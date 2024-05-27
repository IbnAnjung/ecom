package entity

import (
	"context"
	"edot/ecommerce/shop/dto"
)

type IStockiesUsecase interface {
	GetProductStok(ctx context.Context, input dto.GetProductStockInput) (stocks []dto.ProductStock, err error)
	TransferStock(ctx context.Context, input dto.TransferStockInput) (err error)
}
