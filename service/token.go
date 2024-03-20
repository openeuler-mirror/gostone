package service

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/muesli/cache2go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/assignment"
	"work.ctyun.cn/git/GoStack/gostone/mapper/domain"
	"work.ctyun.cn/git/GoStack/gostone/mapper/endpoint"
	"work.ctyun.cn/git/GoStack/gostone/mapper/project"
	"work.ctyun.cn/git/GoStack/gostone/mapper/service"
	"work.ctyun.cn/git/GoStack/gostone/mapper/user"
	"work.ctyun.cn/git/GoStack/gostone/model"
	"work.ctyun.cn/git/GoStack/gostone/policy"
	"work.ctyun.cn/git/GoStack/gostone/request"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

func IssueToken(ctx echo.Context) error {
	var identityInfo map[string]interface{}
	if err := ctx.Bind(&identityInfo); err != nil {
		log.Error(err)
		return err
	}
	var auth Auth
	var defaultDomain *model.Project
	err := utils.Map2Structure(identityInfo["auth"], &auth)
	if err != nil {
		panic(execption.NewGoStoneError(http.StatusConflict, "issue token error"))
	}
	method := auth.Identity.Methods
	if len(method) == 0 {
		panic(execption.NewGoStoneError(http.StatusConflict, "issue token method error"))
	}
	var (
		p  model.Project
		u  model.UserInfo
		pd model.Project
		ud model.Project
	)
	if auth.Scope.Project.Id == "" && auth.Scope.Project.Name == "" {
		panic(execption.NewGoStoneError(http.StatusBadRequest, " project id/name must be present"))
	}
	if auth.Scope.Project.Name == "" {
		p = project.FindProjectById(auth.Scope.Project.Id, http.StatusBadRequest)
	} else {
		p = project.FindProjectByNameAndDomainId(auth.Scope.Project.Name, auth.Scope.Project.Domain.Id, http.StatusBadRequest)
	}
	if auth.Scope.Project.Domain.Id == "" && auth.Scope.Project.Domain.Name == "" {
		if p.DomainId == "" {
			panic(execption.NewGoStoneError(http.StatusBadRequest, " project domain id/name must be present"))
		}
		auth.Scope.Project.Domain.Id = p.DomainId
	}
	if auth.Scope.Project.Domain.Name != "" {
		d2 := domain.FindDomainByName(auth.Scope.Project.Domain.Name, http.StatusBadRequest)
		if auth.Scope.Project.Domain.Name == "default" {
			defaultDomain = &d2
		}
		pd = d2
		auth.Scope.Project.Domain.Id = d2.Id
	} else {
		pd = domain.FindDomainById(auth.Scope.Project.Domain.Id, http.StatusBadRequest)
		if auth.Scope.Project.Domain.Id == "default" {
			defaultDomain = &pd
		}

	}
	tokenValid := utils.GetTokenMethod(utils.TokenType)
	switch method[0] {
	case "token":
		token := auth.Identity.Token
		if token.Id == "" {
			panic(execption.NewGoStoneError(http.StatusBadRequest, "token  is null"))
		}
		authContext := tokenValid.Validate(token.Id)
		result := user.FindUserById(authContext.UserId, http.StatusBadRequest)
		u = result
		if u.Enabled == 0 {
			panic(execption.NewGoStoneError(http.StatusUnauthorized, "user is disabled"))
		}
		if authContext.DomainId == "default" {
			if defaultDomain != nil {
				ud = *defaultDomain
			} else {
				ud = domain.FindDomainById(authContext.DomainId, http.StatusBadRequest)
				defaultDomain = &ud
			}
		} else {
			ud = domain.FindDomainById(authContext.DomainId, http.StatusBadRequest)
		}
		p = project.FindProjectById(authContext.ProjectId, http.StatusBadRequest)
		auth.Identity.Methods = authContext.Method
	case "password":
		if auth.Identity.Password.User.Domain.Id == "" && auth.Identity.Password.User.Domain.Name == "" {
			panic(execption.NewGoStoneError(http.StatusBadRequest, "user domain id/name must be present"))
		}
		if auth.Identity.Password.User.Domain.Name != "" {
			if auth.Identity.Password.User.Domain.Name == "default" {
				if defaultDomain != nil {
					ud = *defaultDomain
				} else {
					ud = domain.FindDomainByName(auth.Identity.Password.User.Domain.Name, http.StatusBadRequest)
					defaultDomain = &ud
				}
			} else {
				ud = domain.FindDomainByName(auth.Identity.Password.User.Domain.Name, http.StatusBadRequest)
			}
			auth.Identity.Password.User.Domain.Id = ud.Id
		} else {
			if auth.Identity.Password.User.Domain.Id == "default" {
				if defaultDomain != nil {
					ud = *defaultDomain
				} else {
					ud = domain.FindDomainById(auth.Identity.Password.User.Domain.Id, http.StatusBadRequest)
					defaultDomain = &ud
				}
			} else {
				ud = domain.FindDomainById(auth.Identity.Password.User.Domain.Id, http.StatusBadRequest)
			}
		}
		if auth.Identity.Password.User.Name == "" {
			panic(execption.NewGoStoneError(http.StatusBadRequest, "user name must be present"))
		}
		var result = user.FindUserByNameAndDomainId(auth.Identity.Password.User.Name, ud.Id, http.StatusBadRequest)
		u = result
		if u.Enabled == 0 {
			panic(execption.NewGoStoneError(http.StatusUnauthorized, "user is disabled"))
		}
		if !utils.CheckSM3Pwd(auth.Identity.Password.User.Password, u.PasswordSm3) {
			if !utils.CheckPwd(auth.Identity.Password.User.Password, u.Password) {
				log.Error("check password failed username:" + auth.Identity.Password.User.Name + " password:" + auth.Identity.Password.User.Password)
				panic(execption.NewGoStoneError(http.StatusUnauthorized, "Invalid username or password"))
			} else {
				var password = user.FindPasswordByLocalUserId(u.LocalUserId, http.StatusUnauthorized)
				password.PasswordSm3 = utils.GenSM3Pwd(auth.Identity.Password.User.Password)
				user.UpdatePassword(&password)
			}
		}
	}
	roleList := assignment.FindRoleByProjectIdAndUserId(p.Id, u.Id)
	if len(roleList) == 0 {
		panic(execption.NewGoStoneError(http.StatusUnauthorized, "user is not belong this project"))
	}
	roleStr := make([]string, 0)
	for _, r := range roleList {
		roleStr = append(roleStr, r.Name)
	}
	authContext := utils.AuthContext{
		UserId:    u.Id,
		ProjectId: p.Id,
		Role:      roleStr,
		DomainId:  u.DomainId,
		Method:    auth.Identity.Methods,
	}
	var (
		token      string
		issueAtZ   utils.JSONRFC3339Milli
		expiresAtz utils.JSONRFC3339Milli
	)
	if auth.Expiration != 0 {
		token, issueAtZ, expiresAtz = tokenValid.SignByExpiration(authContext, auth.Expiration)
	} else {
		token, issueAtZ, expiresAtz = tokenValid.Sign(authContext)
	}

	ctx.Response().Header().Set("X-Subject-Token", token)
	return ctx.JSON(http.StatusCreated, map[string]Token{
		"token": GetToken(u, p, pd, ud, auth.Identity.Methods, issueAtZ, expiresAtz, roleList),
	})
}

