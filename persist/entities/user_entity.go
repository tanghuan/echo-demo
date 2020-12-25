package entities

import (
	"gorm.io/gorm"
)

// User 用户信息
type User struct {
	gorm.Model

	// 用户状态
	Status uint

	// 用户名
	Username string

	// 登录密码
	Password string

	// 昵称
	Nickname string

	// 性别
	Gender uint

	// 角色
	Role uint

	// 头像地址
	Avatar string

	// 头像 hash 值
	AvatarHash string
}
