package entity

import (
	"context"
)

type ISellerUserUsecase interface {
	RegisterUser(ctx context.Context, requestId string, input SellerUser) (user SellerUser, token SellerUserToken, err error)
	Login(ctx context.Context, requestId string, input SellerUser) (user SellerUser, token SellerUserToken, err error)
}
