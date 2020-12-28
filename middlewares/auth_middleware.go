package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"skinshub-api/utils"
)

// Auth 鉴权中间件
func Auth(roles ...string) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("==================Auth Middleware==================")
			user := c.Get("user")
			if user == nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			token, ok := user.(*jwt.Token)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}
			mapClaims := token.Claims.(jwt.MapClaims)

			role, ok := mapClaims["role"].(string)
			if !ok {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			// check role
			hasPermission := utils.ContainsString(roles, role)
			if !hasPermission {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			return next(c)
		}
	}
}
