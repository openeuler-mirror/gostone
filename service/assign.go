package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/conf"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/assignment"
	"work.ctyun.cn/git/GoStack/gostone/mapper/project"
	"work.ctyun.cn/git/GoStack/gostone/mapper/role"
	"work.ctyun.cn/git/GoStack/gostone/mapper/user"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/policy"
	"work.ctyun.cn/git/GoStack/gostone/request"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

func FindRoleListByProjectAndUser(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("list_grants", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	projectId := ctx.Param("project_id")
	userId := ctx.Param("user_id")
	assigns := assignment.FindAssignmentByProjectIdAndUserId(projectId, userId)
	if len(assigns) == 0 {
		return ctx.JSON(http.StatusOK, assigns)
	}
	roleIds := make([]interface{}, 0)
	for _, a := range assigns {
		roleIds = append(roleIds, a.RoleId)
	}
	roles := role.FindRoleByIds(roleIds)
	baseUrl := conf.Url + "/v3/projects/%s/users/%s/roles"
	roleMaps := make([]map[string]interface{}, 0)
	for _, r := range roles {
		roleMaps = append(roleMaps, map[string]interface{}{
			"id":       r.Id,
			"name":     r.Name,
			"domainId": r.DomainId,
			"links": map[string]interface{}{
				"self": conf.Url + "/identity/v3/roles/" + r.Id,
			},
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"roles": roleMaps,
		"links": map[string]interface{}{
			"previous": nil,
			"next":     nil,
			"self":     fmt.Sprintf(baseUrl, projectId, userId),
		},
	})
}

func GetAllAssignment(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("list_grants", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	search := new(request.AssignmentSearch)
	getQuery(ctx.QueryParams(), search)
	var assigns = assignment.FindAssignmentByRoleIdAndActorIdAndTargetId(search.RoleId, search.ActorId,
		search.TargetId, http.StatusBadRequest)
	assignMaps := make([]map[string]interface{}, 0)
	baseUrl := conf.Url + "/v3/projects/%s/users/%s/roles"
	for _, ass := range assigns {
		url := fmt.Sprintf(baseUrl, ass.TargetId, ass.ActorId)
		var r = role.FindRoleById(ass.RoleId, http.StatusNotFound)
		if r.DomainId != "<<null>>" {
			continue
		}
		assignMap := map[string]interface{}{
			"links": map[string]interface{}{
				"assignment": url,
			},
			"scope": map[string]interface{}{
				"project": map[string]interface{}{
					"id": ass.TargetId,
				},
			},
			"user": map[string]interface{}{
				"id": ass.ActorId,
			},
			"role": map[string]interface{}{
				"id":   r.Id,
				"name": r.Name,
			},
		}
		assignMaps = append(assignMaps, assignMap)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"role_assignments": assignMaps,
		"links": map[string]interface{}{
			"previous": nil,
			"next":     nil,
			"self":     conf.Url + "/v3/role_assignments",
		},
	})
}

func SaveRole2UserOnProjectName(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	projectName := ctx.Param("project_name")
	userId := ctx.Param("user_id")
	roleId := ctx.Param("role_id")
	p := project.FindProjectByNameAndDomainId(projectName, auth.DomainId, http.StatusNotFound)
	user.GetUserById(userId, http.StatusNotFound)
	role.FindRoleById(roleId, http.StatusNotFound)
	assign := model.Assignment{
		Type:     "UserProject",
		TargetId: p.Id,
		ActorId:  userId,
		RoleId:   roleId,
	}
	if !policy.Check("create_grant", auth, assign) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	//hasAss := assignment.FindAssignmentByUserId(userId)
	//for _, a := range hasAss {
	//	if a.TargetId != p.Id {
	//		panic(execption.NewGoStoneError(http.StatusForbidden, "user can't in multi project"))
	//	}
	//}
	ass := assignment.FindAssignmentByRoleIdAndActorIdAndTargetIdNoError(roleId, userId, p.Id)
	if ass != nil && len(ass) != 0 {
		return ctx.NoContent(http.StatusNoContent)
	}
	assignment.SaveAssignment(&assign)
	return ctx.NoContent(http.StatusNoContent)
}

func SaveRole2UserOnProject(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	projectId := ctx.Param("project_id")
	userId := ctx.Param("user_id")
	roleId := ctx.Param("role_id")
	project.FindProjectById(projectId, http.StatusNotFound)
	user.GetUserById(userId, http.StatusNotFound)
	role.FindRoleById(roleId, http.StatusNotFound)
	assign := model.Assignment{
		Type:     "UserProject",
		TargetId: projectId,
		ActorId:  userId,
		RoleId:   roleId,
	}
	if !policy.Check("create_grant", auth, assign) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	//hasAss := assignment.FindAssignmentByUserId(userId)
	//for _, a := range hasAss {
	//	if a.TargetId != projectId {
	//		panic(execption.NewGoStoneError(http.StatusForbidden, "user can't in multi project"))
	//	}
	//}
	ass := assignment.FindAssignmentByRoleIdAndActorIdAndTargetIdNoError(roleId, userId, projectId)
	if ass != nil && len(ass) != 0 {
		return ctx.NoContent(http.StatusNoContent)
	}
	assignment.SaveAssignment(&assign)
	return ctx.NoContent(http.StatusNoContent)
}

func ValidRole2UserOnProject(ctx echo.Context) error {
	projectId := ctx.Param("project_id")
	userId := ctx.Param("user_id")
	roleId := ctx.Param("role_id")
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("check_grant", auth, map[string]interface{}{
		"targetId": projectId,
		"actorId":  userId,
		"roleId":   roleId,
	}) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	assign := assignment.FindAssignmentByRoleIdAndActorIdAndTargetIdNoError(roleId, userId, projectId)
	if assign != nil && len(assign) != 0 {
		return ctx.NoContent(http.StatusNoContent)
	}
	return ctx.JSON(http.StatusNotFound, map[string]interface{}{})
}

func DeleteRole2UserOnProject(ctx echo.Context) error {
	projectId := ctx.Param("project_id")
	userId := ctx.Param("user_id")
	roleId := ctx.Param("role_id")
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("check_grant", auth, map[string]interface{}{
		"targetId": projectId,
		"actorId":  userId,
		"roleId":   roleId,
	}) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	assignment.DeleteAssignment(roleId, userId, projectId)
	return ctx.NoContent(http.StatusNoContent)
}
