package entity

type Status struct {
	ID          int    `gorm:"primaryKey" json:"id"`
	Description string `json:"description"`
}
