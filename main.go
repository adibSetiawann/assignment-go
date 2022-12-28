package main

import (
	"github.com/adibSetiawann/transaction-api-go/model"
	"github.com/adibSetiawann/transaction-api-go/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	model.ConnectDatabase()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"hello": "world",
		})
	})

	routes.RouteInit(app)

	app.Listen(":8080")
}
