package user

import (
	"edot/ecommerce/auth/entity"
	"edot/ecommerce/auth/internal/jwt"
	"edot/ecommerce/crypt"
	"edot/ecommerce/structvalidator"
)

type userUsecase struct {
	hasher         crypt.IHash
	jwtService     jwt.IJwtService
	validator      structvalidator.IStructValidator
	userRepository entity.IUserRepository
}

func NewUsecase(
	hasher crypt.IHash,
	jwtService jwt.IJwtService,
	validator structvalidator.IStructValidator,
	userRepository entity.IUserRepository,
) entity.IUserUsecase {
	return &userUsecase{
		hasher, jwtService, validator, userRepository,
	}
}
