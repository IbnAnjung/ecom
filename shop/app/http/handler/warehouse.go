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

type warehouseHandler struct {
	uc entity.IWarehouseUsecase
}

func NewWarehouseHandler(
	uc entity.IWarehouseUsecase,
) *warehouseHandler {
	return &warehouseHandler{uc}
}

func (h warehouseHandler) ToggleStatus(c echo.Context) error {
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)
	userID := c.Get(pkgHttp.SellerUserIdContextKey).(int64)

	c.Logger().Infof("Incomming request: %s, request_id: %s", c.Request().RequestURI, requestdId)
	req := presenter.WarehouseToggleStatusRequest{}
	if err := c.Bind(&req); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	if err := h.uc.ToggleStatus(c.Request().Context(), dto.ToogleStatusWarehouseInput{
		RequestID:    requestdId,
		SellerUserID: userID,
		WarehouseID:  req.WarehouseID,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, pkgHttp.GetStandartSuccessResponse("success", nil))
}
