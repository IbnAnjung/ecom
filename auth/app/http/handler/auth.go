package handler

import (
	"edot/ecommerce/auth/app/http/handler/presenter"
	"edot/ecommerce/auth/dto"
	"edot/ecommerce/auth/entity"
	coreerror "edot/ecommerce/error"
	pkgHttp "edot/ecommerce/http"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authHandler struct {
	uc entity.IAuthUsecase
}

func NewAuthHandler(
	uc entity.IAuthUsecase,
) *authHandler {
	return &authHandler{
		uc,
	}
}

func (h authHandler) ValidateToken(c echo.Context) error {
	req := presenter.ValidateTokenRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	c.Logger().Infof("Incomming request: request_id: %s %s", requestdId, c.Request().URL)

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		return err
	}

	user, err := h.uc.ValidateUserToken(c.Request().Context(), dto.TokenValidationInput{
		RequestId: requestdId,
		Token:     req.Token,
	})
	if err != nil {
		return err
	}

	res := presenter.ValidateTokenResponse{
		ID:   user.ID,
		Name: user.Name,
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", res))
}

func (h authHandler) ValidateSellerToken(c echo.Context) error {
	req := presenter.ValidateTokenRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	c.Logger().Infof("Incomming request: request_id: %s %s", requestdId, c.Request().URL)

	if err := c.Bind(&req); err != nil {
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		return err
	}

	user, err := h.uc.ValidateSellerUserToken(c.Request().Context(), dto.TokenValidationInput{
		RequestId: requestdId,
		Token:     req.Token,
	})
	if err != nil {
		return err
	}

	res := presenter.ValidateSellerTokenResponse{
		ID: user.ID,
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", res))
}
