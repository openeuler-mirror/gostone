package service

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/connect"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/domain"
	"work.ctyun.cn/git/GoStack/gostone/mapper/role"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/policy"
	"work.ctyun.cn/git/GoStack/gostone/request"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

var (
	ignoreRoleField = []string{"extra"}
)

func GetAllRoles(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("list_roles", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	search := new(request.RoleSearch)
	getQuery(ctx.QueryParams(), search)
	var roles = role.FindAllRole(*search, http.StatusBadRequest)
	data := make([]model.Role, 0)
	for _, r := range roles {
		SetNoneDomainId(&r)
		data = append(data, r)
	}
	if search.PageSize == 0 {
		search.PageSize = -1
	}
	search.PageNum = search.PageNum * search.PageSize
	if search.PageNum == 0 {
		search.PageNum = -1
	}
	count := role.FindAllRoleCount(*search, http.StatusBadRequest)
	response := utils.SetArrayLink(data, utils.RolePath, "roles", ignoreRoleField)
	response["count"] = count
	return ctx.JSON(http.StatusOK, response)
}

func SetNoneDomainId(role *model.Role) {
	if role.DomainId == "<<null>>" {
		role.DomainId = ""
	}
}

func CreateRole(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	var u request.CreateRoleRequest
	err := checkCreateRoleRequest(ctx, &u)
	if err != nil {
		return err
	}
	roleRequest := u.Role
	role.CheckRoleName(roleRequest.Name)
	if roleRequest.DomainId != "" {
		domain.FindDomainById(roleRequest.DomainId, http.StatusConflict)
	}
	if roleRequest.DomainId == "" {
		roleRequest.DomainId = "<<null>>"
	}
	e := model.Role{
		Id:       utils.GenerateUUID(),
		DomainId: roleRequest.DomainId,
		Name:     roleRequest.Name,
		Extra:    "{}",
	}
	if !policy.Check("create_role", auth, e) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	role.CreateRole(&e)
	SetNoneDomainId(&e)
	return ctx.JSON(http.StatusCreated, utils.SetSingleLink(e, utils.RolePath, "role", ignoreRoleField))
}

func UpdateRole(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	roleId := ctx.Param("role_id")
	var u request.UpdateRoleRequest
	err := checkUpdateRoleRequest(ctx, &u)
	if err != nil {
		return err
	}
	roleRequest := u.Role
	var old = role.FindRoleById(roleId, http.StatusNotFound)
	if roleRequest.DomainId == "" {
		roleRequest.DomainId = "<<null>>"
	}
	if roleRequest.DomainId != old.DomainId && roleRequest.DomainId != "<<null>>" {
		domain.FindDomainById(roleRequest.DomainId, http.StatusConflict)
	}
	old.DomainId = roleRequest.DomainId
	if roleRequest.Name != "" && roleRequest.Name != old.Name {
		var r model.Role
		ok := connect.GetMysqlConnect().Where("name=? ", roleRequest.Name).First(&r)
		if ok.Error == nil {
			panic(execption.NewGoStoneError(http.StatusConflict, "role name has  exists"))
		}
		old.Name = roleRequest.Name
	}

	if !policy.Check("update_role", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	role.UpdateRole(&old)
	if old.DomainId == "<<null>>" {
		old.DomainId = ""
	}
	return ctx.JSON(http.StatusOK, utils.SetSingleLink(old, utils.RolePath, "role", ignoreRoleField))

}

func GetRole(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	roleId := ctx.Param("role_id")
	var old = role.FindRoleById(roleId, http.StatusNotFound)
	if !policy.Check("get_role", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	if old.DomainId == "<<null>>" {
		old.DomainId = ""
	}
	return ctx.JSON(http.StatusOK, utils.SetSingleLink(old, utils.RolePath, "role", ignoreRoleField))
}

func DeleteRole(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	roleId := ctx.Param("role_id")
	var old = role.FindRoleById(roleId, http.StatusNotFound)
	if !policy.Check("delete_role", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	role.DeleteRole(roleId)
	return ctx.NoContent(http.StatusNoContent)
}

//检查创建节点请求体
func checkCreateRoleRequest(ctx echo.Context, request *request.CreateRoleRequest) error {
	if err := ctx.Bind(request); err != nil {
		log.Error(err)
		return err
	}
	if err := ctx.Validate(request); err != nil {
		log.Error(err)
		err = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return err
	}
	return nil
}

func checkUpdateRoleRequest(ctx echo.Context, request *request.UpdateRoleRequest) error {
	if err := ctx.Bind(request); err != nil {
		log.Error(err)
		return err
	}
	if err := ctx.Validate(request); err != nil {
		log.Error(err)
		err = echo.NewHTTPError(http.StatusBadRequest, err.Error())
		return err
	}
	return nil
}
