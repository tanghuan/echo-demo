package router

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"skinshub-api/apis"
	"skinshub-api/middlewares"
)

// New 创建 router
func New() *echo.Echo {

	// init router
	router := echo.New()

	// add middlewares
	router.Use(middlewares.ServerHeader)

	router.GET("/users", apis.Find)
	router.POST("/login", apis.Login)

	router.GET("/secure", func(c echo.Context) error {
		return c.String(http.StatusOK, "Secure")
	}, middlewares.Auth("Admin"))

	return router
}
