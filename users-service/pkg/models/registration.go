package models

import "gorm.io/gorm"

type Registration struct {
	gorm.Model
	FirstName        string           `json:"firstName"`
	LastName         string           `json:"lastName"`
	Age              int              `json:"age"`
	RegistrationType RegistrationType `json:"registrationType "`
}

type RegistrationType string

const (
	Standard RegistrationType = "Standard"
	Premium                   = "Premium"
	Royal                     = "Royal"
)
