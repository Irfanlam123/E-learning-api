package main

import (
	"doctor-on-demand/config"
	"doctor-on-demand/initializers"
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// Initialize Echo instance
	e := echo.New()
	e.Use(middleware.CORS())
	// Middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // You can restrict this later
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	// Connect to DB
	db := config.ConnectDB()

	// Close DB connection gracefully
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get DB from GORM: %v", err)
	}
	defer sqlDB.Close()

	app := initializers.Initializers()
	app.SetupRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start("0.0.0.0:8080"))
}
