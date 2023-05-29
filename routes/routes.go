package routes

import (
	"go-learn-middleware/controllers"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	//TODO: add logging and jwt middlewares

	userController := controllers.InitUserController()

	auth := e.Group("/api/v1/auth")

	auth.POST("/register", userController.Register)
	auth.POST("/login", userController.Login)

	users := e.Group("/api/v1/users")

	users.GET("/me", userController.GetUser)
}
