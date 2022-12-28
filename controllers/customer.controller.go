package controllers

import (
	"log"
	"time"

	"github.com/adibSetiawann/transaction-api-go/model"
	"github.com/adibSetiawann/transaction-api-go/model/dto"
	"github.com/adibSetiawann/transaction-api-go/utils"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Login(c *fiber.Ctx) error {
	loginForm := new(model.LoginForm)

	if err := c.BodyParser(loginForm); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "can't handle request",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(loginForm)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	var customer model.Customer
	err := model.DB.First(&customer, "email = ?", loginForm.Email).Error
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "wrong email",
		})
	}

	isValid := utils.ConfirmPassword(loginForm.Password, customer.Password)

	if !isValid {
		return c.Status(404).JSON(fiber.Map{
			"error": "wrong password",
		})
	}

	claims := jwt.MapClaims{}
	claims["name"] = customer.Name
	claims["email"] = customer.Email
	claims["exp"] = time.Now().Add(time.Minute * 1500).Unix()
	if customer.Name == "admin" {
		claims["role"] = "admin"
	} else {
		claims["role"] = "customer"
	}

	token, errToken := utils.GenerateToken(&claims)

	if errToken != nil {
		log.Println(errToken)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

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

	validate := validator.New()
	errValidate := validate.Struct(customer)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newCustomer := model.Customer{
		Name:     customer.Name,
		Email:    customer.Email,
		Password: customer.Password,
		Address:  customer.Address,
		Phone:    customer.Phone,
		GenderID: customer.GenderID,
	}
	hashedPassword, err := utils.HashPassword(customer.Password)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	newCustomer.Password = hashedPassword
	if customer.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "field name is required",
		})
	}

	model.DB.Create(&newCustomer)

	return c.JSON(fiber.Map{
		"message": "create customer success",
		"data":    newCustomer,
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

func GetById(ctx *fiber.Ctx) error {

	customerId := ctx.Params("id")

	var customer model.Customer
	err := model.DB.First(&customer, "id = ?", customerId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "customer not found",
		})
	}

	return ctx.JSON(fiber.Map{
		"data": customer,
	})
}
