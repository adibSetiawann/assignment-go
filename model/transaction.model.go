package model

import "gorm.io/gorm"

type CreateTransaction struct {
	CustomerId int     `json:"customer_id" validate:"required"`
	ProductId  int     `json:"product_id" validate:"required"`
	Qty        float64 `json:"qty" validate:"required"`
	GrandTotal float64 `json:"grand_total"`
}

type UpdateTransaction struct {
	Qty       float64 `json:"qty" validate:"required"`
	ProductId int     `json:"product_id" validate:"required"`
}
type TransactionResponse struct {
	ID         int64                    `gorm:"primaryKey" json:"id"`
	Invoice    string                   `json:"invoice"`
	StatusId   int                      `json:"status_id"`
	CustomerId int                      `json:"customer_id"`
	ProductId  int                      `json:"product_id"`
	Qty        float64                  `json:"qty"`
	GrandTotal float64                  `json:"grand_total"`
	Product    ProductRelationResponse  `json:"product"`
	Status     Status                   `json:"status"`
	Customer   CustomerRelationResponse `json:"customer"`
	DeletedAt  gorm.DeletedAt           `json:"deleted_at"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
