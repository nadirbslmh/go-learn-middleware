package routes

import (
	"go-learn-middleware/controllers"
	"go-learn-middleware/middlewares"
	"go-learn-middleware/utils"
	"time"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoutes(e *echo.Echo) {
	loggerConfig := middlewares.LoggerConfig{
		Format: "[${time_rfc3339}] ${status} ${method} ${host} ${path} ${latency_human}" + "\n",
	}

	loggerMiddleware := loggerConfig.Init()

	e.Use(loggerMiddleware)

	e.Use(middleware.Recover())

	config := middleware.RateLimiterMemoryStoreConfig{
		Rate:      10,
		Burst:     30,
		ExpiresIn: 3 * time.Minute,
	}

	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStoreWithConfig(config)))

	jwtConfig := middlewares.JWTConfig{
		SecretKey:       utils.GetConfig("JWT_SECRET_KEY"),
		ExpiresDuration: 1,
	}

	authMiddlewareConfig := jwtConfig.Init()

	userController := controllers.InitUserController(&jwtConfig)

	auth := e.Group("/api/v1/auth")

	auth.POST("/register", userController.Register)
	auth.POST("/login", userController.Login)

	users := e.Group("/api/v1/users", echojwt.WithConfig(authMiddlewareConfig))

	users.GET("/me", userController.GetUser)
}
