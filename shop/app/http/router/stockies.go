package router

import (
	"edot/ecommerce/http/echomiddleware"
	"edot/ecommerce/shop/app/http/config"
	"edot/ecommerce/shop/app/http/handler"
	"edot/ecommerce/shop/entity"

	"github.com/labstack/echo/v4"
)

func LoadStockiesRouter(e *echo.Echo, stockiesUC entity.IStockiesUsecase, config config.Config) {
	h := handler.NewStockiesHandler(stockiesUC)

	e.POST("/stock", h.GetProductStock)
	e.POST("/stock/transfer", h.TransferStock, echomiddleware.SellerAuthenticationMiddleware(config.ServiceURI.Authentication))
}
