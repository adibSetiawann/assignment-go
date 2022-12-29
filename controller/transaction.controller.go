package controller

import (
	"github.com/adibSetiawann/transaction-api-go/model"
	"github.com/adibSetiawann/transaction-api-go/service/transaction"
	"github.com/gofiber/fiber/v2"
)

type TransactionController struct {
	transactionService service.TransactionService
}

func NewTransactionController(transactionService *service.TransactionService) TransactionController {
	return TransactionController{transactionService: *transactionService}
}

func (mc *TransactionController) Create(c *fiber.Ctx) error {
	transactionReq := new(model.CreateTransaction)

	if err := c.BodyParser(transactionReq); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "some field is required or wrong",
		})
	}
	isErrorValidation := mc.transactionService.Validation(*transactionReq)
	if isErrorValidation != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": isErrorValidation.Error(),
		})
	}


	transaction, errCreate := mc.transactionService.Create(*transactionReq)
	if errCreate != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": errCreate.Error(),
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "create data successfully",
		"data":    transaction,
	})
}

func (mc *TransactionController) Update(c *fiber.Ctx) error {
	invoice := c.Params("invoice")
	transactionReq := new(model.UpdateTransaction)

	if err := c.BodyParser(transactionReq); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "request can't go on",
		})
	}

	isErrorValidation := mc.transactionService.Validation(*transactionReq)
	if isErrorValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": isErrorValidation.Error(),
		})
	}

	transaction, errCreate := mc.transactionService.Update(invoice, *transactionReq)

	if errCreate != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "transaction not found",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "update data successfully",
		"data":    transaction,
	})
}

func (mc *TransactionController) GetByInvoice(c *fiber.Ctx) error {
	invoice := c.Params("invoice")

	transactions, err := mc.transactionService.GetByInvoice(invoice)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "transaction not found in database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": transactions,
	})
}

func (mc *TransactionController) GetAll(c *fiber.Ctx) error {
	transactions, _ := mc.transactionService.GetAllData()

	return c.Status(200).JSON(fiber.Map{
		"data": transactions,
	})
}

func (mc *TransactionController) Remove(c *fiber.Ctx) error {
	invoice := c.Params("invoice")
	err := mc.transactionService.Remove(invoice)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "transaction not found in database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "delete success",
	})
}

func (mc *TransactionController) CancelPayment(c *fiber.Ctx) error {
	invoice := c.Params("invoice")
	err := mc.transactionService.CancelPayment(invoice)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "transaction not found in database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "transaction status canceled",
	})
}

func (mc *TransactionController) SuccessPayment(c *fiber.Ctx) error {
	invoice := c.Params("invoice")
	err := mc.transactionService.SuccessPayment(invoice)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "transaction not found in database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "transaction status success",
	})
}