package entity

import (
	"context"
	"edot/ecommerce/shop/dto"
)

type IOrderUsecase interface {
	CheckoutOrder(ctx context.Context, input dto.CheckoutOrderInput) error
	CancelExpiredOrder(ctx context.Context)
}
