package database

import (
	"fmt"

	"ecommerce-api/internal/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/ecommerce?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Gagal connect ke database")
	}

	fmt.Println("Database connected!")

	DB = db

	db.AutoMigrate(
		&entity.User{},
		&entity.Store{},
		&entity.Address{},
		&entity.Category{},
		&entity.Product{},
		&entity.Transaction{},
		&entity.TransactionItem{},
	)
}
