package auth

import (
	"edot/ecommerce/auth/entity"
	"edot/ecommerce/auth/internal/jwt"
)

type uc struct {
	userRepository       entity.IUserRepository
	sellerUserRepository entity.ISellerUserRepository
	jwt                  jwt.IJwtService
}

func NewUsecase(
	userRepository entity.IUserRepository,
	sellerUserRepository entity.ISellerUserRepository,
	jwt jwt.IJwtService,
) entity.IAuthUsecase {
	return &uc{
		userRepository,
		sellerUserRepository,
		jwt,
	}
}
