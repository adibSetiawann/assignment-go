package model

type Product struct {
	ID           int64         `gorm:"primaryKey" json:"id"`
	Name         string        `json:"name" binding:"required"`
	Description  string        `json:"description" binding:"required"`
	Stock        int           `json:"stock"`
	MerchantId   int           `json:"merchant_id"`
	Transactions []Transaction `json:"transactions"`
}

type ProductResponse struct {
	ID          int64                    `gorm:"primaryKey" json:"id"`
	Name        string                   `json:"name" binding:"required"`
	Description string                   `json:"description" binding:"required"`
	MerchantId  int                      `json:"merchant_id"`
	Merchant    MerchantRelationResponse `json:"merchants"`
	Stock       int                      `json:"stock"`
}

func (ProductResponse) TableName() string {
	return "products"
}
