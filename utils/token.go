package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"regexp"
	"work.ctyun.cn/git/GoStack/gostone/connect"
)

var (
	secret       string
	expireTime   = 30
	fernetPath   string
	TokenType    string
	TokenTypeKey = "Token-Type"
	SignMethod   = "HS256"
)

func InitToken() {
	SetSecret(connect.AppConf.GoStone.Secret, connect.AppConf.GoStone.ExpiresTime, connect.AppConf.GoStone.FernetPath, connect.AppConf.GoStone.TokenType)
	SignMethod = connect.AppConf.GoStone.SignMethod
}
func SetSecret(s string, expire int, f string, t string) {
	secret = s
	expireTime = expire
	fernetPath = f
	TokenType = t
	//如果是fernet 则需要监控秘钥的变化
	loadKey()
	watchKey()
}

type Token interface {
	Sign(claims AuthContext) (string, JSONRFC3339Milli, JSONRFC3339Milli)

	Name() string

	Validate(token string) *AuthContext

	SignByExpiration(claims AuthContext, expiration int64) (string, JSONRFC3339Milli, JSONRFC3339Milli)

	ValidateCanExpired(token string) *AuthContext

	GetAuthContext(ctx echo.Context) AuthContext
}

type AuthContext struct {
	jwt.StandardClaims
	UserId     string           `json:"user_id"`
	ProjectId  string           `json:"project_id"`
	DomainId   string           `json:"domain_id"`
	Role       []string         `json:"roles"`
	ConsumerId string           `json:"consumer_id"`
	IssuedAtZ  JSONRFC3339Milli `json:"issued_at"`
	ExpiresAtZ JSONRFC3339Milli `json:"expire_at"`
	Method     []string
}

func GetToken() Token {
	if TokenType == "jwt" {
		return NewJwtToken()
	}
	return NewFernetToken()
}

const tokenReg = "^[A-Za-z0-9-_=]+\\.[A-Za-z0-9-_=]+\\.?[A-Za-z0-9-_.+/=]*$"

var tokenRegx = regexp.MustCompile(tokenReg)

//判断是JWT/Fernet
func IsJwtToken(token string) bool {
	return tokenRegx.Match([]byte(token))
}

//根据类型获取tokenMethod
func GetTokenMethod(t interface{}) Token {
	if t == "jwt" {
		return NewJwtToken()
	}
	return NewFernetToken()
}
