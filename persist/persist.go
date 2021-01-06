package persist

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"skinshub-api/config"
	"skinshub-api/persist/entities"
)

// Con 数据库链接
var Con *gorm.DB

// 初始化数据库链接
func init() {
	// dsn := "root:admin@tcp(127.0.0.1:3306)/echo-demo?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	Con, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                      config.GetDSN(),
		DisableDatetimePrecision: true,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	Con.AutoMigrate(&entities.User{})

	// Con = db
}
