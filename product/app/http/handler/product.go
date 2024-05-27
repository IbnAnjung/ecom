package handler

import (
	coreerror "edot/ecommerce/error"
	pkgHttp "edot/ecommerce/http"
	"edot/ecommerce/product/app/http/handler/presenter"
	"edot/ecommerce/product/dto"
	"edot/ecommerce/product/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	uc entity.IProductUsecase
}

func NewProductHandler(
	uc entity.IProductUsecase,
) *productHandler {
	return &productHandler{uc}
}

func (h productHandler) GetListProduct(c echo.Context) error {
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	c.Logger().Infof("Incomming request: %s, request_id: %s", c.Request().RequestURI, requestdId)
	req := presenter.GetListProductRequest{}
	if err := c.Bind(&req); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	products, err := h.uc.GetListProduct(c.Request().Context(), dto.GetListProductInput{
		RequestID: requestdId,
		Keyword:   req.KeyWord,
		Page:      req.Page,
		Limit:     req.Limit,
	})
	if err != nil {
		return err
	}

	res := make([]presenter.GetListProductRespnose, len(products))
	for i, v := range products {
		res[i] = presenter.GetListProductRespnose{
			ID:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Price:       v.Price,
			Stock:       v.Stock,
		}
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", res))
}
