package routes

import "work.ctyun.cn/git/GoStack/gostone/service"

var versionRoutes = []Route{
	//获取所有版本
	{
		Method:  "get",
		Path:    "/",
		Handler: service.GetAllVersion,
	},
	{
		Method:  "get",
		Path:    "",
		Handler: service.GetAllVersion,
	},
	//获取V3版本信息
	{
		Method:  "get",
		Path:    "/v3",
		Handler: service.GetVersion3,
	},
	{
		Method:  "get",
		Path:    "/v3/",
		Handler: service.GetVersion3,
	},
	//获取V2版本信息
	{
		Method:  "get",
		Path:    "/v2/",
		Handler: service.GetVersion2,
	},
	{
		Method:  "get",
		Path:    "/v2",
		Handler: service.GetVersion2,
	},
}
