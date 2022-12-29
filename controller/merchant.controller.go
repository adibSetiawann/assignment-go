package controller

import (
	"strconv"

	"github.com/adibSetiawann/transaction-api-go/model"
	"github.com/adibSetiawann/transaction-api-go/service/merchant"
	"github.com/gofiber/fiber/v2"
)

type MerchantController struct {
	merchantService service.MerchantService
}

func NewMerchantController(merchantService *service.MerchantService) MerchantController {
	return MerchantController{merchantService: *merchantService}
}

func (mc *MerchantController) Create(c *fiber.Ctx) error {
	merchantReq := new(model.CreateMerchant)

	if err := c.BodyParser(merchantReq); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "request can't go on",
		})
	}

	isErrorValidation := mc.merchantService.Validation(*merchantReq)
	if isErrorValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": isErrorValidation.Error(),
		})
	}

	merchant, errCreate := mc.merchantService.Create(*merchantReq)
	if errCreate != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "internal server error",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message":  "create data successfully",
		"merchant": merchant,
	})
}

func (mc *MerchantController) Update(c *fiber.Ctx) error {
	merchantId := c.Params("id")
	merchantReq := new(model.UpdateMerchant)

	if err := c.BodyParser(merchantReq); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "request can't go on",
		})
	}

	isErrorValidation := mc.merchantService.Validation(*merchantReq)
	if isErrorValidation != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": isErrorValidation.Error(),
		})
	}

	intId, _ := strconv.ParseInt(merchantId, 10, 64)
	merchant, errCreate := mc.merchantService.Update(intId, *merchantReq)

	if errCreate != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "merchant not found in database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "update data successfully",
		"data":    merchant,
	})
}

func (mc *MerchantController) GetById(c *fiber.Ctx) error {
	merchantId := c.Params("id")

	merchants, err := mc.merchantService.GetById(merchantId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "merchant not found in database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"data": merchants,
	})
}

func (mc *MerchantController) GetAll(c *fiber.Ctx) error {
	merchants, _ := mc.merchantService.GetAllData()

	return c.Status(200).JSON(fiber.Map{
		"data": merchants,
	})
}

func (mc *MerchantController) Remove(c *fiber.Ctx) error {
	merchantId := c.Params("id")
	err := mc.merchantService.Remove(merchantId)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"message": "merchant not found in database",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"message": "delete success",
	})
}
