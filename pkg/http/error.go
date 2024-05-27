package http

import (
	coreerror "edot/ecommerce/error"
	"encoding/json"
	"net/http"

	"github.com/labstack/echo/v4"
)

func EchoErroHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	httpCodeError := http.StatusInternalServerError
	data := map[string]interface{}{
		"message": "Internal server error",
	}

	if cErr, ok := err.(coreerror.CoreError); ok {
		if cErr.Type != coreerror.CoreErrorTypeInternalServerError {
			data["message"] = err.Error()
		}

		switch cErr.Type {
		case coreerror.CoreErrorTypeForbidden:
			httpCodeError = http.StatusForbidden
		case coreerror.CoreErrorTypeAuthorization:
			httpCodeError = http.StatusUnauthorized
		case coreerror.CoreErrorTypeNotFound:
			httpCodeError = http.StatusNotFound
		case coreerror.CoreErrorTypeUnprocessable:
			httpCodeError = http.StatusUnprocessableEntity
		}
	} else if valErr, ok := err.(coreerror.ValidationError); ok {
		httpCodeError = http.StatusBadRequest
		data["message"] = "validation error"
		data["validation_error"] = valErr.GetMessage()
	} else if eErr, ok := err.(*echo.HTTPError); ok {
		httpCodeError = eErr.Code

		if httpCodeError != http.StatusInternalServerError {
			data["message"] = eErr.Message
		}
	}

	if httpCodeError == http.StatusInternalServerError {
		s, _ := json.Marshal(map[string]interface{}{
			"uri":        c.Request().RequestURI,
			"request_id": c.Get(RequestIdContextKey),
			"Error":      err.Error(),
		})

		c.Logger().Error(string(s))
	}

	c.JSON(httpCodeError, data)
}
