package models

import "time"

type Party struct {
	ID          uint      `gorm:"primary_key" json:"id"`
	Name        string    `gorm:"not null" json:"name"`
	StartDate   time.Time `gorm:"not null" json:"start_date"`
	EndDate     time.Time `gorm:"not null" json:"end_date"`
	Description string    `json:"description"`
	CancelledAt time.Time `gorm:"null" json:"cancelled_at"`
}
