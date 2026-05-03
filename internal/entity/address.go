package entity

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	UserID   uint
	Detail   string
	City     string
	Province string
}
