package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"net/http"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/domain"
	"work.ctyun.cn/git/GoStack/gostone/mapper/user"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/policy"
	"work.ctyun.cn/git/GoStack/gostone/request"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

var (
	ignoreUserField = []string{"password", "created_at", "last_active_at", "extra", "password_sm3"}
)

func GetUsers(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("list_users", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	search := new(request.UserSearch)
	getQuery(ctx.QueryParams(), search)
	if search.PageSize == 0 {
		search.PageSize = -1
	}
	search.PageNum = search.PageSize * search.PageNum
	if search.PageNum == 0 {
		search.PageNum = -1
	}
	count := user.GetUserCount(*search, http.StatusNotFound)
	result := user.FindAllUser(*search, http.StatusNotFound)
	for _, r := range result {
		setExtra(r)
	}
	response := utils.SetArrayLink(result, utils.UserPath, "users", ignoreUserField)
	response["count"] = count
	return ctx.JSON(http.StatusOK, response)
}

func GetUser(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	userId := ctx.Param("user_id")
	search := request.UserSearch{
		UserId: userId,
	}
	result := user.FindAllUser(search, http.StatusNotFound)
	if len(result) == 0 {
		panic(execption.NewGoStoneError(http.StatusNotFound, "can not find user id:"+userId))
	}
	us := result[0]
	if !policy.Check("get_user", auth, us) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	setExtra(us)
	response := utils.SetSingleLink(us, utils.UserPath, "user", ignoreUserField)
	return ctx.JSON(http.StatusOK, response)
}

func setExtra(user map[string]interface{}) {
	extra := user["extra"].(string)
	if extra != "{}" {
		var extraMap map[string]interface{}
		utils.Byte2Struct([]byte(extra), &extraMap)
		if desc, ok := extraMap["description"]; ok {
			user["description"] = desc
		}
		if email, ok := extraMap["email"]; ok {
			user["email"] = email
		}
	}
}

func CreateUser(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("create_user", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	var u request.CreateUserRequest
	err := checkCreateUserRequest(ctx, &u)
	if err != nil {
		return err
	}
	userRequest := u.User
	if userRequest.DomainId == "" {
		userRequest.DomainId = "default"
	}
	domain.FindDomainById(userRequest.DomainId, http.StatusNotFound)
	user.CheckLocalUserName(userRequest.Name, userRequest.DomainId)
	userId := utils.GenerateUUID()
	u.User = userRequest
	user.CreateUser(u, userId)
	userMap := map[string]interface{}{
		"id":          userId,
		"description": userRequest.Description,
		"email":       userRequest.Email,
		"name":        userRequest.Name,
		"domain_id":   userRequest.DomainId,
	}
	if userRequest.Enabled == nil || *userRequest.Enabled {
		userMap["enabled"] = true
	} else if !*userRequest.Enabled {
		userMap["enabled"] = false
	}
	return ctx.JSON(http.StatusCreated, utils.SetSingleLink(userMap, utils.UserPath, "user", nil))
}

func UpdateUser(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	userId := ctx.Param("user_id")
	if !policy.Check("update_user", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	var u request.UpdateUserRequest
	err := checkUpdateUserRequest(ctx, &u)
	userRequest := u.User
	if err != nil {
		return err
	}
	var (
		us        model.User
		localUser model.LocalUser
	)
	us = user.GetUserById(userId, http.StatusNotFound)
	localUser = user.FindLocalUserByUserId(userId, http.StatusNotFound)
	if userRequest.Enabled != nil && *userRequest.Enabled {
		us.Enabled = true
	} else if userRequest.Enabled != nil && !*userRequest.Enabled {
		us.Enabled = false
	}
	if userRequest.Name != "" && userRequest.Name != localUser.Name {
		if userRequest.DomainId != "" {
			user.CheckLocalUserName(userRequest.Name, userRequest.DomainId)
		} else {
			user.CheckLocalUserName(userRequest.Name, localUser.DomainId)
		}
		localUser.Name = userRequest.Name
	}
	e := us.Extra
	var extraMap map[string]interface{}
	utils.Byte2Struct([]byte(e), &extraMap)
	if userRequest.Description != "" {
		extraMap["description"] = userRequest.Description
	}
	if userRequest.Email != "" {
		extraMap["email"] = userRequest.Email
	}
	us.Extra = utils.Struct2Json(extraMap)
	if userRequest.DomainId != "" {
		domain.FindDomainById(userRequest.DomainId, http.StatusNotFound)
		localUser.DomainId = userRequest.DomainId
		us.DomainId = userRequest.DomainId
	}
	user.UpdateUser(&us, &localUser)
	var extra map[string]interface{}
	utils.Byte2Struct([]byte(us.Extra), &extra)
	userMap := map[string]interface{}{
		"id":          userId,
		"description": extra["description"],
		"email":       extra["email"],
		"name":        localUser.Name,
		"domain_id":   us.DomainId,
		"enabled":     us.Enabled,
	}
	return ctx.JSON(http.StatusOK, utils.SetSingleLink(userMap, utils.UserPath, "user", []string{"extra"}))
}

func DeleteUser(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	userId := ctx.Param("user_id")
	if !policy.Check("delete_user", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	var (
		us        model.User
		localUser model.LocalUser
		password  model.Password
	)
	us = user.GetUserById(userId, http.StatusNotFound)
	localUser = user.FindLocalUserByUserId(userId, http.StatusNotFound)
	password = user.FindPasswordByLocalUserId(localUser.Id, http.StatusNotFound)
	user.DeleteUser(us, localUser, password)
	return ctx.NoContent(http.StatusNoContent)
}

func GetUserProjects(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("get_user_project", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}
	userId := ctx.Param("user_id")
	projects := user.FindProjectByUserId(userId, http.StatusNotFound)
	response := utils.SetArrayLink(projects, fmt.Sprintf(utils.UserProjectsPath, userId), "projects", ignoreProjectField)
	return ctx.JSON(http.StatusOK, response)
}

func ChangePassword(ctx echo.Context) error {
	userId := ctx.Param("user_id")
	var u request.ChangePasswordRequest
	err := checkChangePasswordRequest(ctx, &u)
	if err != nil {
		return err
	}
	passwordRequest := u.User
	var (
		localUser model.LocalUser
		password  model.Password
	)
	localUser = user.FindLocalUserByUserId(userId, http.StatusNotFound)
	password = user.FindPasswordByLocalUserId(localUser.Id, http.StatusNotFound)
	if !utils.CheckSM3Pwd(passwordRequest.OriginalPassword, password.PasswordSm3) &&
		!utils.CheckPwd(passwordRequest.OriginalPassword, password.PasswordHash) {
		panic(execption.NewGoStoneError(http.StatusUnauthorized, "original password not right"))
	}
	password.PasswordSm3 = utils.GenSM3Pwd(passwordRequest.Password)
	password.PasswordHash = utils.GetPwd(passwordRequest.Password)
	user.UpdatePassword(&password)
	return ctx.NoContent(http.StatusNoContent)
}

//检查创建宿主机请求体
func checkCreateUserRequest(ctx echo.Context, request *request.CreateUserRequest) error {
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

//检查创建用户请求体
func checkUpdateUserRequest(ctx echo.Context, request *request.UpdateUserRequest) error {
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

//检查更新用户请求体
func checkChangePasswordRequest(ctx echo.Context, request *request.ChangePasswordRequest) error {
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
