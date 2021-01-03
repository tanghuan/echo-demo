package dtos

import (
	"github.com/dgrijalva/jwt-go"
)

// JwtClaimsDto jwt claims
type JwtClaimsDto struct {
	*jwt.StandardClaims

	// Username 用户名
	Username string `json:"username"`

	// Role 角色
	Role string `json:"role"`
}
