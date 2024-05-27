package entity

import (
	"context"
	"edot/ecommerce/auth/dto"
)

type IUserUsecase interface {
	RegisterUser(ctx context.Context, input dto.RegisterUserInput) (user User, token UserToken, err error)
	Login(ctx context.Context, input dto.LoginUserInput) (user User, token UserToken, err error)
}
