package router

import (
	"edot/ecommerce/http/echomiddleware"
	"edot/ecommerce/shop/app/http/config"
	"edot/ecommerce/shop/app/http/handler"
	"edot/ecommerce/shop/entity"

	"github.com/labstack/echo/v4"
)

func LoadWarehouseRouter(e *echo.Echo, uc entity.IWarehouseUsecase, config config.Config) {
	h := handler.NewWarehouseHandler(uc)

	r := e.Group("/warehouse")
	r.Use(echomiddleware.SellerAuthenticationMiddleware(config.ServiceURI.Authentication))
	r.PUT("/toggle-status/:warehouse_id", h.ToggleStatus)
}
