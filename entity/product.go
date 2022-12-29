package entity

import "gorm.io/gorm"

type Product struct {
	ID           int64          `gorm:"primaryKey" json:"id"`
	Name         string         `json:"name" validate:"required"`
	Description  string         `json:"description" validate:"required"`
	Price        float64        `json:"price" validate:"required"`
	Stock        int            `json:"stock"`
	MerchantId   int            `json:"merchant_id"`
	Transactions []Transaction  `json:"transactions"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at"`
}
