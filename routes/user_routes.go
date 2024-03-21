package routes

import "work.ctyun.cn/git/GoStack/gostone/service"

var userRoutes = []Route{
	//获取用户列表
	{
		Method:  "get",
		Path:    "/v3/users",
		Handler: service.GetUsers,
	},
	//获取指定用户
	{
		Method:  "get",
		Path:    "/v3/users/:user_id",
		Handler: service.GetUser,
	},
	//创建用户
	{
		Method:  "post",
		Path:    "/v3/users",
		Handler: service.CreateUser,
	},
	//更新用户
	{
		Method:  "patch",
		Path:    "/v3/users/:user_id",
		Handler: service.UpdateUser,
	},
	//删除用户
	{
		Method:  "delete",
		Path:    "/v3/users/:user_id",
		Handler: service.DeleteUser,
	},
	//修改用户密码
	{
		Method:  "post",
		Path:    "/v3/users/:user_id/password",
		Handler: service.ChangePassword,
	},
	{
		Method:  "get",
		Path:    "/v3/users/:user_id/projects",
		Handler: service.GetUserProjects,
	},
}
