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

type userHandler struct {
	uc entity.IUserUsecase
}

func NewUserHandler(
	uc entity.IUserUsecase,
) *userHandler {
	return &userHandler{
		uc,
	}
}

func (h userHandler) Register(c echo.Context) error {
	req := presenter.RegisterRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	if err := c.Bind(&req); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	user, token, err := h.uc.RegisterUser(c.Request().Context(), dto.RegisterUserInput{
		RequestId:   requestdId,
		Name:        req.Name,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})
	if err != nil {
		return err
	}

	res := presenter.RegisterResponse{
		ID:           user.ID,
		Name:         user.Name,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return c.JSON(http.StatusCreated, pkgHttp.GetStandartSuccessResponse("success", res))
}

func (h userHandler) Login(c echo.Context) error {
	req := presenter.LoginRequest{}
	requestdId := c.Get(pkgHttp.RequestIdContextKey).(string)

	if err := c.Bind(&req); err != nil {
		e := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, "")
		err = e
		c.Logger().Errorf("fail binding request request_id:%s, %v", requestdId, err)
		return err
	}

	user, token, err := h.uc.Login(c.Request().Context(), dto.LoginUserInput{
		RequestId:   requestdId,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	})
	if err != nil {
		return err
	}

	res := presenter.RegisterResponse{
		ID:           user.ID,
		Name:         user.Name,
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
	}

	return c.JSON(http.StatusOK, pkgHttp.GetStandartSuccessResponse("success", res))
}
