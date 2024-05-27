package order

import (
	"context"
	coreerror "edot/ecommerce/error"
	"edot/ecommerce/shop/dto"
	"edot/ecommerce/shop/entity"
	"fmt"
	"time"
)

type CheckoutOrderInput struct {
	RequestID string                      `validate:"required"`
	StoreID   int64                       `validate:"required,number"`
	UserID    int64                       `validate:"required,number"`
	Products  []CheckoutOrderProductInput `validate:"required,dive"`
}

type CheckoutOrderProductInput struct {
	ProductID string  `validate:"required,alphanum"`
	Quantity  float64 `validate:"required,min=1"`
	Price     float64 `validate:"required,min=0"`
}

func (uc *orderUsecase) CheckoutOrder(ctx context.Context, input dto.CheckoutOrderInput) (err error) {
	// validate input
	if len(input.Products) == 0 {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "can't allowed empty product request")
		return
	}

	productsIds := make([]string, len(input.Products))
	inputValidatorProductObj := make([]CheckoutOrderProductInput, len(input.Products))
	for i, v := range input.Products {
		productsIds[i] = v.ProductID
		inputValidatorProductObj[i] = CheckoutOrderProductInput{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
			Price:     v.Price,
		}
	}

	if err = uc.validator.Validate(CheckoutOrderInput{
		RequestID: input.RequestID,
		StoreID:   input.StoreID,
		UserID:    input.UserID,
		Products:  inputValidatorProductObj,
	}); err != nil {
		fmt.Printf("error validate input %s %s", input.RequestID, err.Error())
		return
	}

	uc.gormUow.Begin(ctx)
	defer func() {
		if r := recover(); r != nil {
			uc.gormUow.Rollback(ctx)
			err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "recovery checkouer proccess\n")
			return
		}
	}()

	activeWharehouseIds := []int64{}
	warehouses, err := uc.warehouseRepository.GetStoreActiveWarehouse(ctx, input.StoreID)
	if err != nil {
		fmt.Printf("error get active warehouse %s %s", input.RequestID, err.Error())
		return
	}

	for _, v := range warehouses {
		activeWharehouseIds = append(activeWharehouseIds, v.ID)
	}

	createTime := time.Now()

	chOrder := make(chan error)
	chStock := make(chan error)
	ctxProc, cancel := context.WithCancel(ctx)
	defer close(chOrder)
	defer close(chStock)
	defer cancel()

	whProducts, err := uc.warehouseProductRepository.GetForUpdateProccess(ctx, activeWharehouseIds, productsIds)
	if err != nil {
		return
	}

	go func(ctxP context.Context) {
		mapProducts := map[string][]entity.WarehouseProduct{}
		updatedProducts := []entity.WarehouseProduct{}
		for _, v := range whProducts {
			if _, ok := mapProducts[v.ProductID]; !ok {
				mapProducts[v.ProductID] = []entity.WarehouseProduct{}
			}
			mapProducts[v.ProductID] = append(mapProducts[v.ProductID], v)
		}
		for _, v := range input.Products {
			p, ok := mapProducts[v.ProductID]
			if !ok {
				e := coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "some product in your order out of stock, please try again\n")
				chStock <- e
				return
			}

			for _, w := range p {
				subStock := w.Stock
				if w.Stock > v.Quantity {
					subStock = v.Quantity
				}

				v.Quantity -= subStock
				w.Stock -= subStock
				updatedProducts = append(updatedProducts, w)
				if v.Quantity == 0 {
					break
				}
			}

			if v.Quantity != 0 {
				e := coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "some product in your order out of stock, please try again\n")
				chStock <- e
				return
			}
		}

		if e := uc.warehouseProductRepository.UpdateStock(ctxP, updatedProducts); err != nil {
			chStock <- e
			return
		}

		chStock <- nil
	}(ctxProc)

	go func(ctxP context.Context) {
		order := entity.Order{
			StoreID:     input.StoreID,
			UserID:      input.UserID,
			TotalPrice:  0,
			CreatedTime: createTime,
			ExpiredTime: createTime.Add(24 * time.Hour),
		}

		if e := uc.orderRepository.CreateOrder(ctxP, &order); err != nil {
			chOrder <- e
			return
		}

		details := make([]entity.OrderDetail, len(input.Products))
		for i, v := range input.Products {
			details[i] = entity.OrderDetail{
				OrderID:    order.ID,
				ProductID:  v.ProductID,
				Quantity:   v.Quantity,
				Price:      v.Price,
				TotalPrice: v.Quantity * v.Price,
			}

			order.TotalPrice += details[i].TotalPrice
		}

		if e := uc.orderDetailRepository.CreateBulk(ctxP, &details); err != nil {
			chOrder <- e
			return
		}

		chOrder <- nil
	}(ctxProc)

	for i := 0; i < 2; i++ {
		select {
		case orderProccess := <-chOrder:
			if orderProccess != nil && err == nil {
				cancel()
				fmt.Printf("order fail  %s %s\n", input.RequestID, orderProccess.Error())
				err = orderProccess
			}
		case stockProccess := <-chStock:
			if stockProccess != nil && err == nil {
				cancel()
				fmt.Printf("stockies fail  %s %s\n", input.RequestID, stockProccess.Error())
				err = stockProccess
			}
		}
	}

	if err != nil {
		uc.gormUow.Rollback(ctx)
		return
	}

	uc.gormUow.Commit(ctx)

	return
}
