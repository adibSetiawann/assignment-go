package controllers

import (
	"github.com/adibSetiawann/transaction-api-go/model"
	"github.com/adibSetiawann/transaction-api-go/model/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func GetAllProduct(c *fiber.Ctx) error {
	var products []model.ProductResponse

	model.DB.Preload("Merchant").Find(&products)

	return c.JSON(fiber.Map{
		"product": products,
	})
}

func RegisterProduct(c *fiber.Ctx) error {
	product := new(dto.CreateProductDto)

	if err := c.BodyParser(product); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "can't handle request",
		})
	}

	newProduct := model.Product{
		Name:        product.Name,
		Description: product.Description,
		MerchantId:  product.MerchantId,
		Price:       product.Price,
		Stock:       product.Stock,
	}

	validate := validator.New()
	errValidate := validate.Struct(product)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	model.DB.Create(&newProduct)

	return c.JSON(fiber.Map{
		"message": "create product success",
		"data":    product,
	})
}

func UpdateProduct(ctx *fiber.Ctx) error {
	ProductReq := new(dto.UpdateProductDto)
	if err := ctx.BodyParser(ProductReq); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "bad request",
		})
	}
	ProductId := ctx.Params("id")

	var product model.Product
	err := model.DB.First(&product, "id = ?", ProductId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "product not found",
		})
	}
	if ProductReq.Name != "" {
		product.Name = ProductReq.Name
	}
	if ProductReq.Description != "" {
		product.Description = ProductReq.Description
	}
	if ProductReq.MerchantId != 0 {
		product.MerchantId = ProductReq.MerchantId
	}

	errUpdate := model.DB.Save(&product).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "updated",
		"data":    product,
	})
}
