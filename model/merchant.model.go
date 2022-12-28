package model

type Merchant struct {
	ID       int64     `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name" validate:"required"`
	Address  string    `json:"address" validate:"required"`
	Email    string    `json:"email" validate:"required"`
	Phone    string    `json:"phone"`
	Products []Product `json:"products"`
}

type MerchantResponse struct {
	ID      int64  `gorm:"primaryKey" json:"id"`
	Name    string `json:"name" form:"name"`
	Address string `json:"address" form:"address"`
	Email   string `json:"email" form:"email"`
	Phone   string `json:"phone" form:"phone"`
}

type MerchantRelationResponse struct {
	ID   int64  `gorm:"primaryKey" json:"id"`
	Name string `json:"name" form:"name"`
}

func (MerchantResponse) TableName() string {
	return "merchants"
}

func (MerchantRelationResponse) TableName() string {
	return "merchants"
}
