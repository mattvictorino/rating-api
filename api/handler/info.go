package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Info struct{}

func NewInfoHandler() Info {
	return Info{}
}

func (h Info) Info(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"service": "rating-api",
		"status":  "running",
	})
}
