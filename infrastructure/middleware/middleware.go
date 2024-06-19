package middleware

import (
	"strings"
	"synapsis-test/infrastructure/config"
	"synapsis-test/infrastructure/database"
	"synapsis-test/pkg/jwt"
	"synapsis-test/pkg/response"

	"github.com/labstack/echo/v4"
)

// Service Authorizer
type MiddlewareAuth struct {
	DB  database.Connection
	Cfg config.Config
}

func NewServiceAuthorizer(db database.Connection, cfg config.Config) MiddlewareAuth {
	return MiddlewareAuth{
		DB:  db,
		Cfg: cfg,
	}
}

func (m MiddlewareAuth) BearerTokenMiddleware(next echo.HandlerFunc, rolesAllowed []string) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		userId := m.getUserIdAndTypeFromJWT(ctx)
		if userId == 0 {
			return response.ResponseErrorUnauthorized(ctx)
		}

		return next(ctx)
	}
}

func (m MiddlewareAuth) getUserIdAndTypeFromJWT(ctx echo.Context) int64 {
	authHeader := ctx.Request().Header.Get("Authorization")
	if authHeader == "" {
		return 0
	}

	headerSplit := strings.Split(authHeader, " ")
	if len(headerSplit) < 1 {
		return 0
	}

	// Get Token
	token := headerSplit[1]

	// Get JWT Claims
	claims, err := jwt.DecodeToken(token, m.Cfg.JWTTokenSecret)
	if err != nil {
		return 0
	}

	// Bind UserID to context
	ctx.Set("user_id", claims.UserId)
	ctx.Set("username", claims.Username)
	return claims.UserId
}

func (m MiddlewareAuth) InitMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		externalPassword := ctx.Request().Header.Get("X-External-Password")
		if externalPassword != m.Cfg.ExternalPassword {
			return response.ResponseErrorUnauthorized(ctx)
		}

		return next(ctx)
	}
}
