package controllers

import (
	"go-learn-middleware/services"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func InitUserController() UserController {
	return UserController{
		service: services.InitUserService(),
	}
}

func (controller *UserController) Register(c echo.Context) error {
	return c.String(200, "Not Implemented!")
}

func (controller *UserController) Login(c echo.Context) error {
	return c.String(200, "Not Implemented!")
}

func (controller *UserController) GetUser(c echo.Context) error {
	return c.String(200, "Not Implemented!")
}
