package model

import "gorm.io/gorm"

type CreateCustomer struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	Address  string `json:"address" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	GenderID int    `json:"gender_id" form:"gender_id" validate:"required"`
}

type UpdateCustomer struct {
	Name    string `json:"name" validate:"required"`
	Address string `json:"address" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
}

type UpdateCustomerEmail struct {
	Email string `json:"email" validate:"required"`
}

type CustomerResponse struct {
	ID        int            `gorm:"primaryKey" form:"id" json:"id"`
	Name      string         `json:"name"`
	Email     string         `json:"email"`
	Phone     string         `json:"phone"`
	Address   string         `json:"address"`
	GenderID  int            `json:"gender_id" form:"gender_id"`
	Gender    Gender         `json:"genders"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type CustomerRelationResponse struct {
	ID   int    `gorm:"primaryKey" form:"id" json:"id"`
	Name string `json:"name" validate:"required"`
}

type LoginForm struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (CustomerResponse) TableName() string {
	return "customers"
}

func (CustomerRelationResponse) TableName() string {
	return "customers"
}
