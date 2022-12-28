package dto

type CreateCustomerDto struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
	GenderID int    `json:"gender_id" form:"gender_id" binding:"required"`
}

type UpdateCustomerDto struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

type UpdateCustomerEmailDto struct {
	Email string `json:"email" validate:"required"`
}

// gorm:"unique;not null"
