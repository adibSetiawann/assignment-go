package entity

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID          int            `gorm:"primaryKey" form:"id" json:"id"`
	Name        string         `json:"name" validate:"required"`
	Email       string         `json:"email" validate:"required"`
	Password    string         `json:"-" gorm:"column:password"`
	Phone       string         `json:"phone"`
	Address     string         `json:"address"`
	GenderID    int            `json:"gender_id" form:"gender_id" validate:"required"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
	Transaction []Transaction  `json:"transactions"`
}