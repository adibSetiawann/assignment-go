package main

import (
	"github.com/adibSetiawann/transaction-api-go/config"
	"github.com/adibSetiawann/transaction-api-go/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {

	config.ConnectDatabase()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"message": "success create API",
		})
	})

	routes.MerchantRoute(app)
	routes.ProductRoute(app)
	routes.CustomerRoute(app)
	routes.TransactionRoute(app)

	app.Listen(":8080")
}
