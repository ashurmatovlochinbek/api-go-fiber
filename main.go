package main

import (
	"api-go-fiber/database"
	noteroutes "api-go-fiber/internals/routes/noteRoutes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	database.ConnectDB()

	noteroutes.SetUpNoteRotes(app)

	log.Fatal(app.Listen(":8080"))
}