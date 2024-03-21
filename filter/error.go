package filter

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"runtime/debug"
	"work.ctyun.cn/git/GoStack/gostone/execption"
)

//统一的前置异常处理,用于panic的异常捕捉
func Error() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			defer func() {
				if err := recover(); err != nil {
					debug.PrintStack()
					log.Errorln(err)
					log.Errorf(string(debug.Stack()))
					code := http.StatusInternalServerError
					ge, ok := err.(execption.GoStoneError)
					if ok {
						code = ge.Code
						ctx.JSON(code, map[string]interface{}{
							"error": map[string]interface{}{
								"status":  code,
								"message": ge.Message,
							},
						})
					} else {
						ctx.JSON(code, map[string]interface{}{
							"error": map[string]interface{}{
								"status":  code,
								"message": debug.Stack(),
							},
						})
					}

				}
			}()
			return next(ctx)
		}
	}

}
