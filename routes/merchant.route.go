package routes

import (
	"github.com/adibSetiawann/transaction-api-go/controller"
	"github.com/adibSetiawann/transaction-api-go/middleware"
	repository "github.com/adibSetiawann/transaction-api-go/repository/merchant"
	service "github.com/adibSetiawann/transaction-api-go/service/merchant"
	"github.com/gofiber/fiber/v2"
)

func MerchantRoute(app *fiber.App) {
	merchantRepo := repository.NewMerchantRepository()
	merchantService := service.NewMerchantService(&merchantRepo)
	merchantController := controller.NewMerchantController(&merchantService)
	
	app.Get("/merchants", merchantController.GetAll)
	app.Get("/merchants/:id", merchantController.GetById)
	app.Post("/merchants", middleware.AuthAsAdmin, merchantController.Create)
	app.Put("/merchants/update/:id", middleware.AuthAsAdmin, merchantController.Update)
	app.Delete("/merchants/delete/:id", middleware.AuthAsAdmin, merchantController.Remove)
}
