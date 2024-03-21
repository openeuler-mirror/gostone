package service

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/domain"
	"work.ctyun.cn/git/GoStack/gostone/mapper/project"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/policy"
	"work.ctyun.cn/git/GoStack/gostone/request"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

var (
	ignoreDomainField = []string{"extra"}
)

func GetDomainList(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("list_domains", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	search := new(request.DomainSearch)
	getQuery(ctx.QueryParams(), search)
	domains := domain.FindAllDomain(*search, http.StatusBadRequest)
	response := utils.SetArrayLink(domains, utils.DomainPath, "domains", ignoreDomainField)
	return ctx.JSON(http.StatusOK, response)
}

func CreateDomain(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	var u request.CreateDomainRequest
	err := checkCreateDomainRequest(ctx, &u)
	domainRequest := u.Domain
	if err != nil {
		return err
	}
	domain.CheckDomainName(domainRequest.Name, http.StatusConflict)
	if len(domainRequest.Name) > 64 {
		panic(execption.NewGoStoneError(http.StatusBadRequest, "domain name must less than 64"))
	}
	d := model.Project{
		Id:          utils.GenerateUUID(),
		Name:        domainRequest.Name,
		Description: domainRequest.Description,
		Extra:       "{}",
		IsDomain:    1,
		ParentId:    sql.NullString{},
		Enabled:     getEnabled(domainRequest.Enabled),
		DomainId:    "<<keystone.domain.root>>",
	}
	if !policy.Check("create_domain", auth, d) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	domain.CreateDomain(&d)
	return ctx.JSON(http.StatusCreated, utils.SetSingleLink(d, utils.DomainPath, "domain", ignoreDomainField))
}

func UpdateDomain(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	domainId := ctx.Param("domain_id")
	var u request.UpdateDomainRequest
	err := checkUpdateDomainRequest(ctx, &u)
	domainRequest := u.Domain
	if err != nil {
		return err
	}
	var old = domain.FindDomainById(domainId, http.StatusNotFound)
	if old.Name != domainRequest.Name {
		domain.CheckDomainName(domainRequest.Name, http.StatusConflict)
		old.Name = domainRequest.Name
	}
	if len(domainRequest.Name) > 64 {
		panic(execption.NewGoStoneError(http.StatusConflict, "domain name must less than 64"))
	}
	utils.CopyProperties(&old, domainRequest)
	if domainRequest.Enabled != nil && *domainRequest.Enabled {
		old.Enabled = 1
	} else if domainRequest.Enabled != nil && !*domainRequest.Enabled {
		old.Enabled = 0
	}
	if !policy.Check("update_domain", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	domain.UpdateDomain(&old, http.StatusConflict)
	return ctx.JSON(http.StatusOK, utils.SetSingleLink(old, utils.DomainPath, "domain", ignoreDomainField))
}

func GetDomain(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	domainId := ctx.Param("domain_id")
	var old = domain.FindDomainById(domainId, http.StatusNotFound)
	if !policy.Check("get_domain", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	return ctx.JSON(http.StatusOK, utils.SetSingleLink(old, utils.DomainPath, "domain", ignoreDomainField))
}

func DeleteDomain(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	domainId := ctx.Param("domain_id")
	var old = domain.FindDomainById(domainId, http.StatusNotFound)
	if !policy.Check("delete_domain", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	if old.Enabled == 1 {
		panic(execption.NewGoStoneError(http.StatusForbidden, " can't delete enabled domain"))
	}
	projects, _ := project.FindAllProjectsByParentId(domainId)
	if projects != nil && len(projects) != 0 {
		panic(execption.NewGoStoneError(http.StatusBadRequest, " can't delete has child domain"))
	}
	domain.DeleteDomain(domainId)
	return ctx.NoContent(http.StatusNoContent)
}

//检查创建项目请求体
func checkCreateDomainRequest(ctx echo.Context, request *request.CreateDomainRequest) error {
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
func checkUpdateDomainRequest(ctx echo.Context, request *request.UpdateDomainRequest) error {
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