func ValidateToken(ctx echo.Context) error {
	token := ctx.Request().Header.Get("X-Subject-Token")
	if token == "" {
		panic(execption.NewGoStoneError(http.StatusUnauthorized, "token  is null"))
	}
	var (
		p  model.Project
		u  model.UserInfo
		pd model.Project
		ud model.Project
	)
	allow := ctx.QueryParam("allow_expired")
	var tokenValid utils.Token
	if utils.IsJwtToken(token) {
		tokenValid = utils.GetTokenMethod("jwt")
	} else {
		tokenValid = utils.GetTokenMethod("fernet")
	}
	var authContext *utils.AuthContext
	if allow == "true" {
		authContext = tokenValid.ValidateCanExpired(token)
	} else {
		authContext = tokenValid.Validate(token)
	}
	var result = user.FindUserById(authContext.UserId, http.StatusBadRequest)
	u = result
	p = project.FindProjectById(authContext.ProjectId, http.StatusBadRequest)
	pd = domain.FindDomainById(p.DomainId, http.StatusBadRequest)
	if strings.ToLower(pd.DomainId) == strings.ToLower(u.DomainId) {
		ud = pd
	} else {
		ud = domain.FindDomainById(u.DomainId, http.StatusBadRequest)
	}
	roles := assignment.FindRoleByProjectIdAndUserId(p.Id, u.Id)
	ctx.Response().Header().Set("X-Subject-Token", token)
	authContext.IssuedAtZ = utils.JSONRFC3339Milli(time.Time(authContext.IssuedAtZ).UTC())
	authContext.ExpiresAtZ = utils.JSONRFC3339Milli(time.Time(authContext.ExpiresAtZ).UTC())
	return ctx.JSON(http.StatusOK, map[string]Token{
		"token": GetToken(u, p, pd, ud, []string{"password"}, authContext.IssuedAtZ, authContext.ExpiresAtZ, roles),
	})
}

