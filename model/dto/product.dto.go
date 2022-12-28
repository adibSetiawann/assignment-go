package dto

type CreateProductDto struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	MerchantId  int     `json:"merchant_id" form:"merchant_id" validate:"required"`
}

type UpdateProductDto struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	MerchantId  int    `json:"merchant_id" form:"merchant_id" validate:"required"`
}
