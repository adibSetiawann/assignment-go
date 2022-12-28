package model

type Transaction struct {
	ID         int64  `gorm:"primaryKey" json:"id"`
	Invoice    string `json:"invoice" validate:"required"`
	Qty        int64  `json:"qty"`
	StatusId   int    `json:"status_id"`
	Status     Status `json:"status"`
	CustomerId int    `json:"customer_id"`
	ProductId  int    `json:"product_id"`
}

type TransactionResponse struct {
	ID         int64                    `gorm:"primaryKey" json:"id"`
	Invoice    string                   `json:"invoice" validate:"required"`
	Qty        int64                    `json:"qty"`
	StatusId   int                      `json:"status_id"`
	Status     Status                   `json:"status"`
	CustomerId int                      `json:"customer_id"`
	ProductId  int                      `json:"product_id"`
	Customer   CustomerRelationResponse `json:"customer"`
	Product    ProductRelationResponse  `json:"product"`
}

func (TransactionResponse) TableName() string {
	return "transactions"
}
