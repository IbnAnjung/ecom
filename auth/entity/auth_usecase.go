package entity

import (
	"context"
	"edot/ecommerce/auth/dto"
)

type IAuthUsecase interface {
	ValidateUserToken(ctx context.Context, input dto.TokenValidationInput) (user User, err error)
	ValidateSellerUserToken(ctx context.Context, input dto.TokenValidationInput) (user SellerUser, err error)
}
