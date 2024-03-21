package routes

import "work.ctyun.cn/git/GoStack/gostone/service"

var projectRoutes = []Route{
	//获取项目列表
	{
		Method:  "get",
		Path:    "/v3/projects",
		Handler: service.GetProjectList,
	},
	//创建项目
	{
		Method:  "post",
		Path:    "/v3/projects",
		Handler: service.CreateProject,
	},
	//更新项目
	{
		Method:  "patch",
		Path:    "/v3/projects/:project_id",
		Handler: service.UpdateProject,
	},
	//获取项目
	{
		Method:  "get",
		Path:    "/v3/projects/:project_id",
		Handler: service.GetProject,
	},
	//删除项目
	{
		Method:  "delete",
		Path:    "/v3/projects/:project_id",
		Handler: service.DeleteProject,
	},
}
