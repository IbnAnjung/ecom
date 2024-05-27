package router

import (
	"edot/ecommerce/shop/app/http/handler"

	"github.com/labstack/echo/v4"
)

func LoadHealtRouter(e *echo.Echo) {
	healtHandler := handler.NewHealtHandler()

	e.GET("/", healtHandler.Check)
}
