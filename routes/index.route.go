package routes

import (
	"github.com/adibSetiawann/transaction-api-go/controllers"
	"github.com/adibSetiawann/transaction-api-go/middleware"
	"github.com/gofiber/fiber/v2"
)

func RouteInit(app *fiber.App)  {
	// customer controller
	app.Post("/login", controllers.Login)
	app.Get("/customers", middleware.Auth, controllers.GetAllCustomer)
	app.Post("/customers", controllers.RegisterCustomer)
	// app.Get("/customers", controllers.)
	app.Put("/customers/:id", middleware.Auth, controllers.UpdateCustomer)

	// product controller
	app.Get("/products", controllers.GetAllProduct)
	app.Post("/products", controllers.RegisterProduct)
	// app.Get("/products", controllers.)
	app.Put("/products/:id", controllers.UpdateProduct)

	// merchant controller
	app.Get("/merchants",middleware.Auth, controllers.GetAllMerchant)
	app.Post("/merchants", controllers.RegisterMerchant)
	// app.Get("/merchants", controllers.)
	app.Put("/merchants/:id", middleware.Auth, controllers.UpdateMerchant)

	// transaction controller
	app.Get("/transactions",middleware.Auth, controllers.GetAllTransaction)
	app.Post("/transactions", middleware.Auth, controllers.CreateTransaction)
	// app.Get("/transactions", controllers.)
	app.Put("/transactions/:id", middleware.Auth, controllers.UpdateTransaction)
}