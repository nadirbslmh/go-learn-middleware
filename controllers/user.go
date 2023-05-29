package controllers

import (
	"go-learn-middleware/middlewares"
	"go-learn-middleware/services"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	service services.UserService
}

func InitUserController(jwtAuth *middlewares.JWTConfig) UserController {
	return UserController{
		service: services.InitUserService(jwtAuth),
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
