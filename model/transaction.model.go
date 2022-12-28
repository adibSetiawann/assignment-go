package model

type Transaction struct {
	ID         int64  `gorm:"primaryKey" json:"id"`
	Invoice    string `json:"invoice" binding:"required"`
	Qty        int64  `json:"qty"`
	StatusId   int    `json:"status_id"`
	Status     Status `json:"status"`
	CustomerId int    `json:"customer_id"`
	ProductId  int    `json:"product_id"`
}

type TransactionResponse struct {
	ID         int64            `gorm:"primaryKey" json:"id"`
	Invoice    string           `json:"invoice" binding:"required"`
	Qty        int64            `json:"qty"`
	StatusId   int              `json:"status_id"`
	Status     Status           `json:"status"`
	CustomerId int              `json:"customer_id"`
	ProductId  int              `json:"product_id"`
	Customer   CustomerResponse `json:"customer"`
	Product    Product          `json:"product"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
