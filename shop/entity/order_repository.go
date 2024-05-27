package entity

import (
	"context"
)

type IOrderRepository interface {
	CreateOrder(ctx context.Context, input *Order) (err error)
	FindOrder(ctx context.Context, id int64) (order Order, err error)
	GetExpiredOrder(ctx context.Context) (orders []Order, err error)
	Update(ctx context.Context, orders *Order) (err error)
}
