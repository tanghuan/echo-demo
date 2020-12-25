package router

import (
	"github.com/labstack/echo/v4"

	"skinshub-api/apis"
)

// New 创建 router
func New() *echo.Echo {

	// init router
	router := echo.New()

	router.GET("/users", apis.Find)
	router.POST("/login", apis.Login)

	return router
}
