package filter

import (
	"github.com/labstack/echo/v4"
	"strings"
	"work.ctyun.cn/git/GoStack/gostone/filter/middleware"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

func FernetSkipper(ctx echo.Context) bool {
	token := ctx.Request().Header.Get("X-Auth-Token")
	return (ctx.Request().Method == echo.POST &&
		strings.Contains(ctx.Path(), "/v3/auth/tokens")) ||
		ctx.Path() == "/v3" ||
		ctx.Path() == "/" ||
		ctx.Path() == "" ||
		ctx.Path() == "/v3/" || ctx.Request().Method == echo.OPTIONS || utils.IsJwtToken(token)
}

func NewFernetMiddleware() echo.MiddlewareFunc {
	return middleware.FERNETWithConfig(middleware.FERNETConfig{
		Skipper:     FernetSkipper,
		TokenLookup: "header:X-Auth-Token",
		Validate:    utils.GetTokenMethod("fernet"),
	})
}
