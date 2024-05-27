package order

import (
	"context"
	coreerror "edot/ecommerce/error"
	"edot/ecommerce/shop/entity"
)

func (uc *orderUsecase) CancelExpiredOrder(ctx context.Context) {
	orders, err := uc.orderRepository.GetExpiredOrder(ctx)
	if err != nil {
		return
	}

	for _, order := range orders {
		details, err := uc.orderDetailRepository.GetDetails(ctx, order.ID)
		if err != nil {
			return
		}
		productsIds := []string{}
		for _, v := range details {
			productsIds = append(productsIds, v.ProductID)
		}

		activeWharehouseIds := []int64{}
		warehouses, err := uc.warehouseRepository.GetStoreActiveWarehouse(ctx, order.StoreID)
		if err != nil {
			return
		}

		for _, v := range warehouses {
			activeWharehouseIds = append(activeWharehouseIds, v.ID)
		}

		uc.gormUow.Begin(ctx)
		defer func() {
			if r := recover(); r != nil {
				uc.gormUow.Rollback(ctx)
				err = coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "recovery checkouer proccess\n")
				return
			}
		}()

		whProducts, err := uc.warehouseProductRepository.GetForUpdateProccess(ctx, activeWharehouseIds, productsIds)
		if err != nil {
			return
		}

		mapProducts := map[string][]entity.WarehouseProduct{}
		updatedProducts := []entity.WarehouseProduct{}
		for _, v := range whProducts {
			if _, ok := mapProducts[v.ProductID]; !ok {
				mapProducts[v.ProductID] = []entity.WarehouseProduct{}
			}
			mapProducts[v.ProductID] = append(mapProducts[v.ProductID], v)
		}

		for _, v := range details {
			p, ok := mapProducts[v.ProductID]
			if !ok {
				err = coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "some product missing from all warehouse\n")
				uc.gormUow.Rollback(ctx)
				return
			}

			selectedWarehouseProduct := p[0]
			selectedWarehouseProduct.Stock += v.Quantity
			updatedProducts = append(updatedProducts, selectedWarehouseProduct)
		}

		if err = uc.warehouseProductRepository.UpdateStock(ctx, updatedProducts); err != nil {
			uc.gormUow.Rollback(ctx)
			return
		}

		order.Status = entity.OrderStatusExpired
		if err = uc.orderRepository.Update(ctx, &order); err != nil {
			uc.gormUow.Rollback(ctx)
			return
		}

		uc.gormUow.Commit(ctx)

	}
}
