package controllers

import (
	"github.com/adibSetiawann/transaction-api-go/model"
	"github.com/adibSetiawann/transaction-api-go/model/dto"
	"github.com/gofiber/fiber/v2"
)

func GetAllMerchant(c *fiber.Ctx) error {
	var merchants []model.MerchantResponse

	model.DB.Preload("Products").Find(&merchants)

	return c.JSON(fiber.Map{
		"merchant": merchants,
	})
}

func RegisterMerchant(c *fiber.Ctx) error {
	merchant := new(dto.CreateMerchantDto)

	if err := c.BodyParser(merchant); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"error": "can't handle request",
		})
	}

	newMerchant := model.Merchant{
		Name:     merchant.Name,
		Email:    merchant.Email,
		Address:  merchant.Address,
		Phone:    merchant.Phone,
	}

	if merchant.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"error": "field name is required",
		})
	}

	model.DB.Create(&newMerchant)

	return c.JSON(fiber.Map{
		"message": "create merchant success",
		"data":    merchant,
	})
}

func UpdateMerchant(ctx *fiber.Ctx) error {
	merchantReq := new(dto.UpdateMerchantDto)
	if err := ctx.BodyParser(merchantReq); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"error": "bad request",
		})
	}
	merchantId := ctx.Params("id")

	var merchant model.Merchant
	err := model.DB.First(&merchant, "id = ?", merchantId).Error
	if err != nil {
		return ctx.Status(404).JSON(fiber.Map{
			"error": "merchant not found",
		})
	}
	if merchantReq.Name != "" {
		merchant.Name = merchantReq.Name
	}
	if merchantReq.Address != "" {
		merchant.Address = merchantReq.Address
	}
	if merchantReq.Phone != "" {
		merchant.Phone = merchantReq.Phone
	}

	errUpdate := model.DB.Save(&merchant).Error
	if errUpdate != nil {
		return ctx.Status(500).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return ctx.JSON(fiber.Map{
		"message": "updated",
		"data":    merchant,
	})
}
