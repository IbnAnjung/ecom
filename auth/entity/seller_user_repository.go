package entity

import (
	"context"
)

type ISellerUserRepository interface {
	Create(ctx context.Context, u *SellerUser) error
	FindUserByUsername(ctx context.Context, username string) (u SellerUser, err error)
	FindById(ctx context.Context, id int64) (user SellerUser, err error)
}
