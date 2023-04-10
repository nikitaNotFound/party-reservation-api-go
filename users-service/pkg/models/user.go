package models

type User struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	Name        string `gorm:"not null" json:"name"`
	Email       string `gorm:"not null;unique" json:"email"`
	PhoneNumber string `gorm:"not null" json:"phone_number"`
	Age         uint   `gorm:"not null" json:"age"`
}
