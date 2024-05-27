package router

import (
	"edot/ecommerce/shop/app/http/config"
	"edot/ecommerce/shop/entity"

	"github.com/labstack/echo/v4"
)

func SetupRouter(
	e *echo.Echo,
	stockiesUC entity.IStockiesUsecase,
	orderUc entity.IOrderUsecase,
	whUc entity.IWarehouseUsecase,
	config config.Config,
) {
	LoadStockiesRouter(e, stockiesUC, config)
	LoadHealtRouter(e)
	LoadOrderRouter(e, orderUc, config)
	LoadWarehouseRouter(e, whUc, config)
}
