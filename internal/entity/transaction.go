package entity

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	UserID    uint
	AddressID uint
	Total     int
}
