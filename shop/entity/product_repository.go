package entity

import (
	"context"
)

type IProductRepository interface {
	FindProductByIds(ctx context.Context, productIDs []string) (products []Product, err error)
}
