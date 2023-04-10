package models

type Organizer struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Description string `json:"description"`
}
