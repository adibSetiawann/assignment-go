package controller

import (
	"strconv"

	"github.com/adibSetiawann/transaction-api-go/model"
	"github.com/adibSetiawann/transaction-api-go/service/product"
	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	productService service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{productService: *productService}
}

func (mc *ProductController) Create(c *fiber.Ctx) error {
	productReq := new(model.CreateProduct)

	if err := c.BodyParser(productReq); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "request can't go on",
		})
	}

	isErrorValidation := mc.productService.Validation(*productReq)
	if isErrorValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": isErrorValidation.Error(),
		})
	}

	product, errCreate := mc.productService.Create(*productReq)
	if errCreate != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":  "create data successfully",
		"data": product,
	})
}

func (mc *ProductController) Update(c *fiber.Ctx) error {
	productId := c.Params("id")
	productReq := new(model.UpdateProduct)

	if err := c.BodyParser(productReq); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "request can't go on",
		})
	}

	isErrorValidation := mc.productService.Validation(*productReq)
	if isErrorValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": isErrorValidation.Error(),
		})
	}

	intId, _ := strconv.ParseInt(productId, 10, 64)
	product, errCreate := mc.productService.Update(intId, *productReq)
	
	if errCreate != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "product not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "update data successfully",
		"data":    product,
	})
}

func (mc *ProductController) GetById(c *fiber.Ctx) error {
	productId := c.Params("id")

	products, err := mc.productService.GetById(productId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "product not found in database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": products,
	})
}

func (mc *ProductController) GetAll(c *fiber.Ctx) error {
	products, _ := mc.productService.GetAllData()

	return c.Status(200).JSON(fiber.Map{
		"data": products,
	})
}

func (mc *ProductController) Remove(c *fiber.Ctx) error {
	productId := c.Params("id")
	err := mc.productService.Remove(productId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "product not found in database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "delete success",
	})
}
