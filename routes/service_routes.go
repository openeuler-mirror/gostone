package routes

import "work.ctyun.cn/git/GoStack/gostone/service"

var serviceRoutes = []Route{
	//获取节点列表
	{
		Method:  "get",
		Path:    "/v3/services",
		Handler: service.GetAllServices,
	},
	//创建节点
	{
		Method:  "post",
		Path:    "/v3/services",
		Handler: service.CreateService,
	},
	//更新节点
	{
		Method:  "patch",
		Path:    "/v3/services/:service_id",
		Handler: service.UpdateService,
	},
	//获取节点
	{
		Method:  "get",
		Path:    "/v3/services/:service_id",
		Handler: service.GetService,
	},
	//删除节点
	{
		Method:  "delete",
		Path:    "/v3/services/:service_id",
		Handler: service.DeleteService,
	},
}
