package routes

import "work.ctyun.cn/git/GoStack/gostone/service"

var assignRoutes = []Route{
	//获取assignment列表
	{
		Method:  "get",
		Path:    "/v3/projects/:project_id/users/:user_id/roles",
		Handler: service.FindRoleListByProjectAndUser,
	},
	//获取所有assignment
	{
		Method:  "get",
		Path:    "/v3/role_assignments",
		Handler: service.GetAllAssignment,
	},
	//保存assignment
	{
		Method:  "put",
		Path:    "/v3/projects/:project_id/users/:user_id/roles/:role_id",
		Handler: service.SaveRole2UserOnProject,
	},
	{
		Method:  "put",
		Path:    "/v3/projects/name/:project_name/users/:user_id/roles/:role_id",
		Handler: service.SaveRole2UserOnProjectName,
	},
	//校验权限
	{
		Method:  "head",
		Path:    "/v3/projects/:project_id/users/:user_id/roles/:role_id",
		Handler: service.ValidRole2UserOnProject,
	},
	//删除权限
	{
		Method:  "delete",
		Path:    "/v3/projects/:project_id/users/:user_id/roles/:role_id",
		Handler: service.DeleteRole2UserOnProject,
	},
}
