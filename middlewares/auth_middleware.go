package middlewares

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/labstack/echo/v4"

	"skinshub-api/config"
	"skinshub-api/models/dtos"
	"skinshub-api/utils"
)

// Auth 鉴权中间件
func Auth(roles ...string) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println("==================Auth Middleware==================")

			claims := dtos.JwtClaimsDto{}
			_, err := request.ParseFromRequestWithClaims(c.Request(), request.AuthorizationHeaderExtractor, &claims, func(token *jwt.Token) (interface{}, error) {
				return config.VerifyKey, nil
			})
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			fmt.Println("claims ======> ", claims)

			// check role
			hasPermission := utils.ContainsString(roles, claims.Role)
			if !hasPermission {
				return echo.NewHTTPError(http.StatusUnauthorized)
			}

			return next(c)
		}
	}
}
