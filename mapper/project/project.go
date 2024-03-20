package project

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/connect"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/common"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/request"
)

func FindProjectById(id string, code int) model.Project {
	var project model.Project
	result := connect.GetMysqlConnect().Table("project").Where("id =?", id).First(&project)
	if result.Error != nil {
		log.Error(result.Error)
		panic(execption.NewGoStoneError(code, "get project by id error id:"+id))
	}
	return project
}

func FindProjectByNameAndDomainId(name string, domainId string, code int) model.Project {
	if domainId == "" {
		domainId = "default"
	}
	var project model.Project
	err := connect.GetMysqlConnect().Table("project").Where("name=?", name).Where("domain_id=?", domainId).First(&project)
	if err.Error != nil {
		log.Error(err)
		panic(execption.NewGoStoneError(code, "get project by id and domain_id error name:"+name+" domain_id:"+domainId))
	}
	return project
}

func CheckProjectName(name string) {
	var project model.Project
	ok := connect.GetMysqlConnect().Table("project").Where("name=? and is_domain=0", name).First(&project)
	if ok.Error == nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "project name has exists"))
	}
}

func CreateProject(project *model.Project) {
	ok := connect.GetMysqlConnect().Table("project").Create(&project)
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "project save error :"+ok.Error.Error()))
	}
}

func UpdateProject(project *model.Project) {
	ok := connect.GetMysqlConnect().Table("project").Where("id =?", project.Id).
		Save(&project)
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "project save error :"+ok.Error.Error()))
	}
}

func FindAllProjectsByParentId(parentId string) ([]model.Project, error) {
	var projects []model.Project
	err := connect.GetMysqlConnect().Table("project").Where("parent_id=?", parentId).Find(&projects)
	if err.Error != nil {
		log.Error(err.Error)
		return nil, err.Error
	}
	return projects, nil
}

func DeleteProject(projectId string) {
	session := connect.GetMysqlConnect().Begin()
	err := session.Table("project").Where("id=?", projectId).Delete(&model.Project{})
	if err.Error != nil {
		session.Rollback()
		panic(execption.NewGoStoneError(http.StatusBadRequest, " delete domain error:"+err.Error.Error()))
	}
	err = session.Table("assignment").Where("target_id=?", projectId).Delete(&model.Assignment{})
	if err.Error != nil {
		session.Rollback()
		panic(execption.NewGoStoneError(http.StatusBadRequest, " delete domain error:"+err.Error.Error()))
	}
	err = session.Commit()
	if err.Error != nil {
		session.Rollback()
		panic(execption.NewGoStoneError(http.StatusBadRequest, " delete domain error:"+err.Error.Error()))
	}
}

func FindAllProjectCount(search request.ProjectSearch, code int) int64 {
	var count int64
	err := connect.GetMysqlConnect().Scopes(common.FindByName(search.Name), common.FindByEnabled(search.Enabled),
		common.FindByDomainId(search.DomainId), findByIsDomain(search.IsDomain), findByParentId(search.ParentId)).
		Limit(search.PageSize).Offset(search.PageNum).
		Table("project").Count(&count)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get project error "))
	}
	return count
}
func FindAllProject(search request.ProjectSearch, code int) []model.Project {
	var projects []model.Project
	err := connect.GetMysqlConnect().Scopes(common.FindByName(search.Name), common.FindByEnabled(search.Enabled),
		common.FindByDomainId(search.DomainId), findByIsDomain(search.IsDomain), findByParentId(search.ParentId)).
		Table("project").Find(&projects)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get project error "))
	}
	return projects
}

func findByIsDomain(isDomain string) func(db *gorm.DB) *gorm.DB {
	if isDomain == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("is_domain = 0")
		}
	}
	if isDomain == "true" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("is_domain = 1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("is_domain = 0")
	}
}

func findByParentId(parentId string) func(db *gorm.DB) *gorm.DB {
	if parentId == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("1=1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("parent_id = ?", parentId)
	}
}
