package routes

import "work.ctyun.cn/git/GoStack/gostone/service"

var tokenRoutes = []Route{
	//获取token
	{
		Method:  "post",
		Path:    "/v3/auth/tokens",
		Handler: service.IssueToken,
	},
	//验证Token
	{
		Method:  "get",
		Path:    "/v3/auth/tokens",
		Handler: service.ValidateToken,
	},
	{
		Method:  "head",
		Path:    "/v3/auth/tokens",
		Handler: service.ValidateToken,
	},
	//获取指定用户token
	{
		Method:  "post",
		Path:    "/v3/admin/tokens",
		Handler: service.GetUserTokenByAdmin,
	},
}
