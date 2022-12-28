package routes

import (
	"github.com/adibSetiawann/transaction-api-go/controllers"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App)  {
	// customer controller
	app.Get("/customers", controllers.GetAllCustomer)
	app.Post("/customers", controllers.RegisterCustomer)
	// app.Get("/customers", controllers.)
	app.Put("/customers/:id", controllers.UpdateCustomer)

	// product controller
	app.Get("/products", controllers.GetAllProduct)
	app.Post("/products", controllers.RegisterProduct)
	// app.Get("/products", controllers.)
	app.Put("/products/:id", controllers.UpdateProduct)

	// merchant controller
	app.Get("/merchants", controllers.GetAllMerchant)
	app.Post("/merchants", controllers.RegisterMerchant)
	// app.Get("/merchants", controllers.)
	app.Put("/merchants/:id", controllers.UpdateMerchant)

	// transaction controller
	app.Get("/transactions", controllers.GetAllTransaction)
	app.Post("/transactions", controllers.CreateTransaction)
	// app.Get("/transactions", controllers.)
	app.Put("/transactions/:id", controllers.UpdateTransaction)
}