package routes

import "work.ctyun.cn/git/GoStack/gostone/service"

var checkRoutes = []Route{
	//进程健康检查
	{
		Method:  "get",
		Path:    "/healthcheck",
		Handler: service.Healthcheck,
	},
}
