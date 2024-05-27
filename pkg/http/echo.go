package http

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func LoadEchoRequiredMiddleware(e *echo.Echo) {
	e.Use(middleware.Secure())
	e.Use(middleware.Recover())
	e.HTTPErrorHandler = EchoErroHandler
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqId := c.Request().Header.Get(RequestIdHeaderKey)

			if reqId == "" {
				uid, _ := uuid.NewV7()
				reqId = uid.String()
			}

			c.Set(RequestIdContextKey, reqId)

			return next(c)
		}
	})
	e.Use(middleware.BodyDumpWithConfig(middleware.BodyDumpConfig{
		Skipper: middleware.DefaultSkipper,
		Handler: func(c echo.Context, reqBody, resBody []byte) {
			s, _ := json.Marshal(map[string]interface{}{
				"request":    string(reqBody),
				"res":        string(resBody),
				"status":     c.Response().Status,
				"uri":        c.Request().RequestURI,
				"request_id": c.Get(RequestIdContextKey),
			})

			switch c.Response().Status {
			case 200, 201, 204:
				c.Logger().Info(string(s))
			case 500:
				c.Logger().Error(string(s))
			default:
				c.Logger().Warn(string(s))
			}
		},
	}))
	e.Use(middleware.RateLimiterWithConfig(middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{Rate: 10, Burst: 30, ExpiresIn: 3 * time.Minute},
		),
		IdentifierExtractor: func(c echo.Context) (string, error) {
			id := c.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return &echo.HTTPError{
				Code:     middleware.ErrExtractorError.Code,
				Message:  middleware.ErrExtractorError.Message,
				Internal: err,
			}
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return &echo.HTTPError{
				Code:     middleware.ErrRateLimitExceeded.Code,
				Message:  middleware.ErrRateLimitExceeded.Message,
				Internal: err,
			}
		},
	}))

	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Skipper:      middleware.DefaultSkipper,
		ErrorMessage: "custom timeout error message returns to client",
		OnTimeoutRouteErrorHandler: func(err error, c echo.Context) {
			fmt.Println(c.Path())
		},
		Timeout: 30 * time.Second,
	}))
}
