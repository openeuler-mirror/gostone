package filter

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"work.ctyun.cn/git/GoStack/gostone/filter/middleware"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

func Skipper(ctx echo.Context) bool {
	return (ctx.Request().Method == echo.POST &&
		strings.Contains(ctx.Path(), "/v3/auth/tokens")) ||
		ctx.Path() == "/v3" ||
		ctx.Path() == "/" ||
		ctx.Path() == "" ||
		ctx.Path() == "/v3/" || ctx.Request().Method == echo.OPTIONS || ctx.Get("HasAuth") != nil
}

func ErrorHandlerWithContext(err error, ctx echo.Context) error {
	ge, ok := err.(jwt.ValidationError)
	if ok {
		if ge.Errors == jwt.ValidationErrorExpired {
			//todo 如果是过期，需要比对是否含有允许过期的标志
		}
	}
	return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
		"error": map[string]interface{}{
			"status":  http.StatusUnauthorized,
			"message": err.Error(),
		},
	})
}

func NewJwtMiddleware(secret string, signMethod string) echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:                 Skipper,
		ErrorHandlerWithContext: ErrorHandlerWithContext,
		SigningKey:              []byte(secret),
		TokenLookup:             "header:X-Auth-Token",
		Claims:                  &utils.AuthContext{},
		SigningMethod:           signMethod,
	})
}
