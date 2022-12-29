package routes

import (
	"github.com/adibSetiawann/transaction-api-go/controller"
	"github.com/adibSetiawann/transaction-api-go/middleware"
	repository "github.com/adibSetiawann/transaction-api-go/repository/product"
	service "github.com/adibSetiawann/transaction-api-go/service/product"
	"github.com/gofiber/fiber/v2"
)

func ProductRoute(app *fiber.App) {
	productRepo := repository.NewProductRepository()
	productService := service.NewProductService(&productRepo)
	productController := controller.NewProductController(&productService)
	app.Get("/products", productController.GetAll)
	app.Get("/products/:id", productController.GetById)
	// app.Post("/products", middleware.AuthAsAdmin, productController.Create)
	app.Post("/products",  productController.Create)
	app.Put("/products/update/:id", middleware.AuthAsAdmin, productController.Update)
	app.Delete("/products/delete/:id",  productController.Remove)
}
