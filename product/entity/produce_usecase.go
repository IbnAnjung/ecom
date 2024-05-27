package entity

import (
	"context"
	"edot/ecommerce/product/dto"
)

type IProductUsecase interface {
	GetListProduct(ctx context.Context, input dto.GetListProductInput) (products []dto.FindProductOutput, err error)
}
