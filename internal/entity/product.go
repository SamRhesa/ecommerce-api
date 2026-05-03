package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name       string `json:"name"`
	Price      int    `json:"price"`
	Stock      int    `json:"stock"`
	StoreID    uint   `json:"store_id"`
	CategoryID uint   `json:"category_id"`
	Image      string `json:"image"`
}
