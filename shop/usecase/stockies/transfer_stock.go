package stockies

import (
	"context"
	coreerror "edot/ecommerce/error"
	"edot/ecommerce/shop/dto"
	"edot/ecommerce/shop/entity"
	"fmt"
)

type TransferStock struct {
	RequestID           string                 `validate:"required"`
	SellerUserID        int64                  `validate:"required"`
	SenderWarehouseID   int64                  `validate:"required"`
	ReceiverWarehouseID int64                  `validate:"required"`
	Products            []TransferStockProduct `validate:"required,dive"`
}

type TransferStockProduct struct {
	ProductID string  `validate:"required,alphanum"`
	Quantity  float64 `validate:"required,min=0"`
}

func (uc *stockies) TransferStock(ctx context.Context, input dto.TransferStockInput) (err error) {
	if len(input.Products) == 0 {
		err = coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "empty products")
		return
	}

	productsIds := make([]string, len(input.Products))
	productValidatorObj := make([]TransferStockProduct, len(input.Products))
	for i, v := range input.Products {
		productsIds[i] = v.ProductID
		productValidatorObj[i] = TransferStockProduct{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		}
	}

	if err = uc.validator.Validate(TransferStock{
		RequestID:           input.RequestID,
		SellerUserID:        input.SellerUserID,
		SenderWarehouseID:   input.SenderWarehouseID,
		ReceiverWarehouseID: input.ReceiverWarehouseID,
		Products:            productValidatorObj,
	}); err != nil {
		fmt.Printf("validation fail, %s %v\n", input.RequestID, err)
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

	// proccess stock like checkout
	whProducts, err := uc.warehouseProductRepo.GetForUpdateProccess(ctx, []int64{input.SenderWarehouseID, input.ReceiverWarehouseID}, productsIds)
	if err != nil {
		return
	}

	valWh := make(chan error)
	trStock := make(chan error)
	trCtx, cancel := context.WithCancel(ctx)
	defer close(valWh)
	defer close(trStock)
	defer cancel()

	// validate user & wareohouse
	warehouse, err := uc.warehouseRepo.GetSellerWarehouse(ctx, input.SellerUserID)
	if err != nil {
		valWh <- err
		return
	}

	go func(c context.Context) {
		defer func() {
			if r := recover(); r != nil {
				err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "recovery validation error\n")
				valWh <- err
				return
			}
		}()

		mapInputWarehouse := map[int64]int8{input.SenderWarehouseID: 1, input.ReceiverWarehouseID: 1}
		inActiveWarehouse := []entity.Warehouse{}
		for _, v := range warehouse {
			if _, ok := mapInputWarehouse[v.ID]; ok {
				if v.Status == entity.WarehouseStatusInActive {
					inActiveWarehouse = append(inActiveWarehouse, v)
				}
				delete(mapInputWarehouse, v.ID)
			}
		}

		if len(mapInputWarehouse) != 0 {
			err := coreerror.NewCoreError(coreerror.CoreErrorTypeForbidden, "You dont have any accesss")
			valWh <- err
			return
		}

		if len(inActiveWarehouse) > 0 {
			err := coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "there is an inactive warehouse, please check request")
			valWh <- err
			return
		}

		valWh <- nil
	}(trCtx)

	// validate and transfer stock
	go func(context.Context) {
		defer func() {
			if r := recover(); r != nil {
				err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "recovery validation error\n")
				trStock <- err
				return
			}
		}()
		mapSenderProducts := map[string]entity.WarehouseProduct{}
		mapReceiverProducts := map[string]entity.WarehouseProduct{}
		updatedProducts := []entity.WarehouseProduct{}
		for _, v := range whProducts {
			if v.WarehouseID != input.SenderWarehouseID {
				mapReceiverProducts[v.ProductID] = v
				continue
			}

			mapSenderProducts[v.ProductID] = v
		}

		for _, v := range input.Products {
			sp, ok := mapSenderProducts[v.ProductID]
			if !ok {
				err := coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "some product in your order out of stock, please try again\n")
				trStock <- err
				return
			}

			rp, ok := mapReceiverProducts[v.ProductID]
			if !ok {
				rp = entity.WarehouseProduct{
					WarehouseID: input.ReceiverWarehouseID,
					ProductID:   v.ProductID,
					Stock:       0,
				}
			}

			if sp.Stock < v.Quantity {
				err = coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "some product in your order out of stock, please try again\n")
				trStock <- err
				return
			}

			sp.Stock -= v.Quantity
			rp.Stock += v.Quantity

			updatedProducts = append(updatedProducts, []entity.WarehouseProduct{sp, rp}...)

		}

		if err = uc.warehouseProductRepo.UpdateStock(ctx, updatedProducts); err != nil {
			trStock <- err
			return
		}

		trStock <- nil
	}(trCtx)

	for i := 0; i < 2; i++ {
		select {
		case validateWarehouse := <-valWh:
			if validateWarehouse != nil && err == nil {
				cancel()
				fmt.Printf("validation error  %s %s\n", input.RequestID, validateWarehouse.Error())
				err = validateWarehouse
			}
		case stockProccess := <-trStock:
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
