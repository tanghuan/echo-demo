package config

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"

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

}
