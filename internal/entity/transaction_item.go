package entity

import "gorm.io/gorm"

type TransactionItem struct {
	gorm.Model
	TransactionID uint
	ProductID     uint
	Name          string
	Price         int
	Qty           int
}
