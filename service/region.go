package service

import (
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/region"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/policy"
	"work.ctyun.cn/git/GoStack/gostone/request"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

var (
	ignoreRegionField = []string{"extra"}
)

func GetAllRegions(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("list_regions", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	search := new(request.RegionSearch)
	getQuery(ctx.QueryParams(), search)
	var regions = region.FindAllRegion(*search, http.StatusBadRequest)
	response := utils.SetArrayLink(regions, utils.RegionPath, "regions", ignoreRegionField)
	return ctx.JSON(http.StatusOK, response)
}

func CreateRegion(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	var u request.CreateRegionRequest
	err := checkCreateRegionRequest(ctx, &u)
	if err != nil {
		return err
	}
	regionRequest := u.Region
	if regionRequest.Id == "" {
		regionRequest.Id = utils.GenerateUUID()
	}
	region.CheckRegionId(regionRequest.Id)
	if regionRequest.ParentRegionId != "" {
		region.FindRegionById(regionRequest.ParentRegionId, http.StatusConflict)
	}

	e := model.Region{
		Id:             regionRequest.Id,
		ParentRegionId: regionRequest.ParentRegionId,
		Description:    regionRequest.Description,
		Extra:          "{}",
	}
	if !policy.Check("create_region", auth, e) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	region.CreateRegion(&e)
	return ctx.JSON(http.StatusCreated, utils.SetSingleLink(e, utils.RegionPath, "region", ignoreRegionField))
}

func UpdateRegion(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	regionId := ctx.Param("region_id")
	var u request.UpdateRegionRequest
	err := checkUpdateRegionRequest(ctx, &u)
	if err != nil {
		return err
	}
	regionRequest := u.Region
	var old = region.FindRegionById(regionId, http.StatusNotFound)
	if regionRequest.Id != "" && regionRequest.Id != old.Id {
		region.CheckRegionId(regionRequest.Id)
		old.Id = regionRequest.Id
	}
	if regionRequest.Description != "" {
		old.Description = regionRequest.Description
	}
	if regionRequest.ParentRegionId != "" && regionRequest.ParentRegionId != old.ParentRegionId {
		region.FindRegionById(regionRequest.ParentRegionId, http.StatusNotFound)
		old.ParentRegionId = regionRequest.ParentRegionId
	}

	if !policy.Check("update_region", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	region.UpdateRegion(&old)
	return ctx.JSON(http.StatusOK, utils.SetSingleLink(old, utils.RegionPath, "region", ignoreRegionField))

}

func GetRegion(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	regionId := ctx.Param("region_id")
	var old = region.FindRegionById(regionId, http.StatusNotFound)
	if !policy.Check("get_region", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	return ctx.JSON(http.StatusOK, utils.SetSingleLink(old, utils.RegionPath, "region", ignoreRegionField))
}

func DeleteRegion(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	regionId := ctx.Param("region_id")
	var old = region.FindRegionById(regionId, http.StatusNotFound)
	if !policy.Check("delete_region", auth, old) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	region.DeleteRegion(regionId)
	return ctx.NoContent(http.StatusNoContent)
}

//检查创建节点请求体
func checkCreateRegionRequest(ctx echo.Context, request *request.CreateRegionRequest) error {
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

func checkUpdateRegionRequest(ctx echo.Context, request *request.UpdateRegionRequest) error {
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
