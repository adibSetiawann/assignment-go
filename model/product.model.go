package model

type Product struct {
	ID           int64         `gorm:"primaryKey" json:"id"`
	Name         string        `json:"name" validate:"required"`
	Description  string        `json:"description" validate:"required"`
	Price        float64           `json:"price" validate:"required"`
	Stock        int           `json:"stock"`
	MerchantId   int           `json:"merchant_id"`
	Transactions []Transaction `json:"transactions"`
}

type ProductResponse struct {
	ID          int64                    `gorm:"primaryKey" json:"id"`
	Name        string                   `json:"name" validate:"required"`
	Description string                   `json:"description" validate:"required"`
	Price       float64                      `json:"price" validate:"required"`
	MerchantId  int                      `json:"merchant_id"`
	Merchant    MerchantRelationResponse `json:"merchants"`
	Stock       int                      `json:"stock"`
}

type ProductRelationResponse struct {
	ID          int64                    `gorm:"primaryKey" json:"id"`
	Name        string                   `json:"name" validate:"required"`
	Description string                   `json:"description" validate:"required"`
	Price       float64                      `json:"price" validate:"required"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductRelationResponse) TableName() string {
	return "products"
}