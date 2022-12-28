package dto

type CreateMerchantDto struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

type UpdateMerchantDto struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UpdateMerchantEmailDto struct {
	Email string `json:"email" validate:"required"`
}

// gorm:"unique;not null"
