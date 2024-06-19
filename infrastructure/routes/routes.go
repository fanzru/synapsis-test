package routes

import (
	"log"
	"net/http"
	"synapsis-test/app/account/domain/constant"
	account "synapsis-test/app/account/http"
	product "synapsis-test/app/product/http"
	"synapsis-test/infrastructure/middleware"

	"github.com/labstack/echo/v4"
)

type ModuleHandler struct {
	MiddlewareAuth middleware.MiddlewareAuth
	AccountHandler account.AccountHandler
	ProductHandler product.ProductHandler
}

func NewRoutes(h ModuleHandler, app *echo.Echo) *echo.Echo {

	log.Println("Starting to create routing...")

	// test api connect or not
	app.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    "good connection",
			"message": "fanzru api 200 OK aman azaaaa!!!!",
		})
	})

	// accounts endpoint
	accountsgateway := app.Group("/accounts")
	accountsgateway.POST("/login", h.AccountHandler.LoginUser)
	accountsgateway.POST("/register", h.AccountHandler.CreateNewAccount)
	accountsgateway.PATCH("/password", h.MiddlewareAuth.BearerTokenMiddleware(h.AccountHandler.ChangePassword, []string{string(constant.User)}))

	// product endpoint
	productgateway := app.Group("/product")
	productgateway.GET("/", h.ProductHandler.FindProducts)

	// categories endpoint
	categories := app.Group("/categories")
	categories.GET("/", h.ProductHandler.FindCategories)

	return app
}
