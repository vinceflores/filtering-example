package main

import (
	"filtering-sample/database"
	"filtering-sample/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()
	app := fiber.New()
	setUpRoutes(app)
	app.Listen(":3000")
}

func setUpRoutes(app *fiber.App){
	app.Post("/student/filtering", handlers.FilteringRequest)
	app.Post("/createListing", handlers.CreateListing)
	app.Get("/listing/details/:id", handlers.GetListingDetails)
}