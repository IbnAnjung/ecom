package user

import (
	"context"
	"edot/ecommerce/auth/dto"
	"edot/ecommerce/auth/entity"
	"edot/ecommerce/auth/internal/jwt"
	coreerror "edot/ecommerce/error"
	"fmt"
	"time"
)

type RegisterInputValidation struct {
	RequestId   string `validate:"required"`
	Name        string `validate:"required,max=50,min=3"`
	PhoneNumber string `validate:"required,min=11,max=15"`
	Email       string `validate:"required,min=5,max=50"`
	Password    string `validate:"required,min=6"`
}

func (uc *userUsecase) RegisterUser(ctx context.Context, input dto.RegisterUserInput) (user entity.User, token entity.UserToken, err error) {
	// validate input
	if err = uc.validator.Validate(RegisterInputValidation{
		RequestId:   input.RequestId,
		Name:        input.Name,
		PhoneNumber: input.PhoneNumber,
		Email:       input.Email,
		Password:    input.Password,
	}); err != nil {
		return
	}

	// find user
	user, err = uc.userRepository.FindUserByEmailOrPhoneNumber(ctx, &input.PhoneNumber, &input.Email)
	if err != nil {
		fmt.Printf("get user, request_id:%s  %v\n", input.RequestId, err.Error())
		return
	}

	if user.ID != 0 {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "phone_number or email already registered")
		err = e
		fmt.Printf("get user, request_id:%s  %v\n", input.RequestId, "phone_number or email already registered")
		return
	}

	// create new user
	hashedPassword, err := uc.hasher.HashString(input.Password)
	if err != nil {
		fmt.Printf("get user, request_id:%s  %v\n", input.RequestId, err)
		return
	}

	user = entity.User{
		Name:        input.Name,
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    hashedPassword,
		CreatedAt:   time.Now(),
	}

	if err = uc.userRepository.Create(ctx, &user); err != nil {
		fmt.Printf("get user, request_id:%s  %v\n", input.RequestId, err)
		return
	}

	// generate token
	userClaim := jwt.UserClaim{
		UserID:   user.ID,
		UserName: user.Name,
	}

	token.AccessToken, err = uc.jwtService.GenerateAccessToken(userClaim, jwt.UserTypeUser)
	if err != nil {
		fmt.Printf("get user, request_id:%s  %v\n", input.RequestId, err)
		return
	}

	token.RefreshToken, err = uc.jwtService.GenerateRefreshToken(userClaim, jwt.UserTypeUser)
	if err != nil {
		fmt.Printf("get user, request_id:%s  %v\n", input.RequestId, err)
		return
	}

	return
}
