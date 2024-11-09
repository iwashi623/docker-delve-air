package handler

import (
	"docker-delve-air/infra/repository"
	"docker-delve-air/usecase"

	"github.com/labstack/echo/v4"
)

func HelloHandler(c echo.Context) error {
	return c.String(200, "Hello, World!")
}

func HelloUserHandler(c echo.Context) error {
	userName := c.Param("user_name")

	usecase := usecase.NewHelloUserUsecase(
		repository.NewUserRepository(),
	)
	msg, err := usecase.Exec(userName)
	if err != nil {
		return c.String(500, "Internal Server Error")
	}

	return c.String(200, msg)
}
