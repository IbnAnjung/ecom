package entity

import (
	"context"
	"edot/ecommerce/product/dto"
)

type IProductRepository interface {
	Find(ctx context.Context, input dto.GetListProductInput) (products []Product, err error)
	FindProductByIds(ctx context.Context, productIDs []string) (products []Product, err error)
}
