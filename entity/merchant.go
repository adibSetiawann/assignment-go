package entity

import "gorm.io/gorm"

type Merchant struct {
	ID        int64          `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name" validate:"required"`
	Address   string         `json:"address" validate:"required"`
	Email     string         `json:"email" validate:"required"`
	Phone     string         `json:"phone"`
	Products  []Product      `json:"products"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
