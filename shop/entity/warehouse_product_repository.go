package entity

import (
	"context"
	"edot/ecommerce/shop/dto"
)

type IWarehouseProductRepository interface {
	GetProductTotalStock(ctx context.Context, productIDs []string) (stocks []dto.ProductStock, err error)
	GetForUpdateProccess(ctx context.Context, warehouseId []int64, productIDs []string) (whProducts []WarehouseProduct, err error)
	UpdateStock(ctx context.Context, whProducts []WarehouseProduct) (err error)
}
