package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	EMiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
	"strings"
	"work.ctyun.cn/git/GoStack/gostone/utils"
)

type (
	// FERNETConfig defines the config for FERNET middleware.
	FERNETConfig struct {
		// Skipper defines a function to skip middleware.
		Skipper EMiddleware.Skipper

		// BeforeFunc defines a function which is executed just before the middleware.
		BeforeFunc EMiddleware.BeforeFunc

		// SuccessHandler defines a function which is executed for a valid token.
		SuccessHandler FERNETSuccessHandler

		// ErrorHandler defines a function which is executed for an invalid token.
		// It may be used to define a custom FERNET error.
		ErrorHandler FERNETErrorHandler

		// ErrorHandlerWithContext is almost identical to ErrorHandler, but it's passed the current context.
		ErrorHandlerWithContext FERNETErrorHandlerWithContext

		// Context key to store user information from the token into context.
		// Optional. Default value "user".
		ContextKey string

		// Claims are extendable claims data defining token content.
		// Optional. Default value jwt.MapClaims
		Claims jwt.Claims

		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "param:<name>"
		// - "cookie:<name>"
		TokenLookup string

		// AuthScheme to be used in the Authorization header.
		// Optional. Default value "Bearer".
		AuthScheme string

		Validate utils.Token
	}

	// FERNETSuccessHandler defines a function which is executed for a valid token.
	FERNETSuccessHandler func(echo.Context)

	// FERNETErrorHandler defines a function which is executed for an invalid token.
	FERNETErrorHandler func(error) error

	// FERNETErrorHandlerWithContext is almost identical to FERNETErrorHandler, but it's passed the current context.
	FERNETErrorHandlerWithContext func(error, echo.Context) error
)

// Errors
var (
	ErrFERNETMissing = echo.NewHTTPError(http.StatusBadRequest, "missing or malformed jwt")
)

var (
	// DefaultFERNETConfig is the default FERNET auth middleware config.
	DefaultFERNETConfig = FERNETConfig{
		Skipper:     EMiddleware.DefaultSkipper,
		ContextKey:  "user",
		TokenLookup: "header:" + echo.HeaderAuthorization,
		AuthScheme:  "Bearer",
		Claims:      jwt.MapClaims{},
		Validate:    utils.GetTokenMethod("fernet"),
	}
)

// FERNET returns a JSON Web Token (FERNET) auth middleware.
//
// For valid token, it sets the user in context and calls next handler.
// For invalid token, it returns "401 - Unauthorized" error.
// For missing token, it returns "400 - Bad Request" error.
//
// See: https://jwt.io/introduction
// See `FERNETConfig.TokenLookup`
func FERNET() echo.MiddlewareFunc {
	c := DefaultFERNETConfig
	return FERNETWithConfig(c)
}

// FERNETWithConfig returns a FERNET auth middleware with config.
// See: `FERNET()`.
func FERNETWithConfig(config FERNETConfig) echo.MiddlewareFunc {
	// Defaults
	if config.Skipper == nil {
		config.Skipper = DefaultFERNETConfig.Skipper
	}
	if config.ContextKey == "" {
		config.ContextKey = DefaultFERNETConfig.ContextKey
	}
	if config.Claims == nil {
		config.Claims = DefaultFERNETConfig.Claims
	}
	if config.TokenLookup == "" {
		config.TokenLookup = DefaultFERNETConfig.TokenLookup
	}
	if config.AuthScheme == "" {
		config.AuthScheme = DefaultFERNETConfig.AuthScheme
	}

	// Initialize
	parts := strings.Split(config.TokenLookup, ":")
	extractor := fernetFromHeader(parts[1], config.AuthScheme)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			if config.BeforeFunc != nil {
				config.BeforeFunc(c)
			}

			auth, err := extractor(c)
			if err != nil {
				if config.ErrorHandler != nil {
					return config.ErrorHandler(err)
				}

				if config.ErrorHandlerWithContext != nil {
					return config.ErrorHandlerWithContext(err, c)
				}
				return err
			}

			token := config.Validate.Validate(auth)
			// Store user information from token into context.
			c.Set(config.ContextKey, token)
			if config.SuccessHandler != nil {
				config.SuccessHandler(c)
			}
			c.Set("HasAuth", true)
			c.Set(utils.TokenTypeKey, "fernet")
			return next(c)

		}
	}
}

// jwtFromHeader returns a `jwtExtractor` that extracts token from the request header.
func fernetFromHeader(header string, authScheme string) jwtExtractor {
	return func(c echo.Context) (string, error) {
		auth := c.Request().Header.Get(header)
		if len(auth) != 0 {
			return auth, nil
		}
		return "", ErrFERNETMissing
	}
}
