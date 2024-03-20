package domain

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"work.ctyun.cn/git/GoStack/gostone/connect"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/common"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/request"
)

//为了防止过多的default domain查询在运行前自动缓存default
var defaultDomain model.Project

func InitDomain() {
	connect.GetMysqlConnect().Table("project").Where("name=?", "default").Where("is_domain=?", 1).Find(&defaultDomain)
}

func FindDomainById(id string, code int) model.Project {
	var domain model.Project
	if id == defaultDomain.Id {
		return defaultDomain
	}
	result := connect.GetMysqlConnect().Table("project").
		Where("id =?", id).First(&domain)
	if result.Error != nil {
		log.Error(result.Error)
		panic(execption.NewGoStoneError(code, "get domain by id error domain id:"+id))
	}
	return domain
}

func FindDomainByName(name string, code int) model.Project {
	var domain model.Project
	if strings.ToLower(name) == strings.ToLower(defaultDomain.Name) {
		return defaultDomain
	}
	result := connect.GetMysqlConnect().Table("project").Where("name=? and is_domain=1", name).First(&domain)
	if result.Error != nil {
		log.Error(result.Error)
		panic(execption.NewGoStoneError(code, "get domain by name error name:"+name))
	}
	return domain
}

func CheckDomainName(name string, code int) {
	var domain model.Project
	result := connect.GetMysqlConnect().Table("project").Where("name=? and is_domain=1", name).First(&domain)
	if result.Error == nil {
		panic(execption.NewGoStoneError(code, "domain name has exists name:"+name))
	}
}

func CreateDomain(domain *model.Project) {
	result := connect.GetMysqlConnect().Table("project").Create(domain)
	if result.Error != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "domain save error :"+result.Error.Error()))
	}
}

func UpdateDomain(domain *model.Project, code int) {
	result := connect.GetMysqlConnect().Table("project").
		Where("id=? ", domain.Id).Save(domain)
	if result.Error != nil {
		panic(execption.NewGoStoneError(code, "domain save error:"+result.Error.Error()))
	}
}

func DeleteDomain(domainId string) {
	ok := connect.GetMysqlConnect().Table("project").
		Where("id=? ", domainId).Delete(&model.Project{})
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusBadRequest, " delete domain error:"+ok.Error.Error()))
	}
}

func FindAllDomain(search request.DomainSearch, code int) []model.Project {
	var projects []model.Project
	err := connect.GetMysqlConnect().
		Where("is_domain=1").Scopes(common.FindByName(search.Name), common.FindByEnabled(search.Enabled)).
		Table("project").Find(&projects)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get domain error "))
	}
	return projects
}
