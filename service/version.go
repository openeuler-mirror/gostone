package service

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/conf"
)

func GetVersion3(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"version": conf.Version3,
	})
}

func GetVersion2(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"version": conf.Version2,
	})
}

func GetAllVersion(ctx echo.Context) error {
	return ctx.JSON(http.StatusMultipleChoices, map[string]interface{}{
		"versions": map[string]interface{}{
			"values": []interface{}{
				conf.Version2, conf.Version3,
			},
		},
	})
}
