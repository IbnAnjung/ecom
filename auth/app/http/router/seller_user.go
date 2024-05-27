package router

import (
	"edot/ecommerce/auth/app/http/handler"
	"edot/ecommerce/auth/entity"

	"github.com/labstack/echo/v4"
)

func LoadSellerUserRouter(e *echo.Echo, uc entity.ISellerUserUsecase) {
	h := handler.NewSellerUserHandler(uc)

	r := e.Group("/seller")
	r.POST("/register", h.Register)
	r.POST("/login", h.Login)
}
