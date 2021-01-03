package apis

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"skinshub-api/config"
	"skinshub-api/models/dtos"
	"skinshub-api/models/vos"
	"skinshub-api/persist"
	"skinshub-api/persist/entities"
)

func init() {
	fmt.Println("package apis user_apis init ......")
	persist.Con.Create(&entities.User{
		Status:   1,
		Username: "tanghuan",
		Password: "admin",
		Nickname: "tang",
		Gender:   1,
		Role:     1,
		Avatar:   "https://oss.yuefantutor.com",
	})
}

// Find 查询用户列表
func Find(c echo.Context) error {

	dto := &dtos.PageDto{}

	err := c.Bind(dto)
	if err != nil {
		log.Fatal(err)
	}

	var users []entities.User
	persist.Con.Offset(int(dto.Offset())).Limit(int(dto.Size)).Find(&users)

	vo := vos.PageVo{}
	vo.Page = dto.Page
	vo.Size = dto.Size
	vo.Data = users

	return c.JSON(http.StatusOK, users)
}

// Login 用户登录
func Login(c echo.Context) error {

	dto := dtos.UsernamePasswordDto{}

	err := c.Bind(&dto)
	if err != nil {
		log.Fatal(err)
	}

	// validate params
	if strings.TrimSpace(dto.Username) == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid_params")
	}

	if strings.TrimSpace(dto.Password) == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid_params")
	}
	// user := entities.User{}
	var user entities.User

	res := persist.Con.Where("username = ?", dto.Username).First(&user)

	if res.Error != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid_username_password")
	}

	// TODO add bcrypt
	if dto.Password != user.Password {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid_username_password")
	}

	// Create token
	token := jwt.New(jwt.SigningMethodRS256)

	// Set claims
	// claims := token.Claims.(jwt.MapClaims)
	// claims["username"] = user.Username
	// claims["role"] = "Admin"
	// claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token.Claims = &dtos.JwtClaimsDto{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
		user.Username,
		"Admin",
	}

	// Generate encoded token and send it as response.
	t, err := token.SignedString(config.SignKey)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})

}
