package auth

import (
	"context"
	"edot/ecommerce/auth/dto"
	"edot/ecommerce/auth/entity"
	"edot/ecommerce/auth/internal/jwt"
	coreerror "edot/ecommerce/error"
	"fmt"
)

func (uc *uc) ValidateUserToken(ctx context.Context, input dto.TokenValidationInput) (user entity.User, err error) {
	claim, err := uc.jwt.ValidateToken(input.Token, jwt.UserTypeUser)
	if err != nil {
		return
	}

	user, err = uc.userRepository.FindById(ctx, claim.UserID)
	if err != nil {
		fmt.Printf("error get user request_id: %s %s\n", input.RequestId, err.Error())
		if cErr, ok := err.(coreerror.CoreError); ok {
			if cErr.Type == coreerror.CoreErrorTypeNotFound {
				e := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, "")
				err = e
			}
		}
		return
	}

	return
}

func (uc *uc) ValidateSellerUserToken(ctx context.Context, input dto.TokenValidationInput) (user entity.SellerUser, err error) {
	claim, err := uc.jwt.ValidateToken(input.Token, jwt.UserTypeSeller)
	if err != nil {
		return
	}

	user, err = uc.sellerUserRepository.FindById(ctx, claim.UserID)
	if err != nil {
		fmt.Printf("error get user request_id: %s %s\n", input.RequestId, err.Error())
		if cErr, ok := err.(coreerror.CoreError); ok {
			if cErr.Type == coreerror.CoreErrorTypeNotFound {
				e := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, "")
				err = e
			}
		}
		return
	}

	return
}
