package service

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/conf"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/project"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/policy"
	"work.ctyun.cn/git/GoStack/gostone/request"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

var (
	ignoreProjectField = []string{"extra"}
)

func GetProjectList(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("list_projects", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	search := new(request.ProjectSearch)
	getQuery(ctx.QueryParams(), search)
	if search.PageSize == 0 {
		search.PageSize = -1
	}
	search.PageNum = search.PageSize * search.PageNum
	if search.PageNum == 0 {
		search.PageNum = -1
	}
	count := project.FindAllProjectCount(*search, http.StatusBadRequest)
	var domains = project.FindAllProject(*search, http.StatusBadRequest)
	response := utils.SetArrayLink(domains, utils.ProjectPath, "projects", ignoreProjectField)
	response["count"] = count
	return ctx.JSON(http.StatusOK, response)
}

func CreateProject(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	var u request.CreateProjectRequest
	err := checkCreateProjectRequest(ctx, &u)
	projectRequest := u.Project
	if err != nil {
		return err
	}
	project.CheckProjectName(projectRequest.Name)
	if len(projectRequest.Name) > 64 {
		panic(execption.NewGoStoneError(http.StatusBadRequest, "domain name must less than 64"))
	}
	if projectRequest.DomainId == "" {
		projectRequest.DomainId = "default"
	}
	if projectRequest.ParentId == "" {
		projectRequest.ParentId = projectRequest.DomainId
	}
	project.FindProjectById(projectRequest.ParentId, http.StatusConflict)
	d := model.Project{
		Id:          utils.GenerateUUID(),
		Name:        projectRequest.Name,
		Description: projectRequest.Description,
		Extra:       "{}",
		IsDomain:    0,
		ParentId: sql.NullString{
			String: projectRequest.ParentId,
			Valid:  true,
		},
		Enabled:  getEnabled(projectRequest.Enabled),
		DomainId: projectRequest.DomainId,
	}
	if !policy.Check("create_project", auth, d) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	project.CreateProject(&d)
	return ctx.JSON(http.StatusCreated, utils.SetSingleLink(d, utils.ProjectPath, "project", ignoreProjectField))
}

func UpdateProject(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	projectId := ctx.Param("project_id")
	var u request.UpdateProjectRequest
	err := checkUpdateProjectRequest(ctx, &u)
	projectRequest := u.Project
	if err != nil {
		return err
	}
	var old = project.FindProjectById(projectId, http.StatusNotFound)
	if old.Name != projectRequest.Name {
		project.CheckProjectName(projectRequest.Name)
	}
	if len(projectRequest.Name) > 64 {
		panic(execption.NewGoStoneError(http.StatusConflict, "domain name must less than 64"))
	}
	utils.CopyProperties(&old, projectRequest)
	if projectRequest.Enabled != nil && *projectRequest.Enabled {
		old.Enabled = 1
	} else if projectRequest.Enabled != nil && !*projectRequest.Enabled {
		old.Enabled = 0
	}
	if projectRequest.ParentId != "" {
		project.FindProjectById(projectRequest.ParentId, http.StatusNotFound)
	}
	if !policy.Check("update_project", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	project.UpdateProject(&old)
	return ctx.JSON(http.StatusOK, utils.SetSingleLink(old, utils.ProjectPath, "project", ignoreDomainField))
}

func GetProject(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	projectId := ctx.Param("project_id")
	var old = project.FindProjectById(projectId, http.StatusNotFound)
	if !policy.Check("get_project", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	data := Struct2Map(old)
	parent := make([]interface{}, 0)
	if old.ParentId.Valid && old.ParentId.String != "default" {
		parent = getParents(parent, old.ParentId)
	}
	if len(parent) != 0 {
		data["parents"] = parent
	}
	child := make([]interface{}, 0)
	child = getChild(child, old.Id)
	if len(child) != 0 {
		data["subtree"] = child
	}
	return ctx.JSON(http.StatusOK, utils.SetSingleLink(data, utils.ProjectPath, "project", ignoreProjectField))
}

func getChild(projects []interface{}, id string) []interface{} {
	ps, err := project.FindAllProjectsByParentId(id)
	if err != nil || len(ps) == 0 {
		return projects
	}
	for _, p := range ps {
		data := Struct2Map(p)
		data["links"] = map[string]interface{}{
			"self": conf.Url + "/v3/projects/" + p.Id,
		}
		projects = append(projects, dealMap(data))
		getChild(projects, p.Id)
	}
	return projects
}

func getParents(projects []interface{}, parentId sql.NullString) []interface{} {
	if !parentId.Valid || parentId.String == "default" {
		return projects
	}
	p := project.FindProjectById(parentId.String, http.StatusNotFound)
	if p.IsDomain == 1 {
		return projects
	}
	data := Struct2Map(p)
	data["links"] = map[string]interface{}{
		"self": conf.Url + "/v3/projects/" + p.Id,
	}
	projects = append(projects, dealMap(data))
	return getParents(projects, p.ParentId)
}

func dealMap(target map[string]interface{}) map[string]interface{} {
	enable, ok := target["enabled"]
	if ok {
		switch enable.(type) {
		case int:
			if enable.(int) == 1 {
				target["enabled"] = true
			} else {
				target["enabled"] = false
			}
		case int64:
			if enable.(int64) == 1 {
				target["enabled"] = true
			} else {
				target["enabled"] = false
			}
		}
	}
	isDomain, ok := target["is_domain"]
	if ok {
		switch isDomain.(type) {
		case int:
			if isDomain.(int) == 1 {
				target["is_domain"] = true
			} else {
				target["is_domain"] = false
			}
		case int64:
			if isDomain.(int64) == 1 {
				target["is_domain"] = true
			} else {
				target["is_domain"] = false
			}
		}
	}
	parentId, ok := target["parent_id"]
	if ok {
		pd := parentId.(sql.NullString)
		target["parent_id"] = pd.String
	}
	if _, ok := target["extra"]; ok {
		delete(target, "extra")
	}
	return target
}

func DeleteProject(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	projectId := ctx.Param("project_id")
	var old = project.FindProjectById(projectId, http.StatusNotFound)
	if !policy.Check("delete_project", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	projects, _ := project.FindAllProjectsByParentId(projectId)
	if projects != nil && len(projects) != 0 {
		panic(execption.NewGoStoneError(http.StatusBadRequest, " can't delete has child project"))
	}
	project.DeleteProject(projectId)
	return ctx.NoContent(http.StatusNoContent)
}

//检查创建项目请求体
func checkCreateProjectRequest(ctx echo.Context, request *request.CreateProjectRequest) error {
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

//检查创建项目请求体
func checkUpdateProjectRequest(ctx echo.Context, request *request.UpdateProjectRequest) error {
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
