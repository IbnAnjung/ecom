package router

import (
	"edot/ecommerce/auth/app/http/handler"
	"edot/ecommerce/auth/entity"

	"github.com/labstack/echo/v4"
)

func LoadUserRouter(e *echo.Echo, userUC entity.IUserUsecase) {
	h := handler.NewUserHandler(userUC)

	e.POST("/register", h.Register)
	e.POST("/login", h.Login)
}
