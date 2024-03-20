package service

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
)

func Healthcheck(ctx echo.Context) error {
	log.Debug("healthcheck")
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"status": "healthy",
	})
}
