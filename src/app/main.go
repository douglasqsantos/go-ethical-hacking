package main

import (
	"app/config"
	"app/database"
	"app/routes"
	"app/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"time"
)


func main() {

	// Setting the App Configuration
	cfg := config.New()

	// Database connection.
	database.Connect()

	// Running the migrations
	database.AutoMigrate()

	// Getting the fiber
	app := fiber.New()

	// Enabling Cors
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	// Enabling logs
	// https://github.com/gofiber/fiber/tree/master/middleware/logger
	app.Use(logger.New(logger.Config{
		Format:       "[${time}] ${yellow} ${status} ${reset} - ${latency} ${magenta} ${method} ${green} ${path} ${queryParams}\n",
		TimeFormat:   "02-Jan-2006 15:04:05",
		TimeZone:     cfg.Timezone,
		TimeInterval: 500 * time.Millisecond,
	}))

	// Calling the Routes
	routes.Setup(app)

	// starting the server
	err := app.Listen(cfg.AppPort)
	utils.HandleError(err)
}
