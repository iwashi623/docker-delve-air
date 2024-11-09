package main

import (
	"docker-delve-air/handler"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", rootHandler)
	e.GET("/hello", handler.HelloHandler)
	e.GET("/hello/:user_name", handler.HelloUserHandler)
	e.Logger.Fatal(e.Start(":3333"))
}

func rootHandler(c echo.Context) error {
	return c.String(200, "Root Handler")
}
