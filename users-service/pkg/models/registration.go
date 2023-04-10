package models

type Registration struct {
	ID      uint `gorm:"primary_key" json:"id"`
	UserID  uint `gorm:"not null" json:"user_id"`
	PartyID uint `gorm:"not null" json:"party_id"`
}
