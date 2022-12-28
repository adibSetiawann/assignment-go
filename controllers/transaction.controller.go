package controllers

import (
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/adibSetiawann/transaction-api-go/model"
	"github.com/adibSetiawann/transaction-api-go/model/dto"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const charset = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func GenerateInvoice() string {
	rndmString := strings.ToUpper(String(14))
	year, month, day := time.Now().Date()

	var str strings.Builder
	str.WriteString("TRSC-")
	str.WriteString(strconv.Itoa(int(year)))
	str.WriteString(strconv.Itoa(int(month)))
	str.WriteString(strconv.Itoa(day))
	str.WriteString(rndmString)

	return str.String()
}

func GetAllTransaction(c *fiber.Ctx) error {
	var transactions []model.TransactionResponse

	model.DB.Preload("Customer").Preload("Product").Preload("Status").Find(&transactions)

	return c.JSON(fiber.Map{
		"transaction": transactions,
	})
}

func CreateTransaction(c *fiber.Ctx) error {
	transaction := new(dto.CreateTransactionDto)

	if err := c.BodyParser(transaction); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "can't handle request",
		})
	}

	validate := validator.New()
	errValidate := validate.Struct(transaction)
	if errValidate != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "failed",
			"error":   errValidate.Error(),
		})
	}

	newTransaction := model.Transaction{
		StatusId:   1,
		Qty:        transaction.Qty,
		CustomerId: transaction.CustomerId,
		ProductId:  transaction.ProductId,
		Invoice:    GenerateInvoice(),
	}

	if transaction.Qty < 1 {
		return c.Status(400).JSON(fiber.Map{
			"error": "qty is more than 0 required",
		})
	}

	model.DB.Create(&newTransaction)

	return c.JSON(fiber.Map{
		"message": "create transaction success",
		"data":    transaction,
	})
}

func UpdateTransaction(ctx *fiber.Ctx) error {
	TransactionReq := new(dto.UpdateTransactionDto)
	if err := ctx.BodyParser(TransactionReq); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "bad request",
		})
	}
	TransactionId := ctx.Params("id")

	var transaction model.Transaction
	err := model.DB.First(&transaction, "id = ?", TransactionId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "transaction not found",
		})
	}
	if TransactionReq.Qty != 0 {
		transaction.Qty = TransactionReq.Qty
	}

	errUpdate := model.DB.Save(&transaction).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "updated",
		"data":    transaction,
	})
}
