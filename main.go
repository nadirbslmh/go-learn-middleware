package main

import (
	"go-learn-middleware/database"
	"go-learn-middleware/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	database.InitDB()

	database.Migrate()

	e := echo.New()

	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":1323"))
}
