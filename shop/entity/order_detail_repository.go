package entity

import "context"

type IOrderDetailRepository interface {
	CreateBulk(ctx context.Context, input *[]OrderDetail) (err error)
	GetDetails(ctx context.Context, orderID int64) (details []OrderDetail, err error)
}
