package routes

import (
	"github.com/adibSetiawann/transaction-api-go/controller"
	"github.com/adibSetiawann/transaction-api-go/middleware"
	repository "github.com/adibSetiawann/transaction-api-go/repository/transaction"
	service "github.com/adibSetiawann/transaction-api-go/service/transaction"
	"github.com/gofiber/fiber/v2"
)

func TransactionRoute(app *fiber.App) {
	transactionRepo := repository.NewTransactionRepository()
	transactionService := service.NewTransactionService(&transactionRepo)
	transactionController := controller.NewTransactionController(&transactionService)
	
	app.Get("/transactions", middleware.AuthAsAdmin, transactionController.GetAll)
	app.Get("/transactions/:invoice", middleware.AuthForRegistered, transactionController.GetByInvoice)
	app.Post("/transactions", middleware.AuthAsCustomer, transactionController.Create)
	app.Put("/transactions/update/:invoice", middleware.AuthAsCustomer, transactionController.Update)
	app.Delete("/transactions/delete/:invoice", middleware.AuthAsCustomer, transactionController.Remove)
}
