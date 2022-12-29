package entity

import "gorm.io/gorm"

type Transaction struct {
	ID         int64          `gorm:"primaryKey" json:"id"`
	Invoice    string         `json:"invoice" validate:"required"`
	Qty        float64        `json:"qty"`
	StatusId   int            `json:"status_id"`
	Status     Status         `json:"status"`
	CustomerId int            `json:"customer_id"`
	ProductId  int            `json:"product_id"`
	DeletedAt  gorm.DeletedAt `json:"deleted_at"`
}
