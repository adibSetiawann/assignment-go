package model

type Gender struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Description string `json:"description"`
}
