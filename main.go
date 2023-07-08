package main

import (
	"go-fiber-gorm/config"
	"go-fiber-gorm/migration"
	"go-fiber-gorm/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Inital Database
	config.DBInit()

	// Run Migration
	migration.RunMigration()

	// Initial Route
	route.RouteInit(app)

	app.Listen(":8080")
}
