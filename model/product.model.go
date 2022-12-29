package model

import "gorm.io/gorm"

type CreateProduct struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	MerchantId  int     `json:"merchant_id" form:"merchant_id" validate:"required"`
}

type UpdateProduct struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

type ProductResponse struct {
	ID          int64                    `gorm:"primaryKey" json:"id"`
	Name        string                   `json:"name"`
	Description string                   `json:"description"`
	Price       float64                  `json:"price"`
	Stock       int                      `json:"stock"`
	MerchantId  int                      `json:"merchant_id"`
	Merchant    MerchantRelationResponse `json:"merchants"`
	DeletedAt   gorm.DeletedAt           `json:"deleted_at"`
}

type ProductRelationResponse struct {
	ID          int64   `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
}

func (ProductResponse) TableName() string {
	return "products"
}

func (ProductRelationResponse) TableName() string {
	return "products"
}
