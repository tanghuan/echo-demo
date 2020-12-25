package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

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
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World!")
	})

	e.Logger.Fatal(e.Start(":8000"))
}
