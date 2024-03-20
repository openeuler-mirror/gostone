package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"work.ctyun.cn/git/GoStack/gostone/conf"
	"work.ctyun.cn/git/GoStack/gostone/connect"
	"work.ctyun.cn/git/GoStack/gostone/filter"
	"work.ctyun.cn/git/GoStack/gostone/mapper/domain"
	"work.ctyun.cn/git/GoStack/gostone/routes"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	InitGoStone()
	ctx, cancel := context.WithCancel(context.Background())
	wait := &sync.WaitGroup{}
	wait.Add(len(connect.AppConf.GoStone.Port))
	for i := 0; i < len(connect.AppConf.GoStone.Port); i++ {
		port := connect.AppConf.GoStone.Port[i]
		go func(port int) {
			startOne(port, ctx, wait)
		}(port)
	}
	<-quit
	log.Info("wait stop server")
	cancel()
	wait.Wait()
	log.Info("Shutdown Server success")
}

func InitGoStone() {
	connect.InitConf()
	appConf := connect.AppConf
	setLogLevel(appConf.GoStone.LogLevel)
	setLogOutput()
	conf.InitUrl(appConf.GoStone.BaseUrl)
	conf.InitPath(appConf.GoStone.ConfPath)
	connect.InitDB()
	utils.InitTable(connect.GetMysqlConnect(), appConf.GoStone.InitEndpoint, appConf.GoStone.InitRegion)
	utils.InitToken()
	domain.InitDomain()
}

func startOne(port int, ctx context.Context, wait *sync.WaitGroup) {
	e := echo.New()
	e.Use(filter.Error())
	e.Use(filter.SkipAuth())
	e.Use(filter.NewFernetMiddleware())
	e.Use(filter.NewJwtMiddleware(connect.AppConf.GoStone.Secret, connect.AppConf.GoStone.SignMethod))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		ExposeHeaders: []string{"X-Subject-Token", "X-Auth-Token"},
	}))
	//e.Use(filter.RefreshTokenWithConfig(filter.RefreshTokenConfig{
	//	RefreshTime: connect.AppConf.GoStone.RefreshTime,
	//}))
	//启用echo logger
	e.Use(middleware.Logger())
	//自定义数据绑定
	e.Binder = &CustomBinder{}
	//自定义校验
	e.Validator = &CustomValidator{validator: validator.New()}
	//自定义错误处理
	e.HTTPErrorHandler = customHTTPErrorHandler
	//加载所有路由
	for _, route := range routes.GetRoutes() {
		method := route.Method
		switch method {
		case "get":
			e.GET(route.Path, route.Handler)
		case "post":
			e.POST(route.Path, route.Handler)
		case "put":
			e.PUT(route.Path, route.Handler)
		case "delete":
			e.DELETE(route.Path, route.Handler)
		case "options":
			e.OPTIONS(route.Path, route.Handler)
		case "patch":
			e.PATCH(route.Path, route.Handler)
		case "head":
			e.HEAD(route.Path, route.Handler)
		default:
			panic("route http method error!")
		}
	}
	go func() {
		select {
		case <-ctx.Done():
			log.Info("start stop server")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer func() {
				wait.Done()
				cancel()
			}()
			if err := e.Shutdown(ctx); err != nil {
				log.Fatal("Server Shutdown:", err)
			}
			log.Info("Server exiting")
		}
	}()
	log.Fatal(e.Start(":" + strconv.Itoa(port)))
}

func setLogLevel(level string) {
	switch level {
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warning":
		log.SetLevel(log.WarnLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.InfoLevel)
	}
	log.SetReportCaller(true)
	log.SetFormatter(&LogFormatter{})
}

//日志自定义格式
type LogFormatter struct{}

//格式详情
func (s *LogFormatter) Format(entry *log.Entry) ([]byte, error) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	var file string
	var line int
	if entry.Caller != nil {
		file = filepath.Base(entry.Caller.File)
		line = entry.Caller.Line
	}
	level := strings.ToUpper(entry.Level.String())
	msg := fmt.Sprintf("%s [%s] [GOID:%d] [%s:%d] #msg:%s \n", timestamp, level, getGID(), file, line, entry.Message)
	return []byte(msg), nil
}

// 获取当前协程id
func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}

func setLogOutput() {
	logf, err := rotatelogs.New(
		connect.AppConf.GoStone.LogPath+"/gostone_log.%Y%m%d%H%M",
		rotatelogs.WithLinkName(connect.AppConf.GoStone.LogPath+"/gostone.log"),
		rotatelogs.WithMaxAge(time.Duration(connect.AppConf.GoStone.LogMaxAge)*time.Hour),
		rotatelogs.WithRotationTime(time.Duration(connect.AppConf.GoStone.LogRotateTime)*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	defer logf.Close()
	// remove stdout log
	mw := io.MultiWriter(logf)
	log.SetOutput(mw)
}

//自定义请求结构体body内参数校验
type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

//自定义错误处理
func customHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	var message interface{}
	message = err.Error()
	he, ok := err.(*echo.HTTPError)
	if ok {
		code = he.Code
		message = he.Message
	}
	if err := c.JSON(code, map[string]interface{}{
		"status": code,
		"error":  message,
	}); err != nil {
		log.Fatal(err)
		panic("Unexpected Internal Error")
	}
}

//自定义数据绑定
type CustomBinder struct{}

func (cb *CustomBinder) Bind(i interface{}, c echo.Context) (err error) {
	req := c.Request()
	if !strings.HasPrefix(req.Header.Get(echo.HeaderContentType), echo.MIMEApplicationJSON) {
		return echo.ErrUnsupportedMediaType
	}
	//use default binder
	db := new(echo.DefaultBinder)
	if err = db.Bind(i, c); err != echo.ErrUnsupportedMediaType {
		return
	}
	return
}
