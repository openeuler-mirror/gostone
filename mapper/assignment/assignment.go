package assignment

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/connect"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/common"
	"work.ctyun.cn/git/GoStack/gostone/model"
)

func FindAssignmentByProjectIdAndUserId(projectId string, userId string) []model.Assignment {
	assigns := make([]model.Assignment, 0)
	err := connect.GetMysqlConnect().Table("assignment").
		Where("target_id=?", projectId).
		Where("actor_id=?", userId).
		Where("type = ?", "UserProject").
		Find(&assigns)
	if err.Error != nil {
		log.Error(err.Error)
		return assigns
	}
	return assigns
}

func FindRoleByProjectIdAndUserId(projectId string, userId string) []Role {
	var roles []Role
	err := connect.GetMysqlConnect().Select("id , name ").Table("role r").
		Joins("left join assignment a on r.id=a.role_id").
		Where("a.target_id=?", projectId).
		Where("a.actor_id=?", userId).
		Where("type = ?", "UserProject").
		Find(&roles)
	if err.Error != nil {
		panic(execption.NewGoStoneError(http.StatusUnauthorized, "get role error:"+err.Error.Error()))
	}
	return roles
}

type Role struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func FindAssignmentByRoleIdAndActorIdAndTargetId(roleId, actorId, targetId string, code int) []model.Assignment {
	var assigns []model.Assignment
	err := connect.GetMysqlConnect().Table("assignment").
		Scopes(common.FindByRoleId(roleId), common.FindByActorId(actorId), common.FindByTargetId(targetId)).Find(&assigns)
	if err.Error != nil {
		log.Error(err)
		panic(execption.NewGoStoneError(code, "get assignment error:"+err.Error.Error()))
	}
	return assigns
}

func FindAssignmentByUserId(userId string) []model.Assignment {
	var assigns []model.Assignment
	err := connect.GetMysqlConnect().Table("assignment").
		Where("actor_id=?", userId).Find(&assigns)
	if err.Error != nil {
		log.Error(err)
		return nil
	}
	return assigns
}

func FindAssignmentByRoleIdAndActorIdAndTargetIdNoError(roleId, actorId, targetId string) []model.Assignment {
	var assigns []model.Assignment
	err := connect.GetMysqlConnect().Table("assignment").Where("role_id=?", roleId).
		Where("actor_id=?", actorId).Where("target_id=?", targetId).Find(&assigns)
	if err.Error != nil {
		log.Error(err)
		return nil
	}
	return assigns
}

func SaveAssignment(assign *model.Assignment) {
	err := connect.GetMysqlConnect().Table("assignment").Create(&assign)
	if err.Error != nil {
		panic(execption.NewGoStoneError(http.StatusBadRequest, "save assignment error:"+err.Error.Error()))
	}
}

func DeleteAssignment(roleId, actorId, targetId string) {
	connect.GetMysqlConnect().Table("assignment").Where("role_id=?", roleId).
		Where("actor_id=?", actorId).Where("target_id=?", targetId).Delete(&model.Assignment{})
}
