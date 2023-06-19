package controllers

import "gorm.io/gorm"

type UserServiceHandler struct {
	Db *gorm.DB
}
