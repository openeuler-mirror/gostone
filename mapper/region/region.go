package region

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/connect"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/request"
)

func CheckRegionId(id string) {
	var region model.Region
	ok := connect.GetMysqlConnect().Table("region").Where("id=?", id).
		First(&region)
	if ok.Error == nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "region id has  exists"))
	}
}

func FindRegionById(id string, code int) model.Region {
	var region model.Region
	err := connect.GetMysqlConnect().Table("region").Where("id =?", id).
		First(&region)
	if err.Error != nil {
		panic(execption.NewGoStoneError(code, "region  has not exists"))
	}
	return region
}

func CreateRegion(region *model.Region) {
	ok := connect.GetMysqlConnect().Table("region").Create(&region)
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "region save error :"+ok.Error.Error()))
	}
}

func UpdateRegion(region *model.Region) {
	ok := connect.GetMysqlConnect().Table("region").
		Where("id = ?", region.Id).Save(&region)
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "region save error :"+ok.Error.Error()))
	}
}

func DeleteRegion(regionId string) {
	ok := connect.GetMysqlConnect().Table("region").Where("id = ?", regionId).Delete(&model.Region{})
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusBadRequest, " delete region error:"+ok.Error.Error()))
	}
}

func FindAllRegion(search request.RegionSearch, code int) []model.Region {
	var regions []model.Region
	err := connect.GetMysqlConnect().Scopes(findByParentRegionId(search.ParentRegionId)).Table("region").Find(&regions)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get region error "))
	}
	return regions
}

func findByParentRegionId(parentId string) func(db *gorm.DB) *gorm.DB {
	if parentId == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("1=1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("parent_region_id = ?", parentId)
	}
}
