package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

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

		// TODO
		return c.String(http.StatusOK, "Secure")
	}, middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte("secret"),
	}), middlewares.Auth("Admin"))

	return router
}
