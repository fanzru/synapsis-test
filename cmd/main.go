package main

import (
	"fmt"
	"log"
	"synapsis-test/cmd/services"
	"synapsis-test/infrastructure/config"
	"synapsis-test/infrastructure/database"
	"synapsis-test/infrastructure/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Println("Start Services....")

	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to build config: %v", err)
	}

	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	// middleware
	middlewareAuth := services.RegisterMiddleware(db, cfg)

	// services
	accountsHandler := services.RegisterServiceAccounts(db, cfg)
	productHandler := services.RegisterServiceProduct(db, cfg)

	// register routes
	mHandler := routes.ModuleHandler{
		AccountHandler: accountsHandler,
		MiddlewareAuth: middlewareAuth,
		ProductHandler: productHandler,
	}

	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e = routes.NewRoutes(mHandler, e)

	log.Fatal(e.Start(fmt.Sprintf(":%v", cfg.PortApps)))
}
