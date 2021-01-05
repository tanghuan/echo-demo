package config

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/dgrijalva/jwt-go"
)

const (
	privKeyPath = "config/private.pem" // openssl genrsa -out app.rsa keysize
	pubKeyPath  = "config/public.pub"  // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

var (
	// VerifyKey rsa jwt 校验公钥
	VerifyKey *rsa.PublicKey

	// SignKey rsa jwt 签发私钥
	SignKey *rsa.PrivateKey

	// DatabaseHost 数据库地址
	DatabaseHost = os.Getenv("DatabaseHost")

	// DatabasePort 数据库端口
	DatabasePort = os.Getenv("DatabasePort")

	// DatabaseName 数据库名称
	DatabaseName = os.Getenv("DatabaseName")

	// DatabaseUsername 数据库链接用户名
	DatabaseUsername = os.Getenv("DatabaseUsername")

	// DatabasePassword 数据库链接密码
	DatabasePassword = os.Getenv("DatabasePassword")
)

func init() {
	fmt.Println("package config init ...... ")
	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}

	if DatabaseHost == "" {
		DatabaseHost = "127.0.0.1"
	}

	if DatabasePort == "" {
		DatabasePort = "3306"
	}

	if DatabaseName == "" {
		DatabaseName = "echo-demo"
	}

	if DatabaseUsername == "" {
		DatabaseUsername = "root"
	}

	if DatabasePassword == "" {
		DatabasePassword = "admin"
	}

}

// GetDSN 获取数据库链接
func GetDSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", DatabaseUsername, DatabasePassword, DatabaseHost, DatabasePort, DatabaseName)
}
