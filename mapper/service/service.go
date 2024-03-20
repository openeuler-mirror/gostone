package service

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/connect"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/request"
)

func FindServiceList() []model.Service {
	var serviceList []model.Service
	ok := connect.GetMysqlConnect().Table("service").Find(&serviceList)
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "service create error:"+ok.Error.Error()))
	}
	return serviceList
}

func FindServiceById(serviceId string, code int) model.Service {
	var service model.Service
	err := connect.GetMysqlConnect().Table("service").
		Where("id = ?", serviceId).First(&service)
	if err.Error != nil {
		panic(execption.NewGoStoneError(code, "service  has not exists"))
	}
	return service
}

func CreateService(service *model.Service) {
	ok := connect.GetMysqlConnect().Table("service").Create(&service)
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "endpoint save error :"+ok.Error.Error()))
	}
}

func UpdateService(service *model.Service) {
	ok := connect.GetMysqlConnect().Table("service").Where("id = ?", service.Id).Save(&service)
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "service update error :"+ok.Error.Error()))
	}
}

func DeleteService(serviceId string) {
	ok := connect.GetMysqlConnect().Table("service").
		Where("id = ?", serviceId).Delete(&model.Service{})
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusBadRequest, " delete service error:"+ok.Error.Error()))
	}
}

func FindEnabledService() []model.Service {
	var service []model.Service
	err := connect.GetMysqlConnect().Table("service").Where("enabled=1").Find(&service)
	if err.Error != nil {
		panic(execption.NewGoStoneError(execption.StatusBadRequest, "get service error"))
	}
	return service
}

func FindAllService(search request.ServiceSearch, code int) []model.Service {
	var services []model.Service
	err := connect.GetMysqlConnect().Scopes(findByType(search.Type)).Table("service").Find(&services)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get service error "))
	}
	return services
}

func findByType(t string) func(db *gorm.DB) *gorm.DB {
	if t == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("1=1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("type=?", t)
	}
}
