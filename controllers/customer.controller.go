package controllers

import (
	"github.com/adibSetiawann/transaction-api-go/model"
	"github.com/adibSetiawann/transaction-api-go/model/dto"
	"github.com/gofiber/fiber/v2"
)

func GetAllCustomer(c *fiber.Ctx) error {
	var customers []model.CustomerResponse

	model.DB.Preload("Gender").Find(&customers)

	return c.JSON(fiber.Map{
		"customer": customers,
	})
}

func RegisterCustomer(c *fiber.Ctx) error {
	customer := new(dto.CreateCustomerDto)

	if err := c.BodyParser(customer); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "can't handle request",
		})
	}

	newCustomer := model.Customer{
		Name:     customer.Name,
		Email:    customer.Email,
		Address:  customer.Address,
		Phone:    customer.Phone,
		GenderID: customer.GenderID,
	}

	if customer.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "field name is required",
		})
	}

	model.DB.Create(&newCustomer)

	return c.JSON(fiber.Map{
		"message": "create customer success",
		"data":    customer,
	})
}

func UpdateCustomer(ctx *fiber.Ctx) error {
	customerReq := new(dto.UpdateCustomerDto)
	if err := ctx.BodyParser(customerReq); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "bad request",
		})
	}
	customerId := ctx.Params("id")

	var customer model.Customer
	err := model.DB.First(&customer, "id = ?", customerId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "customer not found",
		})
	}
	if customerReq.Name != "" {
		customer.Name = customerReq.Name
	}
	if customerReq.Address != "" {
		customer.Address = customerReq.Address
	}
	if customerReq.Phone != "" {
		customer.Phone = customerReq.Phone
	}

	errUpdate := model.DB.Save(&customer).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "updated",
		"data":    customer,
	})
}
