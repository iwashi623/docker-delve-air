package handler

import "github.com/labstack/echo"

func helloHandler(c echo.Context) error {
	return c.String(200, "Hello, Echo!")
}
