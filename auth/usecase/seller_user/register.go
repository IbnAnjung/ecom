package seller_user

import (
	"context"
	"edot/ecommerce/auth/entity"
	"edot/ecommerce/auth/internal/jwt"
	coreerror "edot/ecommerce/error"
	"fmt"
)

type RegisterInputValidation struct {
	RequestId string `validate:"required"`
	Username  string `validate:"required,max=50,min=3"`
	Password  string `validate:"required,min=6"`
}

func (uc *sellerUserUsecase) RegisterUser(ctx context.Context, requestId string, input entity.SellerUser) (user entity.SellerUser, token entity.SellerUserToken, err error) {
	// validate input
	if err = uc.validator.Validate(RegisterInputValidation{
		RequestId: requestId,
		Username:  input.Username,
		Password:  input.Password,
	}); err != nil {
		return
	}

	// find user
	user, err = uc.sellerUserRepository.FindUserByUsername(ctx, input.Username)
	if err != nil {
		fmt.Printf("get user, request_id:%s  %v\n", requestId, err.Error())
		return
	}

	if user.ID != 0 {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeUnprocessable, "username already registered")
		err = e
		fmt.Printf("get user, request_id:%s  %v\n", requestId, "username already registered")
		return
	}

	// create new user
	hashedPassword, err := uc.hasher.HashString(input.Password)
	if err != nil {
		fmt.Printf("get user, request_id:%s  %v\n", requestId, err)
		return
	}

	user = entity.SellerUser{
		Username: input.Username,
		Password: hashedPassword,
	}

	if err = uc.sellerUserRepository.Create(ctx, &user); err != nil {
		fmt.Printf("get user, request_id:%s  %v\n", requestId, err)
		return
	}

	// generate token
	userClaim := jwt.UserClaim{
		UserID: user.ID,
	}

	token.AccessToken, err = uc.jwtService.GenerateAccessToken(userClaim, jwt.UserTypeSeller)
	if err != nil {
		fmt.Printf("get user, request_id:%s  %v\n", requestId, err)
		return
	}

	token.RefreshToken, err = uc.jwtService.GenerateRefreshToken(userClaim, jwt.UserTypeSeller)
	if err != nil {
		fmt.Printf("get user, request_id:%s  %v\n", requestId, err)
		return
	}

	return
}
