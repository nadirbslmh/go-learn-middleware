package routes

import (
	"go-learn-middleware/controllers"
	"go-learn-middleware/middlewares"
	"go-learn-middleware/utils"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	loggerConfig := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	loggerMiddleware := loggerConfig.Init()

	e.Use(loggerMiddleware)

	jwtConfig := middlewares.JWTConfig{
		SecretKey:       utils.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	authMiddlewareConfig := jwtConfig.Init()

	userController := controllers.InitUserController(&jwtConfig)

	auth := e.Group("/api/v1/auth", echojwt.WithConfig(authMiddlewareConfig))

	auth.POST("/register", userController.Register)
	auth.POST("/login", userController.Login)

	users := e.Group("/api/v1/users")

	users.GET("/me", userController.GetUser)
}
