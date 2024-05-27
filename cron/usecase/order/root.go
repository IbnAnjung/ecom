package order

import "edot/ecommerce/cron/entity"

type orderUsecase struct {
}

func NewUsecase() entity.IOrder {
	return &orderUsecase{}
}
