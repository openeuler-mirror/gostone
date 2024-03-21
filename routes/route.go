package routes

import "github.com/labstack/echo/v4"

//路由信息
type Route struct {
	Method  string
	Path    string
	Handler func(ctx echo.Context) error
}

//获取所有路由
func GetRoutes() []Route {
	var routes []Route
	for i := range versionRoutes {
		routes = append(routes, versionRoutes[i])
	}
	for i := range userRoutes {
		routes = append(routes, userRoutes[i])
	}
	for i := range tokenRoutes {
		routes = append(routes, tokenRoutes[i])
	}
	for i := range assignRoutes {
		routes = append(routes, assignRoutes[i])
	}
	for i := range domainRoutes {
		routes = append(routes, domainRoutes[i])
	}
	for i := range projectRoutes {
		routes = append(routes, projectRoutes[i])
	}
	for i := range endpointRoutes {
		routes = append(routes, endpointRoutes[i])
	}
	for i := range serviceRoutes {
		routes = append(routes, serviceRoutes[i])
	}
	for i := range regionRoutes {
		routes = append(routes, regionRoutes[i])
	}
	for i := range roleRoutes {
		routes = append(routes, roleRoutes[i])
	}
	for i := range checkRoutes {
		routes = append(routes, checkRoutes[i])
	}
	return routes
}
