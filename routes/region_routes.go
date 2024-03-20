package routes

import "work.ctyun.cn/git/GoStack/gostone/service"

var regionRoutes = []Route{
	//获取节点列表
	{
		Method:  "get",
		Path:    "/v3/regions",
		Handler: service.GetAllRegions,
	},
	//创建节点
	{
		Method:  "post",
		Path:    "/v3/regions",
		Handler: service.CreateRegion,
	},
	//更新节点
	{
		Method:  "patch",
		Path:    "/v3/regions/:region_id",
		Handler: service.UpdateRegion,
	},
	//获取节点
	{
		Method:  "get",
		Path:    "/v3/regions/:region_id",
		Handler: service.GetRegion,
	},
	//删除节点
	{
		Method:  "delete",
		Path:    "/v3/regions/:region_id",
		Handler: service.DeleteRegion,
	},
}
