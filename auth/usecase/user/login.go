package user

import (
	"context"
	"edot/ecommerce/auth/dto"
	"edot/ecommerce/auth/entity"
	"edot/ecommerce/auth/internal/jwt"
	coreerror "edot/ecommerce/error"
	"fmt"
)

type LoginInputValidation struct {
	RequestId   string `validate:"required"`
	PhoneNumber string `validate:"required_without=Email,min=11,max=15"`
	Email       string `validate:"required_without=PhoneNumber,min=5,max=50"`
	Password    string `validate:"required"`
}

func (uc *userUsecase) Login(ctx context.Context, input dto.LoginUserInput) (user entity.User, token entity.UserToken, err error) {
	// validate input
	if err = uc.validator.Validate(LoginInputValidation{
		RequestId:   input.RequestId,
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

	if user.ID == 0 {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, "invalid credentials")
		err = e
		return
	}

	// create new user
	if err = uc.hasher.CompareHash(user.Password, input.Password); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, "invalid credentials")
		err = e
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
