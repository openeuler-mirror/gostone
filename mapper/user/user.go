package user

import (
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
	"work.ctyun.cn/git/GoStack/gostone/connect"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/common"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/request"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

func FindUserById(id string, code int) model.UserInfo {
	var user model.UserInfo
	db := connect.GetMysqlConnect()
	err := db.Table("user u").Scopes(SelectUserInfo(), JoinLocalUserAndPassword()).Where("u.id=? ", id).Scan(&user)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get service error "))
	}
	return user
}

func FindProjectByUserId(userId string, code int) []model.Project {
	var projects []model.Project
	err := connect.GetMysqlConnect().Select("p.id as id,p.name as name, p.description as description,p.enabled as enabled, p.domain_id as domain_id ,p.parent_id as parent_id").
		Joins("left join project p on p.id=target_id and p.is_domain=0").
		Where("a.type=?", "UserProject").
		Where("a.actor_id=?", userId).
		Table("assignment a").Find(&projects)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get project error "))
	}
	return projects
}

func FindUserByNameAndDomainId(id, domainId string, code int) model.UserInfo {
	var user model.UserInfo
	err := connect.GetMysqlConnect().Select("u.id as id,name,enabled,u.domain_id,default_project_id,u.created_at,"+
		"last_active_at,p.password_hash as password ,p.password_sm3 as password_sm3,u.extra,l.id as local_user_id").
		Joins("left join local_user as l on u.id=l.user_id").
		Joins("left join password as p on l.id = p.local_user_id").
		Where("l.name=? ", id).
		Where("l.domain_id=? ", domainId).
		Where("ISNULL(p.expires_at)").
		Table("user u").Find(&user)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get service error "))
	}
	return user
}

func GetUserCount(search request.UserSearch, code int) int64 {
	var count int64
	err := connect.GetMysqlConnect().Table("user u").Scopes(SelectUserInfoWithoutPwd(), JoinLocalUser(),
		findByDomainId(search.DomainId), findById(search.UserId),
		common.FindByEnabled(search.Enabled), common.FindByName(search.Name)).Count(&count)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get service error "))
	}
	return count
}

func FindAllUser(search request.UserSearch, code int) []map[string]interface{} {
	var info []map[string]interface{}
	err := connect.GetMysqlConnect().Table("user u").Scopes(SelectUserInfoWithoutPwd(), JoinLocalUser(),
		findByDomainId(search.DomainId), findById(search.UserId),
		common.FindByEnabled(search.Enabled), common.FindByName(search.Name)).
		Limit(search.PageSize).Offset(search.PageNum).Find(&info)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get service error "))
	}
	return info
}

func findById(id string) func(db *gorm.DB) *gorm.DB {
	if id == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("1=1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("u.id =? ", id)
	}
}

func findByDomainId(id string) func(db *gorm.DB) *gorm.DB {
	if id == "" {
		return func(db *gorm.DB) *gorm.DB {
			return db.Where("1=1")
		}
	}
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("u.domain_id =? ", id)
	}
}
func SelectUserInfo() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select("u.id as id,name,enabled,u.domain_id,default_project_id,u.created_at," +
			"last_active_at,p.password_hash as password ,p.password_sm3 as password_sm3 ,u.extra,l.id as local_user_id")
	}
}

func SelectUserInfoWithoutPwd() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Select("u.id as id,name,enabled,u.domain_id,default_project_id,u.created_at," +
			"last_active_at,u.extra,l.id as local_user_id")
	}
}

func JoinLocalUser() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins("left join local_user as l on u.id=l.user_id")
	}
}

func JoinLocalUserAndPassword() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Joins("left join local_user as l on u.id=l.user_id").
			Joins("left join password as p on l.id = p.local_user_id")
	}
}

func GetUserById(id string, code int) model.User {
	var user model.User
	err := connect.GetMysqlConnect().Table("user").Where("id =?", id).First(&user)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(code, "get user error,id:"+id))
	}
	return user
}

func FindLocalUserByUserId(id string, code int) model.LocalUser {
	var localUser model.LocalUser
	ok := connect.GetMysqlConnect().Table("local_user").
		Where("user_id=?", id).First(&localUser)
	if ok.Error != nil {
		panic(execption.NewGoStoneError(code, "user is not exist"))
	}
	return localUser
}

func GetLocalUserByNameAndDomainId(name string, domainId string) model.LocalUser {
	var user model.LocalUser
	err := connect.GetMysqlConnect().Where("name=?", name).Where("domain_id=?", domainId).Find(&user)
	if err.Error != nil {
		log.Error(err)
		panic(execption.NewGoStoneError(execption.StatusBadRequest, "get user error name:"+name+" domainId:"+domainId))
	}
	return user
}

func GetLocalUserById(id string) model.LocalUser {
	var user model.LocalUser
	err := connect.GetMysqlConnect().Where("user_id=?", id).Find(&user)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(execption.StatusBadRequest, "get local_user error id:"+id))
	}
	return user
}

