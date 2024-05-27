package handler

import (
	"edot/ecommerce/auth/app/http/handler/presenter"
	"edot/ecommerce/auth/entity"
	coreerror "edot/ecommerce/error"
	pkgHttp "edot/ecommerce/http"
	"net/http"

	"github.com/labstack/echo/v4"
)

type sellerUserHandler struct {
	uc entity.ISellerUserUsecase
}

func NewSellerUserHandler(
	uc entity.ISellerUserUsecase,
) *sellerUserHandler {
	return &sellerUserHandler{
		uc,
	}
}

func (h sellerUserHandler) Register(c echo.Context) error {
	req := presenter.SellerUserRegisterRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	if err := c.Bind(&req); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	user, token, err := h.uc.RegisterUser(c.Request().Context(), requestdId, entity.SellerUser{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return err
	}

	res := presenter.SellerUserRegisterResponse{
		ID:           user.ID,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return c.JSON(http.StatusCreated, pkgHttp.GetStandartSuccessResponse("success", res))
}

func (h sellerUserHandler) Login(c echo.Context) error {
	req := presenter.SellerUserLoginRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	if err := c.Bind(&req); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	user, token, err := h.uc.Login(c.Request().Context(), requestdId, entity.SellerUser{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return err
	}

	res := presenter.RegisterResponse{
		ID:           user.ID,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", res))
}
