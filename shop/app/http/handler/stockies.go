package handler

import (
	coreerror "edot/ecommerce/error"
	pkgHttp "edot/ecommerce/http"
	"edot/ecommerce/shop/app/http/handler/presenter"
	"edot/ecommerce/shop/dto"
	"edot/ecommerce/shop/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type stockiesHandler struct {
	uc entity.IStockiesUsecase
}

func NewStockiesHandler(
	uc entity.IStockiesUsecase,
) *stockiesHandler {
	return &stockiesHandler{
		uc,
	}
}

func (h stockiesHandler) GetProductStock(c echo.Context) error {
	req := presenter.GetProductStockRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	if err := c.Bind(&req); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	stocks, err := h.uc.GetProductStok(c.Request().Context(), dto.GetProductStockInput{
		ProductIDs: req.ProductIDs,
		RequestID:  requestdId,
	})
	if err != nil {
		return err
	}

	resData := make([]presenter.GetProductStockData, len(stocks))
	for i, v := range stocks {
		resData[i] = presenter.GetProductStockData{
			ProductID:  v.ProductID,
			TotalStock: v.TotalStock,
		}
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", resData))
}

func (h stockiesHandler) TransferStock(c echo.Context) error {
	req := presenter.TransferStockRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)
	sellerUserId := c.Get(pkgHttp.SellerUserIdContextKey).(int64)

	if err := c.Bind(&req); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	products := make([]dto.TransferStockProductInput, len(req.Products))
	for i, v := range req.Products {
		products[i] = dto.TransferStockProductInput{
			ProductID: v.ProductID,
			Quantity:  v.Quantity,
		}
	}
	if err := h.uc.TransferStock(c.Request().Context(), dto.TransferStockInput{
		RequestID:           requestdId,
		SellerUserID:        sellerUserId,
		SenderWarehouseID:   req.SenderWarehouseID,
		ReceiverWarehouseID: req.ReceiverWarehouseID,
		Products:            products,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", nil))
}