func GetPasswordByLocalUserId(id int) model.Password {
	var user model.Password
	err := connect.GetMysqlConnect().Where("local_user_id=?", id).Find(&user)
	if err.Error != nil {
		log.Error(err.Error)
		panic(execption.NewGoStoneError(execption.StatusBadRequest, "get password error local_user_id:"+strconv.Itoa(id)))
	}
	return user
}

func CheckLocalUserName(name string, domainId string) {
	var lu model.LocalUser
	ok := connect.GetMysqlConnect().Where("name=?", name).
		Where("domain_id=?", domainId).Table("local_user").First(&lu)
	if ok.Error == nil {
		panic(execption.NewGoStoneError(http.StatusNotFound, "user name has exist"))
	}
}

func CreateUser(request request.CreateUserRequest, userId string) {
	userRequest := request.User
	err := connect.GetMysqlConnect().Transaction(func(tx *gorm.DB) error {
		baseUser := model.User{
			Id:        userId,
			CreatedAt: time.Now(),
			DomainId:  userRequest.DomainId,
		}
		if userRequest.Enabled == nil || *userRequest.Enabled {
			baseUser.Enabled = true
		} else if !*userRequest.Enabled {
			baseUser.Enabled = false
		}

		extra := make(map[string]interface{})
		if userRequest.Description != "" {
			extra["description"] = userRequest.Description
		}
		if userRequest.Email != "" {
			extra["email"] = userRequest.Email
		}
		baseUser.Extra = utils.Struct2Json(extra)
		ok := tx.Table("user").Omit("LastActiveAt").Create(&baseUser)
		if ok.Error != nil {
			return ok.Error
		}

		localUser := model.LocalUser{
			UserId:   userId,
			DomainId: userRequest.DomainId,
			Name:     userRequest.Name,
		}
		ok = tx.Table("local_user").Omit("FailedAuthAt").Create(&localUser)
		if ok.Error != nil {
			return ok.Error
		}
		password := model.Password{
			LocalUserId:  localUser.Id,
			PasswordSm3:  utils.GenSM3Pwd(userRequest.Password),
			SelfService:  false,
			PasswordHash: utils.GetPwd(userRequest.Password),
			CreatedAtInt: time.Now().Unix(),
			CreatedAt:    time.Now(),
		}
		ok = tx.Table("password").Omit("ExpiresAt").Create(&password)
		if ok.Error != nil {

			return ok.Error
		}
		return nil
	})
	if err != nil {
		panic(execption.NewGoStoneError(execption.StatusConflict, "create User error::"+err.Error()))
	}
}

func UpdateUser(us *model.User, localUser *model.LocalUser) {
	session := connect.GetMysqlConnect().Begin()
	ok := session.Omit("LastActiveAt").Table("user").Save(us)
	if ok.Error != nil {
		session.Rollback()
		panic(execption.NewGoStoneError(http.StatusConflict, ok.Error.Error()))
	}
	ok = session.Table("local_user").Omit("FailedAuthAt").Save(localUser)
	if ok.Error != nil {
		session.Rollback()
		panic(execption.NewGoStoneError(http.StatusConflict, ok.Error.Error()))
	}
	_ = session.Commit()
}

func FindPasswordByLocalUserId(id int, code int) model.Password {
	var password model.Password
	err := connect.GetMysqlConnect().Table("password").
		Where("local_user_id=?", id).First(&password)
	if err.Error != nil {
		panic(execption.NewGoStoneError(code, "password is not exist"))
	}
	return password
}

func DeleteUser(user model.User, localUser model.LocalUser, password model.Password) {
	session := connect.GetMysqlConnect().Begin()
	ok := session.Table("user").Where("id = ?", user.Id).Delete(&model.User{})
	if ok.Error != nil {
		session.Rollback()
		panic(execption.NewGoStoneError(http.StatusConflict, ok.Error.Error()))
	}
	ok = session.Table("local_user").Where("id = ?", localUser.Id).
		Delete(&model.LocalUser{}, localUser.Id)
	if ok.Error != nil {
		session.Rollback()
		panic(execption.NewGoStoneError(http.StatusConflict, ok.Error.Error()))
	}
	ok = session.Table("password").Where("id =?", password.Id).
		Delete(&model.Password{})
	if ok.Error != nil {
		session.Rollback()
		panic(execption.NewGoStoneError(http.StatusConflict, ok.Error.Error()))
	}
	ok = session.Table("assignment").Where(" actor_id=?", user.Id).
		Delete(model.Assignment{})
	if ok.Error != nil {
		session.Rollback()
		panic(execption.NewGoStoneError(http.StatusConflict, ok.Error.Error()))
	}
	_ = session.Commit()
}

func UpdatePassword(password *model.Password) {
	ok := connect.GetMysqlConnect().Table("password").
		Where("id = ?", password.Id).Omit("ExpiresAt", "CreatedAt").Save(&password)
	if ok.Error != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, ok.Error.Error()))
	}
}
