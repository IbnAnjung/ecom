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

type orderHandler struct {
	uc entity.IOrderUsecase
}

func NewOrderHandler(
	uc entity.IOrderUsecase,
) *orderHandler {
	return &orderHandler{uc}
}

func (h orderHandler) Chekcout(c echo.Context) error {
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)
	userID := c.Get(pkgHttp.UserIdContextKey).(int64)

	c.Logger().Infof("Incomming request: %s, request_id: %s", c.Request().RequestURI, requestdId)
	req := presenter.CheckoutRequest{}
	if err := c.Bind(&req); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	inputProducts := make([]dto.CheckoutOrderProductInput, len(req.Products))
	for i, v := range req.Products {
		inputProducts[i] = dto.CheckoutOrderProductInput{
			ProductID: v.ID,
			Quantity:  v.Quantity,
			Price:     v.Price,
		}
	}
	if err := h.uc.CheckoutOrder(c.Request().Context(), dto.CheckoutOrderInput{
		RequestID: requestdId,
		StoreID:   req.StoreID,
		UserID:    userID,
		Products:  inputProducts,
	}); err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, pkgHttp.GetStandartSuccessResponse("success", nil))
}
