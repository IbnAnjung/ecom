package router

import (
	"edot/ecommerce/auth/app/http/handler"
	"edot/ecommerce/auth/entity"

	"github.com/labstack/echo/v4"
)

func LoadAuthRouter(e *echo.Echo, uc entity.IAuthUsecase) {
	h := handler.NewAuthHandler(uc)

	e.PATCH("/validate", h.ValidateToken)
	e.PATCH("/validate_seller", h.ValidateSellerToken)
}
