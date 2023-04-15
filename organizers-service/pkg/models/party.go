package models

import "time"

type Party struct {
	ID          uint       `gorm:"primaryKey" json:"id"`
	Name        string     `gorm:"not null" json:"name"`
	StartDate   time.Time  `gorm:"not null" json:"start_date"`
	EndDate     time.Time  `gorm:"not null" json:"end_date"`
	Description string     `gorm:"not null" json:"description"`
	CancelledAt *time.Time `json:"cancelled_at"`
	IsDraft     bool       `gorm:"not null" json:"is_draft"`
	IsApproved  bool       `gorm:"not null" json:"is_approved"`
	IsPublished bool       `gorm:"not null" json:"is_published"`
}
