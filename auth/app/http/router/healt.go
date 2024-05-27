package router

import (
	"edot/ecommerce/auth/app/http/handler"

	"github.com/labstack/echo/v4"
)

func LoadHealtRouter(e *echo.Echo) {
	healtHandler := handler.NewHealtHandler()

	e.GET("/", healtHandler.Check)
}
