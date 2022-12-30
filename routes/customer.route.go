package routes

import (
	"github.com/adibSetiawann/transaction-api-go/controller"
	"github.com/adibSetiawann/transaction-api-go/middleware"
	repository "github.com/adibSetiawann/transaction-api-go/repository/customer"
	service "github.com/adibSetiawann/transaction-api-go/service/customer"
	"github.com/gofiber/fiber/v2"
)

func CustomerRoute(app *fiber.App) {
	customerRepo := repository.NewCustomerRepository()
	customerService := service.NewCustomerService(&customerRepo)
	customerController := controller.NewCustomerController(&customerService)

	app.Post("/register", customerController.Create)
	app.Post("/login", customerController.Login)
	app.Post("/logout", middleware.AuthForRegistered, customerController.Logout)
	app.Get("/customers", middleware.AuthAsAdmin, customerController.GetAll)
	app.Get("/customers/:id", middleware.AuthForRegistered, customerController.GetById)
	app.Put("/customers/update/:id", middleware.AuthAsCustomer, customerController.Update)
	app.Delete("/customers/delete/:id", middleware.AuthAsCustomer, customerController.Remove)
}
