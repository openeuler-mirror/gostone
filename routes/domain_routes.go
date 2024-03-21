package routes

import "work.ctyun.cn/git/GoStack/gostone/service"

var domainRoutes = []Route{
	//获取项目列表
	{
		Method:  "get",
		Path:    "/v3/domains",
		Handler: service.GetDomainList,
	},
	//创建项目
	{
		Method:  "post",
		Path:    "/v3/domains",
		Handler: service.CreateDomain,
	},
	//更新项目
	{
		Method:  "patch",
		Path:    "/v3/domains/:domain_id",
		Handler: service.UpdateDomain,
	},
	//获取项目
	{
		Method:  "get",
		Path:    "/v3/domains/:domain_id",
		Handler: service.GetDomain,
	},
	//删除项目
	{
		Method:  "delete",
		Path:    "/v3/domains/:domain_id",
		Handler: service.DeleteDomain,
	},
}
