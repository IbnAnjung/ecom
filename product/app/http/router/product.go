package router

import (
	"edot/ecommerce/http/echomiddleware"
	"edot/ecommerce/product/app/http/config"
	"edot/ecommerce/product/app/http/handler"
	"edot/ecommerce/product/entity"

	"github.com/labstack/echo/v4"
)

func LoadProductRouter(e *echo.Echo, config config.Config, uc entity.IProductUsecase) {
	h := handler.NewProductHandler(uc)

	r := e.Group("/product")
	r.Use(echomiddleware.AuthenticationMiddleware(config.ServiceURI.Authentication))
	r.GET("", h.GetListProduct)
}
