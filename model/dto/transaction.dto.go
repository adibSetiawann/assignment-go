package dto

type CreateTransactionDto struct {
	// TransactionCode string           `json:"transaction_code" binding:"required"`
	Qty       int64 `json:"qty"`
	CustomerId int   `json:"customer_id"`
	ProductId  int   `json:"product_id"`
}

type UpdateTransactionDto struct {
	// TransactionCode string           `json:"transaction_code" binding:"required"`
	Qty       int64 `json:"qty"`
	CustomerId int   `json:"customer_id"`
	ProductId  int   `json:"product_id"`
}
