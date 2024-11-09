package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", rootHandler)
	e.GET("/hello", helloHandler)
	e.Logger.Fatal(e.Start(":3333"))
}

func rootHandler(c echo.Context) error {
	fmt.Println("Hello, World!")
	return c.String(200, "Hello, World!")
}
