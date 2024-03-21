package routes

import "work.ctyun.cn/git/GoStack/gostone/service"

var roleRoutes = []Route{
	//获取节点列表
	{
		Method:  "get",
		Path:    "/v3/roles",
		Handler: service.GetAllRoles,
	},
	//创建节点
	{
		Method:  "post",
		Path:    "/v3/roles",
		Handler: service.CreateRole,
	},
	//更新节点
	{
		Method:  "patch",
		Path:    "/v3/roles/:role_id",
		Handler: service.UpdateRole,
	},
	//获取节点
	{
		Method:  "get",
		Path:    "/v3/roles/:role_id",
		Handler: service.GetRole,
	},
	//删除节点
	{
		Method:  "delete",
		Path:    "/v3/roles/:role_id",
		Handler: service.DeleteRole,
	},
}
