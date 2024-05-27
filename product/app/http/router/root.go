package router

import (
	"edot/ecommerce/product/app/http/config"
	"edot/ecommerce/product/entity"

	"github.com/labstack/echo/v4"
)

func SetupRouter(
	e *echo.Echo,
	config config.Config,
	productUc entity.IProductUsecase,
) {

	LoadHealtRouter(e)
	LoadProductRouter(e, config, productUc)
}
