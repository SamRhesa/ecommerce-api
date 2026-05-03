package main

import (
	"ecommerce-api/pkg/database"

	"ecommerce-api/internal/delivery"
	"ecommerce-api/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	app.Get("/profile", middleware.AuthMiddleware, delivery.GetProfile)
	app.Post("/register", delivery.Register)
	app.Post("/login", delivery.Login)
	app.Put("/profile", middleware.AuthMiddleware, delivery.UpdateProfile)
	app.Get("/store", middleware.AuthMiddleware, delivery.GetMyStore)
	app.Put("/store", middleware.AuthMiddleware, delivery.UpdateMyStore)
	app.Post("/address", middleware.AuthMiddleware, delivery.CreateAddress)
	app.Get("/address", middleware.AuthMiddleware, delivery.GetAddresses)
	app.Put("/address/:id", middleware.AuthMiddleware, delivery.UpdateAddress)
	app.Delete("/address/:id", middleware.AuthMiddleware, delivery.DeleteAddress)
	app.Post("/category", middleware.AuthMiddleware, middleware.AdminMiddleware, delivery.CreateCategory)
	app.Get("/category", delivery.GetCategories)
	app.Put("/category/:id", middleware.AuthMiddleware, middleware.AdminMiddleware, delivery.UpdateCategory)
	app.Delete("/category/:id", middleware.AuthMiddleware, middleware.AdminMiddleware, delivery.DeleteCategory)
	app.Post("/product", middleware.AuthMiddleware, delivery.CreateProduct)
	app.Get("/product", delivery.GetProducts)
	app.Put("/product/:id", middleware.AuthMiddleware, delivery.UpdateProduct)
	app.Delete("/product/:id", middleware.AuthMiddleware, delivery.DeleteProduct)
	app.Post("/transaction", middleware.AuthMiddleware, delivery.CreateTransaction)
	app.Get("/transaction", middleware.AuthMiddleware, delivery.GetTransactions)
	app.Get("/transaction/:id", middleware.AuthMiddleware, delivery.GetTransactionDetail)
	app.Post("/product/:id/upload", middleware.AuthMiddleware, delivery.UploadProductImage)
	app.Static("/uploads", "./uploads")
	app.Listen(":3000")
}
