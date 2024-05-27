package seller_user

import (
	"context"
	"edot/ecommerce/auth/entity"
	"edot/ecommerce/auth/internal/jwt"
	coreerror "edot/ecommerce/error"
	"fmt"
)

type UserSellerLoginInputValidation struct {
	RequestId string `validate:"required"`
	Username  string `validate:"required,min=5,max=50"`
	Password  string `validate:"required"`
}

func (uc *sellerUserUsecase) Login(ctx context.Context, requestId string, input entity.SellerUser) (user entity.SellerUser, token entity.SellerUserToken, err error) {
	// validate input
	if err = uc.validator.Validate(UserSellerLoginInputValidation{
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
