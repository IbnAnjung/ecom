package entity

import (
	"context"
)

type IUserRepository interface {
	Create(ctx context.Context, u *User) error
	FindUserByEmailOrPhoneNumber(ctx context.Context, phoneNumber, email *string) (u User, err error)
	FindById(ctx context.Context, id int64) (user User, err error)
}
