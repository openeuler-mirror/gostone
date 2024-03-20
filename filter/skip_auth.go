package filter

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"work.ctyun.cn/git/GoStack/gostone/connect"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

func SkipAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if !connect.AppConf.GoStone.SkipAuth {
				return next(c)
			}
			log.Info("skip auth........")
			tokenValid := utils.GetTokenMethod(utils.TokenType)
			token, _, _ := tokenValid.Sign(utils.AuthContext{
				UserId:    "8f3e215075c54a998d334ca3582869e5",
				ProjectId: "001a18cadd4b401e9fdeab6c411d9816",
				Role:      []string{"admin"},
				DomainId:  "default",
				Method:    []string{"password"},
			})
			c.Request().Header.Set("X-Auth-Token", token)
			return next(c)
		}
	}
}
