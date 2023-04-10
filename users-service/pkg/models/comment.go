package models

type Comment struct {
	ID      uint   `gorm:"primary_key" json:"id"`
	Message string `gorm:"not null" json:"message"`
	UserID  uint   `gorm:"not null" json:"user_id"`
	PartyID uint   `gorm:"not null" json:"party_id"`
}