func GetUserTokenByAdmin(ctx echo.Context) error {
	auth := utils.GetTokenMethod(ctx.Get(utils.TokenTypeKey)).GetAuthContext(ctx)
	if !policy.Check("admin_get_token", auth, nil) {
		panic(execption.NewGoStoneError(http.StatusForbidden, "policy not valid"))
	}

	var tq request.TokenRequest
	err := checkTokenRequest(ctx, &tq)
	if err != nil {
		return err
	}
	var (
		p  model.Project
		u  model.UserInfo
		pd model.Project
		ud model.Project
	)
	assign := assignment.FindAssignmentByUserId(tq.UserId)
	if len(assign) == 0 {
		panic(execption.NewGoStoneError(http.StatusNotFound, "user don't have project_id"))
	}
	if len(assign) > 1 {
		log.Warnf("user has multi project assign:[%+v]", assign)
	}
	projectId := assign[0].TargetId
	p = project.FindProjectById(projectId, http.StatusBadRequest)
	var result = user.FindUserById(tq.UserId, http.StatusBadRequest)
	u = result
	pd = domain.FindDomainById(p.DomainId, http.StatusBadRequest)
	ud = domain.FindDomainById(u.DomainId, http.StatusBadRequest)
	roleList := assignment.FindRoleByProjectIdAndUserId(p.Id, u.Id)
	roleStr := make([]string, 0)
	for _, r := range roleList {
		roleStr = append(roleStr, r.Name)
	}
	authContext := utils.AuthContext{
		UserId:    u.Id,
		ProjectId: p.Id,
		Role:      roleStr,
		DomainId:  u.DomainId,
		Method:    []string{"token"},
	}
	var (
		token      string
		issueAtZ   utils.JSONRFC3339Milli
		expiresAtz utils.JSONRFC3339Milli
	)
	tokenValid := utils.GetTokenMethod(utils.TokenType)
	token, issueAtZ, expiresAtz = tokenValid.Sign(authContext)
	ctx.Response().Header().Set("X-Subject-Token", token)
	return ctx.JSON(http.StatusCreated, GetToken(u, p, pd, ud, []string{"password"}, issueAtZ, expiresAtz, roleList))
}
func GetUserToken(id string) string {
	var (
		p model.Project
		u model.UserInfo
	)
	assign := assignment.FindAssignmentByUserId(id)
	if len(assign) == 0 {
		panic(execption.NewGoStoneError(http.StatusNotFound, "user don't have project_id"))
	}
	projectId := assign[0].TargetId
	p = project.FindProjectById(projectId, http.StatusBadRequest)
	var result = user.FindUserById(id, http.StatusBadRequest)
	u = result
	roleList := assignment.FindRoleByProjectIdAndUserId(p.Id, u.Id)
	roleStr := make([]string, 0)
	for _, r := range roleList {
		roleStr = append(roleStr, r.Name)
	}
	authContext := utils.AuthContext{
		UserId:    u.Id,
		ProjectId: p.Id,
		Role:      roleStr,
		DomainId:  u.DomainId,
		Method:    []string{"token"},
	}
	var (
		token string
	)
	tokenValid := utils.GetTokenMethod(utils.TokenType)
	token, _, _ = tokenValid.Sign(authContext)
	return token
}

