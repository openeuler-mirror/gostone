package service

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/endpoint"
	"work.ctyun.cn/git/GoStack/gostone/mapper/region"
	"work.ctyun.cn/git/GoStack/gostone/mapper/service"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/policy"
	"work.ctyun.cn/git/GoStack/gostone/request"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

var (
	ignoreEndpointField = []string{"extra"}
)

func GetAllEndpoints(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("list_endpoints", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	search := new(request.EndpointSearch)
	getQuery(ctx.QueryParams(), search)
	endpoints := endpoint.FindAllEndpoint(*search, http.StatusBadRequest)
	data := make([]model.EndpointResponse, 0)
	for _, end := range endpoints {
		e := getEndpointResponse(end)
		data = append(data, e)
	}
	response := utils.SetArrayLink(data, utils.EndpointPath, "endpoints", ignoreEndpointField)
	return ctx.JSON(http.StatusOK, response)
}
func getEndpointResponse(end model.Endpoint) model.EndpointResponse {
	return model.EndpointResponse{
		Id:               end.Id,
		LegacyEndpointId: end.LegacyEndpointId,
		Interface:        end.Interface,
		ServiceId:        end.ServiceId,
		Url:              end.URL,
		Extra:            end.Extra,
		Enabled:          transEnabled(end.Enabled),
		RegionId:         end.RegionId,
		Region:           end.RegionId,
	}
}
func CreateEndpoint(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	var u request.CreateEndpointRequest
	err := checkCreateEndpointRequest(ctx, &u)
	if err != nil {
		return err
	}
	endpointRequest := u.Endpoint
	service.FindServiceById(endpointRequest.ServiceId, http.StatusConflict)
	if endpointRequest.RegionId == "" && endpointRequest.Region == "" {
		panic(execption.NewGoStoneError(http.StatusConflict, "region must be present"))
	}
	if endpointRequest.RegionId == "" {
		endpointRequest.RegionId = endpointRequest.Region
	}
	region.FindRegionById(endpointRequest.RegionId, http.StatusConflict)
	e := model.Endpoint{
		Id:        utils.GenerateUUID(),
		Interface: endpointRequest.Interface,
		ServiceId: endpointRequest.ServiceId,
		URL:       endpointRequest.Url,
		Extra:     "{}",
		Enabled:   getEnabled(endpointRequest.Enabled),
		RegionId:  endpointRequest.RegionId,
	}
	if !policy.Check("create_endpoint", auth, e) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	endpoint.CreateEndpoint(&e)
	res := getEndpointResponse(e)
	return ctx.JSON(http.StatusCreated, utils.SetSingleLink(res, utils.EndpointPath, "endpoint", ignoreEndpointField))
}

func UpdateEndpoint(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	endpointId := ctx.Param("endpoint_id")
	var u request.UpdateEndpointRequest
	err := checkUpdateEndpointRequest(ctx, &u)
	if err != nil {
		return err
	}
	endpointRequest := u.Endpoint
	var old = endpoint.FindEndpointById(endpointId, http.StatusNotFound)
	if endpointRequest.ServiceId != "" && endpointRequest.ServiceId != old.ServiceId {
		service.FindServiceById(endpointRequest.ServiceId, http.StatusConflict)
		old.ServiceId = endpointRequest.ServiceId
	}
	if endpointRequest.RegionId == "" {
		endpointRequest.RegionId = endpointRequest.Region
	}
	if endpointRequest.RegionId != "" && endpointRequest.RegionId != old.RegionId {
		region.FindRegionById(endpointRequest.RegionId, http.StatusConflict)
		old.RegionId = endpointRequest.RegionId
	}
	if endpointRequest.Enabled != nil && *endpointRequest.Enabled {
		old.Enabled = 1
	} else if endpointRequest.Enabled != nil && !*endpointRequest.Enabled {
		old.Enabled = 0
	}
	if endpointRequest.Url != "" {
		old.URL = endpointRequest.Url
	}
	if endpointRequest.Interface != "" {
		old.Interface = endpointRequest.Interface
	}
	if !policy.Check("update_endpoint", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	endpoint.UpdateEndpoint(&old)
	res := getEndpointResponse(old)
	return ctx.JSON(http.StatusOK, utils.SetSingleLink(res, utils.EndpointPath, "endpoint", ignoreEndpointField))

}

func GetEndpoint(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	endpointId := ctx.Param("endpoint_id")
	var old = endpoint.FindEndpointById(endpointId, http.StatusNotFound)
	if !policy.Check("get_endpoint", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	res := getEndpointResponse(old)
	return ctx.JSON(http.StatusOK, utils.SetSingleLink(res, utils.EndpointPath, "endpoint", ignoreEndpointField))
}

func DeleteEndpoint(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	endpointId := ctx.Param("endpoint_id")
	var old = endpoint.FindEndpointById(endpointId, http.StatusNotFound)
	if !policy.Check("delete_endpoint", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	endpoint.DeleteEndpoint(endpointId)
	return ctx.NoContent(http.StatusNoContent)
}

//检查创建节点请求体
func checkCreateEndpointRequest(ctx echo.Context, request *request.CreateEndpointRequest) error {
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

func checkUpdateEndpointRequest(ctx echo.Context, request *request.UpdateEndpointRequest) error {
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
