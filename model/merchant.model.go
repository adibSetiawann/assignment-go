package model

import "gorm.io/gorm"

type CreateMerchant struct {
	ID      int64  `gorm:"primaryKey" json:"id"`
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}

type UpdateMerchant struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}

type UpdateMerchantEmail struct {
	Email string `json:"email" validate:"required"`
}

type MerchantResponse struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name" form:"name"`
	Address   string         `json:"address" form:"address"`
	Email     string         `json:"email" form:"email"`
	Phone     string         `json:"phone" form:"phone"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type MerchantRelationResponse struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `json:"name" form:"name"`
}

func (MerchantResponse) TableName() string {
	return "merchants"
}

func (MerchantRelationResponse) TableName() string {
	return "merchants"
}
