package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type healtHandler struct{}

func NewHealtHandler() *healtHandler {
	return &healtHandler{}
}

func (h healtHandler) Check(c echo.Context) error {
	time.Sleep(5 * time.Second)
	return c.JSON(http.StatusOK, "OK")
}
