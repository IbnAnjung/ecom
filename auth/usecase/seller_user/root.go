package seller_user

import (
	"edot/ecommerce/auth/entity"
	"edot/ecommerce/auth/internal/jwt"
	"edot/ecommerce/crypt"
	"edot/ecommerce/structvalidator"
)

type sellerUserUsecase struct {
	hasher               crypt.IHash
	jwtService           jwt.IJwtService
	validator            structvalidator.IStructValidator
	sellerUserRepository entity.ISellerUserRepository
}

func NewUsecase(
	hasher crypt.IHash,
	jwtService jwt.IJwtService,
	validator structvalidator.IStructValidator,
	sellerUserRepository entity.ISellerUserRepository,
) entity.ISellerUserUsecase {
	return &sellerUserUsecase{
		hasher, jwtService, validator, sellerUserRepository,
	}
}
