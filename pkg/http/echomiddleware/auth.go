package echomiddleware

import (
	"bytes"
	coreerror "edot/ecommerce/error"
	pkghttp "edot/ecommerce/http"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AuthDataResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type AuthResponse struct {
	Message string           `json:"string"`
	Data    AuthDataResponse `json:"data"`
}

func AuthenticationMiddleware(authServiceBaseURI string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestId := c.Get(pkghttp.RequestIdContextKey).(string)

			bearer := c.Request().Header.Get("Authorization")
			tokens := strings.Split(bearer, " ")

			if bearer == "" || len(tokens) != 2 {
				err := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, "")
				return err
			}

			authRequest := map[string]string{"token": tokens[1]}
			jsonBody, _ := json.Marshal(authRequest)

			req, err := http.NewRequestWithContext(c.Request().Context(), "PATCH", fmt.Sprintf("%s/validate", authServiceBaseURI), bytes.NewReader(jsonBody))
			if err != nil {
				err := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
				return err
			}

			req.Header.Set(pkghttp.RequestIdHeaderKey, requestId)
			req.Header.Set("Content-Type", "application/json")
			res, err := http.DefaultClient.Do(req)
			if err != nil {
				err := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
				return err
			}
			defer res.Body.Close()
			if res.StatusCode != http.StatusOK {
				err := coreerror.NewCoreError(coreerror.CoreErrorTypeAuthorization, "")
				return err
			}

			resBody := AuthResponse{}
			bodyBytes, err := io.ReadAll(res.Body)
			if err != nil {
				err := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
				return err
			}

			if err := json.Unmarshal(bodyBytes, &resBody); err != nil {
				err := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
				return err
			}

			if err := json.Unmarshal(bodyBytes, &resBody); err != nil {
				err := coreerror.NewCoreError(coreerror.CoreErrorTypeInternalServerError, err.Error())
				return err
			}

			c.Set(pkghttp.UserIdContextKey, resBody.Data.ID)

			return next(c)
		}
	}
}
