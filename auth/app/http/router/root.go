package router

import (
	"edot/ecommerce/auth/entity"

	"github.com/labstack/echo/v4"
)

func SetupRouter(
	e *echo.Echo,
	userUc entity.IUserUsecase,
	sellerUserUc entity.ISellerUserUsecase,
	authUc entity.IAuthUsecase,
) {

	LoadHealtRouter(e)
	LoadUserRouter(e, userUc)
	LoadSellerUserRouter(e, sellerUserUc)
	LoadAuthRouter(e, authUc)
}
