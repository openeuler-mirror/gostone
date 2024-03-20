package utils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"time"
	"work.ctyun.cn/git/GoStack/gostone/execption"
)

type JwtToken struct {
}

func NewJwtToken() Token {
	return &JwtToken{}
}

func (j *JwtToken) Name() string {
	return "jwt"
}

func (j *JwtToken) Sign(claims AuthContext) (string, JSONRFC3339Milli, JSONRFC3339Milli) {
	t := time.Duration(expireTime) * time.Minute
	now := time.Now().UTC()
	claims.IssuedAtZ = JSONRFC3339Milli(now)
	claims.IssuedAt = now.Unix()
	claims.ExpiresAtZ = JSONRFC3339Milli(now.Add(t))
	claims.ExpiresAt = now.Add(t).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod(SignMethod), claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Error(err)
		panic(execption.NewGoStoneError(execption.StatusInternalServerError, "token sign error"))
	}
	return tokenString, claims.IssuedAtZ, claims.ExpiresAtZ
}

func (j *JwtToken) SignByExpiration(claims AuthContext, expiration int64) (string, JSONRFC3339Milli, JSONRFC3339Milli) {
	t := time.Duration(expiration) * time.Minute
	now := time.Now().UTC()
	claims.IssuedAtZ = JSONRFC3339Milli(now)
	claims.IssuedAt = now.Unix()
	claims.ExpiresAtZ = JSONRFC3339Milli(now.Add(t))
	claims.ExpiresAt = now.Add(t).Unix()
	token := jwt.NewWithClaims(jwt.GetSigningMethod(SignMethod), claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Error(err)
		panic(execption.NewGoStoneError(execption.StatusInternalServerError, "token sign error"))
	}
	return tokenString, claims.IssuedAtZ, claims.ExpiresAtZ
}

func (j *JwtToken) Validate(token string) *AuthContext {
	authContext := new(AuthContext)
	t, err := jwt.ParseWithClaims(token, authContext, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !t.Valid {
		panic(execption.NewGoStoneError(execption.StatusUnauthorized, "token valid error"))
	}
	return t.Claims.(*AuthContext)
}

func (j *JwtToken) ValidateCanExpired(token string) *AuthContext {
	authContext := new(AuthContext)
	t, err := jwt.ParseWithClaims(token, authContext, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil || !t.Valid {
		if ge, ok := err.(jwt.ValidationError); ok {
			if ge.Errors == jwt.ValidationErrorExpired {
				return t.Claims.(*AuthContext)
			}
		}
		panic(execption.NewGoStoneError(execption.StatusUnauthorized, "token valid error"))
	}
	return t.Claims.(*AuthContext)
}

func (j *JwtToken) GetAuthContext(ctx echo.Context) AuthContext {
	c := ctx.Get("user").(*jwt.Token)
	authContext := c.Claims.(*AuthContext)
	return *authContext
}

func init() {
	jwt.RegisterSigningMethod(SM3, func() jwt.SigningMethod {
		return &SigningMethodSM3{
			Name: SM3,
		}
	})
}

const SM3 = "SM3"

type SigningMethodSM3 struct {
	Name string
}

func (m *SigningMethodSM3) Alg() string {
	return m.Name
}

// Implements the Verify method from SigningMethod
func (m *SigningMethodSM3) Verify(signingString, signature string, key interface{}) error {
	if !CheckSM3Pwd(signingString, signature) {
		return jwt.ErrSignatureInvalid
	}
	return nil
}

// Implements the Sign method from SigningMethod
func (m *SigningMethodSM3) Sign(signingString string, key interface{}) (string, error) {
	return GenSM3Pwd(signingString), nil
}
