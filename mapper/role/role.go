package role

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/connect"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/common"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/request"
)

func FindRoleById(id string, code int) model.Role {
	var role model.Role
	err := connect.GetMysqlConnect().Table("role").Where("id =?", id).First(&role)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get role error id:"+id))
	}
	return role

}

func CheckRoleName(name string) {
	var role model.Role
	ok := connect.GetMysqlConnect().Table("role").Where("name=? ", name).First(&role)
	if ok.Error == nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "role name has  exists"))
	}
}

func FindRoleByIds(ids []interface{}) []model.Role {
	roles := make([]model.Role, 0)
	connect.GetMysqlConnect().Table("role").Where("id in ?", ids).Find(&roles)
	return roles
}

func CreateRole(role *model.Role) {
	ok := connect.GetMysqlConnect().Table("role").Create(&role)
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "role save error :"+ok.Error.Error()))
	}
}

func UpdateRole(role *model.Role) {
	ok := connect.GetMysqlConnect().Table("role").Where("id = ?", role.Id).Save(&role)
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "region update error :"+ok.Error.Error()))
	}
}

func DeleteRole(roleId string) {
	session := connect.GetMysqlConnect().Begin()
	err := session.Table("role").Where("id =?", roleId).Delete(&model.Role{})
	if err.Error != nil {
		session.Rollback()
		panic(execption.NewGoStoneError(http.StatusBadRequest, " delete role error:"+err.Error.Error()))
	}
	err = session.Table("assignment").Where("role_id = ?", roleId).Delete(&model.Assignment{})
	if err.Error != nil {
		session.Rollback()
		panic(execption.NewGoStoneError(http.StatusBadRequest, " delete role error:"+err.Error.Error()))
	}
	err = session.Commit()
	if err.Error != nil {
		panic(execption.NewGoStoneError(http.StatusBadRequest, " delete role error:"+err.Error.Error()))
	}

}

func FindRoleNameByUserIdAndProjectId(userId, projectId string, code int) []string {
	var result []string
	err := connect.GetMysqlConnect().Table("assignment").Model(&model.Assignment{}).Select("r.name").
		Joins("left JOIN role r on assignment.role_id=r.id").
		Where("actor_id = ?", userId).Where("target_id = ?", projectId).Find(&result)
	if err.Error != nil {
		log.Error(err)
		panic(execption.NewGoStoneError(code, "get role error:"+err.Error.Error()))
	}
	return result
}

func FindAllRole(search request.RoleSearch, code int) []model.Role {
	var roles []model.Role
	err := connect.GetMysqlConnect().Table("role").Select(" id,name,extra,CASE  WHEN domain_id='<<null>>' THEN       ''    ELSE        domain_id   END domain_id ").
		Scopes(common.FindByName(search.Name), common.FindByDomainId(search.DomainId)).
		Find(&roles)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get role error "))
	}
	return roles
}

func FindAllRoleCount(search request.RoleSearch, code int) int64 {
	var count int64
	err := connect.GetMysqlConnect().Table("role").
		Select(" id,name,extra,CASE  WHEN domain_id='<<null>>' THEN       ''    ELSE        domain_id   END domain_id ").
		Scopes(common.FindByName(search.Name), common.FindByDomainId(search.DomainId)).
		Limit(search.PageSize).Offset(search.PageNum).
		Count(&count)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get role error"))
	}
	return count
}
