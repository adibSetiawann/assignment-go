package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          int            `gorm:"primaryKey" form:"id" json:"id"`
	Name        string         `json:"name" binding:"required"`
	Email       string         `json:"email" binding:"required"`
	Password    string         `json:"password"`
	Phone       string         `json:"phone"`
	Address     string         `json:"address"`
	GenderID    int            `json:"gender_id" form:"gender_id" binding:"required"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Transaction []Transaction  `json:"transactions"`
}

type CustomerResponse struct {
	ID       int    `gorm:"primaryKey" form:"id" json:"id"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	GenderID int    `json:"gender_id" form:"gender_id" binding:"required"`
	Gender   Gender `json:"genders"`
}

type CustomerRelationResponse struct {
	ID   int    `gorm:"primaryKey" form:"id" json:"id"`
	Name string `json:"name" binding:"required"`
}

func (CustomerResponse) TableName() string {
	return "customers"
}
