package services

import (
	"synapsis-test/infrastructure/config"
	"synapsis-test/infrastructure/database"
	"synapsis-test/infrastructure/middleware"
)

func RegisterMiddleware(db database.Connection, cfg config.Config) middleware.MiddlewareAuth {
	middlewareAuth := middleware.NewServiceAuthorizer(db, cfg)
	return middlewareAuth
}
