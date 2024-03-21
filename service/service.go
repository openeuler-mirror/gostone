package service

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/service"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/policy"
	"work.ctyun.cn/git/GoStack/gostone/request"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

var (
	ignoreServiceField = []string{"extra"}
)

func GetAllServices(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("list_services", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	search := new(request.ServiceSearch)
	getQuery(ctx.QueryParams(), search)
	var services = service.FindAllService(*search, http.StatusBadRequest)
	data := make([]model.ServiceResponse, 0)
	for _, end := range services {
		res := getServiceResponse(end)
		data = append(data, res)
	}
	response := utils.SetArrayLink(data, utils.ServicePath, "services", ignoreServiceField)
	return ctx.JSON(http.StatusOK, response)
}

func getServiceResponse(s model.Service) model.ServiceResponse {
	var ex map[string]string
	utils.Byte2Struct([]byte(s.Extra), &ex)
	res := model.ServiceResponse{
		Id:   s.Id,
		Type: s.Type,
		Name: ex["name"],
	}
	if desc, ok := ex["description"]; ok {
		res.Description = desc
	}
	if s.Enabled == 0 {
		res.Enabled = false
	} else {
		res.Enabled = true
	}
	return res
}

func CreateService(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	var u request.CreateServiceRequest
	err := checkCreateServiceRequest(ctx, &u)
	if err != nil {
		return err
	}
	serviceRequest := u.Service
	var serviceList = service.FindServiceList()
	for _, s := range serviceList {
		var ex map[string]string
		utils.Byte2Struct([]byte(s.Extra), &ex)
		if ex["name"] == serviceRequest.Name {
			panic(execption.NewGoStoneError(http.StatusConflict, "service name has exists"))
		}
	}
	extra := map[string]interface{}{
		"name": serviceRequest.Name,
	}
	if serviceRequest.Description != "" {
		extra["description"] = serviceRequest.Description
	}
	e := model.Service{
		Id:      utils.GenerateUUID(),
		Extra:   utils.Struct2Json(extra),
		Enabled: getEnabled(serviceRequest.Enabled),
		Type:    serviceRequest.Type,
	}
	if !policy.Check("create_service", auth, e) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	service.CreateService(&e)
	res := getServiceResponse(e)
	return ctx.JSON(http.StatusCreated, utils.SetSingleLink(res, utils.ServicePath, "service", ignoreServiceField))
}

func UpdateService(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	serviceId := ctx.Param("service_id")
	var u request.UpdateServiceRequest
	err := checkUpdateServiceRequest(ctx, &u)
	if err != nil {
		return err
	}
	serviceRequest := u.Service
	var old = service.FindServiceById(serviceId, http.StatusNotFound)
	var ex map[string]string
	utils.Byte2Struct([]byte(old.Extra), &ex)
	if serviceRequest.Name != "" && serviceRequest.Name != ex["name"] {
		var serviceList = service.FindServiceList()
		for _, s := range serviceList {
			var ex map[string]string
			utils.Byte2Struct([]byte(s.Extra), &ex)
			if ex["name"] == serviceRequest.Name {
				panic(execption.NewGoStoneError(http.StatusConflict, "service name has exists"))
			}
		}
	}
	if serviceRequest.Description != "" {
		ex["description"] = serviceRequest.Description
	}
	if serviceRequest.Name != "" {
		ex["name"] = serviceRequest.Name
	}
	old.Extra = utils.Struct2Json(ex)

	if serviceRequest.Enabled != nil && *serviceRequest.Enabled {
		old.Enabled = 1
	} else if serviceRequest.Enabled != nil && !*serviceRequest.Enabled {
		old.Enabled = 0
	}
	if serviceRequest.Type != "" {
		old.Type = serviceRequest.Type
	}
	if !policy.Check("update_service", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	service.UpdateService(&old)
	res := getServiceResponse(old)
	return ctx.JSON(http.StatusOK, utils.SetSingleLink(res, utils.ServicePath, "service", ignoreServiceField))

}

func GetService(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	serviceId := ctx.Param("service_id")
	var old = service.FindServiceById(serviceId, http.StatusNotFound)
	if !policy.Check("get_service", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	res := getServiceResponse(old)
	return ctx.JSON(http.StatusOK, utils.SetSingleLink(res, utils.ServicePath, "service", ignoreServiceField))
}

func DeleteService(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	serviceId := ctx.Param("service_id")
	var old = service.FindServiceById(serviceId, http.StatusNotFound)
	if !policy.Check("delete_service", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	service.DeleteService(serviceId)
	return ctx.NoContent(http.StatusNoContent)
}

//检查创建节点请求体
func checkCreateServiceRequest(ctx echo.Context, request *request.CreateServiceRequest) error {
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

func checkUpdateServiceRequest(ctx echo.Context, request *request.UpdateServiceRequest) error {
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
