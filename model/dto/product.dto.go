package dto

type CreateProductDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Stock       int    `json:"stock" binding:"required"`
	MerchantId  int    `json:"merchant_id" form:"merchant_id" binding:"required"`
}

type UpdateProductDto struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	MerchantId  int    `json:"merchant_id" form:"merchant_id" binding:"required"`
}
