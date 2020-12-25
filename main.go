package main

import (
	"fmt"

	_ "skinshub-api/persist"
	"skinshub-api/router"
	"skinshub-api/utils"
)

func init() {
	fmt.Println("package main init ......")
}

func main() {

	name := "TangHuan"

	res, err := utils.SayHello(name)
	if err != nil {
		panic(err)
	}

	fmt.Println(res)

	// instance echo
	e := router.New()

	e.Logger.Fatal(e.Start(":8000"))
}
