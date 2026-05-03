package entity

import "gorm.io/gorm"

type Store struct {
	gorm.Model
	Name   string
	UserID uint
}
