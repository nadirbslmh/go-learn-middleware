package controllers

import (
	"go-learn-middleware/middlewares"
	"go-learn-middleware/models"
	"go-learn-middleware/services"
	"net/http"
	"strconv"

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
	var userInput models.RegisterInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	user, err := controller.service.Register(userInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "failed",
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, models.Response{
		Status:  "success",
		Message: "user registered",
		Data:    user,
	})
}

func (controller *UserController) Login(c echo.Context) error {
	var userInput models.LoginInput

	if err := c.Bind(&userInput); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "failed",
			Message: "invalid request",
		})
	}

	token, err := controller.service.Login(userInput)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Status:  "failed",
			Message: "invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "login success",
		Data:    token,
	})
}

func (controller *UserController) GetUser(c echo.Context) error {
	claims, err := middlewares.GetUser(c)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, models.Response{
			Status:  "failed",
			Message: "invalid token",
		})
	}

	user, err := controller.service.GetUser(strconv.Itoa(claims.ID))

	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
			Status:  "failed",
			Message: "user not found",
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Status:  "success",
		Message: "user data",
		Data:    user,
	})
}
