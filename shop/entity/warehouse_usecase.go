package entity

import (
	"context"
	"edot/ecommerce/shop/dto"
)

type IWarehouseUsecase interface {
	ToggleStatus(ctx context.Context, input dto.ToogleStatusWarehouseInput) (err error)
}
