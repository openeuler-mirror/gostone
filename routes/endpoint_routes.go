package routes

import "work.ctyun.cn/git/GoStack/gostone/service"

var endpointRoutes = []Route{
	//获取节点列表
	{
		Method:  "get",
		Path:    "/v3/endpoints",
		Handler: service.GetAllEndpoints,
	},
	//创建节点
	{
		Method:  "post",
		Path:    "/v3/endpoints",
		Handler: service.CreateEndpoint,
	},
	//更新节点
	{
		Method:  "patch",
		Path:    "/v3/endpoints/:endpoint_id",
		Handler: service.UpdateEndpoint,
	},
	//获取节点
	{
		Method:  "get",
		Path:    "/v3/endpoints/:endpoint_id",
		Handler: service.GetEndpoint,
	},
	//删除节点
	{
		Method:  "delete",
		Path:    "/v3/endpoints/:endpoint_id",
		Handler: service.DeleteEndpoint,
	},
}
