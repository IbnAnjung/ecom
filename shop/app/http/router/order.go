package router

import (
	"edot/ecommerce/http/echomiddleware"
	"edot/ecommerce/shop/app/http/config"
	"edot/ecommerce/shop/app/http/handler"
	"edot/ecommerce/shop/entity"

	"github.com/labstack/echo/v4"
)

func LoadOrderRouter(e *echo.Echo, uc entity.IOrderUsecase, config config.Config) {
	h := handler.NewOrderHandler(uc)

	r := e.Group("/order")
	r.Use(echomiddleware.AuthenticationMiddleware(config.ServiceURI.Authentication))
	r.POST("/checkout", h.Chekcout)
}
