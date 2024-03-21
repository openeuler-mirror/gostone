package endpoint

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/connect"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/request"
)

func GetEndpointsByServiceId(serverId string) []model.Endpoint {
	var endpoints []model.Endpoint
	err := connect.GetMysqlConnect().Table("endpoint").Where("service_id=?", serverId).Find(&endpoints)
	if err.Error != nil {
		log.Error(err)
		panic(execption.NewGoStoneError(execption.StatusBadRequest, "get endpoint error serviceId:"+serverId))
	}
	return endpoints
}

func CreateEndpoint(endpoint *model.Endpoint) {
	ok := connect.GetMysqlConnect().Table("endpoint").Create(endpoint)
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "endpoint save error :"+ok.Error.Error()))
	}
}

func FindEndpointById(id string, code int) model.Endpoint {
	var endpoint model.Endpoint
	ok := connect.GetMysqlConnect().Table("endpoint").Where("id = ?", id).First(&endpoint)
	if ok.Error != nil {
		panic(execption.NewGoStoneError(code, "endpoint not found"))
	}
	return endpoint
}

func UpdateEndpoint(endpoint *model.Endpoint) {
	err := connect.GetMysqlConnect().Table("endpoint").Where("id =?", endpoint.Id).Save(&endpoint)
	if err.Error != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "endpoint save error :"+err.Error.Error()))
	}
}
func DeleteEndpoint(endpointId string) {
	err := connect.GetMysqlConnect().Table("endpoint").
		Where("id =?", endpointId).Delete(&model.Endpoint{})
	if err.Error != nil {
		panic(execption.NewGoStoneError(http.StatusBadRequest, " delete endpoint error:"+err.Error.Error()))
	}
}

func FindAllEndpoint(search request.EndpointSearch, code int) []model.Endpoint {
	var endpoints []model.Endpoint
	err := connect.GetMysqlConnect().Scopes(findByInterface(search.Interface),
		findByServiceId(search.ServiceId)).Table("endpoint").Find(&endpoints)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get endpoint error "))
	}
	return endpoints
}

func findByInterface(interfaceId string) func(db *gorm.DB) *gorm.DB {
	if interfaceId == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("1=1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("interface = ?", interfaceId)
	}
}

func findByServiceId(serviceId string) func(db *gorm.DB) *gorm.DB {
	if serviceId == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("1=1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("service_id = ?", serviceId)
	}
}
