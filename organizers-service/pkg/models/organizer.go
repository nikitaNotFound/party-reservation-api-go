package models

type Organizer struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Description string `gorm:"not null" json:"description"`
	IsVerified  bool   `gorm:"not null" json:"is_verified"`
	Phone       string `gorm:"not null" json:"phone"`
}
