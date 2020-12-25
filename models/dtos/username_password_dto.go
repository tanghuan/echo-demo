package dtos

// UsernamePasswordDto 登录用户名和密码
type UsernamePasswordDto struct {
	Username string `json:"username"`

	Password string `json:"password"`
}