//检查创建宿主机请求体
func checkTokenRequest(ctx echo.Context, request *request.TokenRequest) error {
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

func GetToken(u model.UserInfo,
	p model.Project, pd model.Project, ud model.Project, methods []string, issueAtZ utils.JSONRFC3339Milli,
	expiresAtz utils.JSONRFC3339Milli, roleList []assignment.Role) Token {

	response := Token{
		IssuedAt:  issueAtZ,
		ExpiresAt: expiresAtz,
		IsDomain:  false,
		Roles:     roleList,
		Project: Project{
			Id:   p.Id,
			Name: p.Name,
			Domain: Domain{
				Id:   pd.Id,
				Name: pd.Name,
			},
		},
		User: User{
			Domain: Domain{
				Id:   ud.Id,
				Name: ud.Name,
			},
			Id:   u.Id,
			Name: u.Name,
		},
		Catalog: getCatalog(p.Id, u.Id),
		Methods: methods,
	}
	return response
}

func getCatalog(projectId string, userId string) []Catalog {
	catalogs := GetCatalogByCache()
	var catalog []Catalog
	for _, c := range catalogs {
		es := make([]EndPoint, 0)
		for _, end := range c.Endpoints {
			url := formatUrl(end.Url, projectId, userId)
			end.Url = url
			es = append(es, end)
		}
		c.Endpoints = es
		catalog = append(catalog, c)
	}
	return catalog
}

var cache *cache2go.CacheTable
var cacheLock sync.Mutex

func init() {
	cache = cache2go.Cache("Catalog")
}

//缓存catalog 5s过期
// fixme 利用本地cache可能会造成数据延迟,但目前endpoint几乎修改所以5s延迟应该是被允许的
func GetCatalogByCache() []Catalog {
	cacheLock.Lock()
	defer cacheLock.Unlock()
	res, err := cache.Value("catalog")
	if err != nil {
		log.Debug("Item is not cached (anymore).")
		catalogs := GetBaseCateLog()
		cache.Add("catalog", 5*time.Second, catalogs)
		return catalogs
	}
	return res.Data().([]Catalog)
}

func GetBaseCateLog() []Catalog {
	services := service.FindEnabledService()
	catalogs := make([]Catalog, 0)
	for _, sev := range services {
		var extra map[string]interface{}
		utils.Byte2Struct([]byte(sev.Extra), &extra)
		endpoints := endpoint.GetEndpointsByServiceId(sev.Id)
		es := make([]EndPoint, 0)
		for _, end := range endpoints {
			es = append(es, EndPoint{
				Id:        end.Id,
				Interface: end.Interface,
				Region:    end.RegionId,
				RegionId:  end.RegionId,
				Url:       end.URL,
			})
		}
		catalog := Catalog{
			Type:      sev.Type,
			Id:        sev.Id,
			Name:      extra["name"].(string),
			Endpoints: es,
		}
		catalogs = append(catalogs, catalog)
	}
	return catalogs
}

var regx = regexp.MustCompile("\\((.+)\\)")

func formatUrl(url string, projectId string, userId string) string {
	if regx.Match([]byte(url)) {
		values := make([]interface{}, 0)
		for _, str := range regx.FindStringSubmatch(url) {
			if str == "project_id" || str == "tenant_id" {
				values = append(values, projectId)
			} else if str == "user_id" {
				values = append(values, userId)
			}
		}
		url = strings.ReplaceAll(url, "(tenant_id)", "")
		url = strings.ReplaceAll(url, "(project_id)", "")
		url = strings.ReplaceAll(url, "(user_id)", "")
		url = fmt.Sprintf(url, values...)
	}
	return url
}

type Auth struct {
	Identity struct {
		Methods  []string
		Password struct {
			User struct {
				Name   string
				Domain struct {
					Name string
					Id   string
				}
				Password string
			}
		}
		Token struct {
			Id string
		}
	}
	Scope struct {
		Project struct {
			Domain struct {
				Name string
				Id   string
			}
			Name string
			Id   string
		}
	}
	Expiration int64
}

type Token struct {
	IssuedAt  utils.JSONRFC3339Milli `json:"issued_at"`
	ExpiresAt utils.JSONRFC3339Milli `json:"expires_at"`
	IsDomain  bool                   `json:"is_domain"`
	Roles     []assignment.Role      `json:"roles"`
	Project   Project                `json:"project"`
	User      User                   `json:"user"`
	Catalog   []Catalog              `json:"catalog"`
	Methods   []string               `json:"methods"`
}

type Catalog struct {
	Type      string     `json:"type"`
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	Endpoints []EndPoint `json:"endpoints"`
}

type EndPoint struct {
	Id        string `json:"id"`
	Interface string `json:"interface"`
	Region    string `json:"region"`
	RegionId  string `json:"region_id"`
	Url       string `json:"url"`
}

type User struct {
	Domain Domain `json:"domain"`
	Id     string `json:"id"`
	Name   string `json:"name"`
}

type Project struct {
	Domain Domain `json:"domain"`
	Id     string `json:"id"`
	Name   string `json:"name"`
}

type Domain struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
