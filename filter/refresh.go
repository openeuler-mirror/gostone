package filter

import (
	"github.com/labstack/echo/v4"
	EMiddleware "github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
	"work.ctyun.cn/git/GoStack/gostone/service"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

//用于自动刷新jwt token 检查token到期时间是否小于5分钟,是的话则重新生成token 存入到header中 X-Refresh-Token

type RefreshTokenConfig struct {
	Skipper     EMiddleware.Skipper
	HeaderName  string
	RefreshTime float64
}

func DefaultSkipper(ctx echo.Context) bool {
	return (ctx.Request().Method == echo.POST &&
		strings.Contains(ctx.Path(), "/v3/auth/tokens")) ||
		ctx.Path() == "/v3" ||
		ctx.Path() == "/" ||
		ctx.Path() == "" ||
		ctx.Path() == "/v3/" || ctx.Request().Method == echo.OPTIONS
}

var (
	DefaultRefreshTokenConfig = RefreshTokenConfig{
		Skipper:     DefaultSkipper,
		HeaderName:  "X-Refresh-Token",
		RefreshTime: 5,
	}
)

func RefreshToken() echo.MiddlewareFunc {
	return RefreshTokenWithConfig(DefaultRefreshTokenConfig)
}

func RefreshTokenWithConfig(config RefreshTokenConfig) echo.MiddlewareFunc {
	if config.Skipper == nil {
		config.Skipper = DefaultRefreshTokenConfig.Skipper
	}
	if config.RefreshTime == 0 {
		config.RefreshTime = DefaultRefreshTokenConfig.RefreshTime
	}
	if config.HeaderName == "" {
		config.HeaderName = DefaultRefreshTokenConfig.HeaderName
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}
			auth := utils.GetTokenMethod(c.Get(utils.TokenTypeKey)).GetAuthContext(c)
			t := time.Time(auth.ExpiresAtZ).Sub(time.Now().UTC()).Minutes()
			log.Infof("token will expire after %v minutes ", t)
			if t < 5 {
				//需要刷新token
				token := service.GetUserToken(auth.UserId)
				c.Response().Header().Add(config.HeaderName, token)
			}
			return next(c)
		}
	}
}
