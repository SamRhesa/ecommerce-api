package delivery

import (
	"ecommerce-api/internal/entity"
	"ecommerce-api/pkg/database"
	"ecommerce-api/pkg/jwt"
	"fmt"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func UploadProductImage(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	id := c.Params("id")

	var product entity.Product
	database.DB.First(&product, id)

	var store entity.Store
	database.DB.Where("user_id = ?", userID).First(&store)

	if product.StoreID != store.ID {
		return c.Status(403).JSON(fiber.Map{
			"error": "forbidden",
		})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "image required",
		})
	}

	filename := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), file.Filename)

	c.SaveFile(file, filename)

	product.Image = filename
	database.DB.Save(&product)

	return c.JSON(fiber.Map{
		"message": "upload success",
		"image":   filename,
	})
}

func GetTransactionDetail(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	id := c.Params("id")

	var transaction entity.Transaction

	if err := database.DB.
		Where("user_id = ?", userID).
		First(&transaction, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "transaction not found",
		})
	}

	var items []entity.TransactionItem
	database.DB.Where("transaction_id = ?", transaction.ID).Find(&items)

	return c.JSON(fiber.Map{
		"transaction": transaction,
		"items":       items,
	})
}

func GetTransactions(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))

	var transactions []entity.Transaction

	if err := database.DB.
		Where("user_id = ?", userID).
		Find(&transactions).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to get transactions",
		})
	}

	return c.JSON(fiber.Map{
		"data": transactions,
	})
}

func CreateTransaction(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))

	type ItemInput struct {
		ProductID uint `json:"product_id"`
		Qty       int  `json:"qty"`
	}

	type Request struct {
		AddressID uint        `json:"address_id"`
		Items     []ItemInput `json:"items"`
	}

	req := new(Request)

	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	var address entity.Address
	if err := database.DB.First(&address, req.AddressID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "address not found",
		})
	}

	if address.UserID != userID {
		return c.Status(403).JSON(fiber.Map{
			"error": "forbidden",
		})
	}

	total := 0

	tx := database.DB.Begin()

	transaction := entity.Transaction{
		UserID:    userID,
		AddressID: req.AddressID,
		Total:     0,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create transaction",
		})
	}

	for _, item := range req.Items {
		var product entity.Product

		if err := tx.First(&product, item.ProductID).Error; err != nil {
			tx.Rollback()
			return c.Status(404).JSON(fiber.Map{
				"error": "product not found",
			})
		}

		if product.Stock < item.Qty {
			tx.Rollback()
			return c.Status(400).JSON(fiber.Map{
				"error": "stock not enough",
			})
		}

		product.Stock -= item.Qty

		if err := tx.Save(&product).Error; err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to update stock",
			})
		}

		subtotal := product.Price * item.Qty
		total += subtotal

		transactionItem := entity.TransactionItem{
			TransactionID: transaction.ID,
			ProductID:     product.ID,
			Name:          product.Name,
			Price:         product.Price,
			Qty:           item.Qty,
		}

		if err := tx.Create(&transactionItem).Error; err != nil {
			tx.Rollback()
			return c.Status(500).JSON(fiber.Map{
				"error": "failed to create transaction item",
			})
		}
	}

	transaction.Total = total

	if err := tx.Save(&transaction).Error; err != nil {
		tx.Rollback()
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to update total",
		})
	}

	tx.Commit()

	return c.JSON(fiber.Map{
		"message": "transaction success",
		"data":    transaction,
	})
}

func DeleteProduct(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	id := c.Params("id")

	var product entity.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "product not found",
		})
	}

	// ambil store user
	var store entity.Store
	if err := database.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "store not found",
		})
	}

	//  validasi kepemilikan
	if product.StoreID != store.ID {
		return c.Status(403).JSON(fiber.Map{
			"error": "forbidden",
		})
	}

	if err := database.DB.Delete(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to delete product",
		})
	}

	return c.JSON(fiber.Map{
		"message": "product deleted",
	})
}

func UpdateProduct(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	id := c.Params("id")

	var product entity.Product
	if err := database.DB.First(&product, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "product not found",
		})
	}

	// ambil store milik user
	var store entity.Store
	if err := database.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "store not found",
		})
	}

	//  validasi kepemilikan
	if product.StoreID != store.ID {
		return c.Status(403).JSON(fiber.Map{
			"error": "forbidden",
		})
	}

	input := new(entity.Product)
	if err := c.BodyParser(input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	product.Name = input.Name
	product.Price = input.Price
	product.Stock = input.Stock
	product.CategoryID = input.CategoryID

	if err := database.DB.Save(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to update product",
		})
	}

	return c.JSON(fiber.Map{
		"message": "product updated",
		"data":    product,
	})
}

func GetProducts(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	categoryID := c.Query("category_id")
	name := c.Query("name")

	offset := (page - 1) * limit

	var products []entity.Product
	var total int64

	query := database.DB.Model(&entity.Product{})

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	query.Count(&total)

	if err := query.
		Limit(limit).
		Offset(offset).
		Find(&products).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to get products",
		})
	}

	return c.JSON(fiber.Map{
		"page":  page,
		"limit": limit,
		"total": total,
		"data":  products,
	})
}

func CreateProduct(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))

	// ambil store milik user
	var store entity.Store
	if err := database.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "store not found",
		})
	}

	product := new(entity.Product)

	if err := c.BodyParser(product); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	// paksa store_id dari user
	product.StoreID = store.ID

	if err := database.DB.Create(&product).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create product",
		})
	}

	return c.JSON(fiber.Map{
		"message": "product created",
		"data":    product,
	})
}

func DeleteCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	var category entity.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "category not found",
		})
	}

	if err := database.DB.Delete(&category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to delete category",
		})
	}

	return c.JSON(fiber.Map{
		"message": "category deleted",
	})
}

func UpdateCategory(c *fiber.Ctx) error {
	id := c.Params("id")

	var category entity.Category
	if err := database.DB.First(&category, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "category not found",
		})
	}

	input := new(entity.Category)
	if err := c.BodyParser(input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	category.Name = input.Name

	if err := database.DB.Save(&category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to update category",
		})
	}

	return c.JSON(fiber.Map{
		"message": "category updated",
		"data":    category,
	})
}

func GetCategories(c *fiber.Ctx) error {
	var categories []entity.Category

	if err := database.DB.Find(&categories).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to get categories",
		})
	}

	return c.JSON(fiber.Map{
		"data": categories,
	})
}

func CreateCategory(c *fiber.Ctx) error {
	category := new(entity.Category)

	if err := c.BodyParser(category); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	if err := database.DB.Create(&category).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create category",
		})
	}

	return c.JSON(fiber.Map{
		"message": "category created",
		"data":    category,
	})
}

func DeleteAddress(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	id := c.Params("id")

	var address entity.Address
	if err := database.DB.First(&address, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "address not found",
		})
	}

	//  validasi kepemilikan
	if address.UserID != userID {
		return c.Status(403).JSON(fiber.Map{
			"error": "forbidden",
		})
	}

	if err := database.DB.Delete(&address).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to delete address",
		})
	}

	return c.JSON(fiber.Map{
		"message": "address deleted",
	})
}

func UpdateAddress(c *fiber.Ctx) error {
	userID := uint(c.Locals("user_id").(float64))
	id := c.Params("id")

	var address entity.Address
	if err := database.DB.First(&address, id).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "address not found",
		})
	}

	//  validasi kepemilikan
	if address.UserID != userID {
		return c.Status(403).JSON(fiber.Map{
			"error": "forbidden",
		})
	}

	input := new(entity.Address)
	if err := c.BodyParser(input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	address.Detail = input.Detail
	address.City = input.City
	address.Province = input.Province

	if err := database.DB.Save(&address).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to update address",
		})
	}

	return c.JSON(fiber.Map{
		"message": "address updated",
		"data":    address,
	})
}

func GetAddresses(c *fiber.Ctx) error {
	userID := c.Locals("user_id")

	var addresses []entity.Address

	if err := database.DB.Where("user_id = ?", uint(userID.(float64))).Find(&addresses).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to get addresses",
		})
	}

	return c.JSON(fiber.Map{
		"data": addresses,
	})
}

func CreateAddress(c *fiber.Ctx) error {
	userID := c.Locals("user_id")

	address := new(entity.Address)

	if err := c.BodyParser(address); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	address.UserID = uint(userID.(float64))

	if err := database.DB.Create(&address).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create address",
		})
	}

	return c.JSON(fiber.Map{
		"message": "address created",
		"data":    address,
	})
}

func UpdateMyStore(c *fiber.Ctx) error {
	userID := c.Locals("user_id")

	var store entity.Store
	if err := database.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "store not found",
		})
	}

	input := new(entity.Store)
	if err := c.BodyParser(input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	store.Name = input.Name

	if err := database.DB.Save(&store).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to update store",
		})
	}

	return c.JSON(fiber.Map{
		"message": "store updated",
		"data":    store,
	})
}

func GetMyStore(c *fiber.Ctx) error {
	userID := c.Locals("user_id")

	var store entity.Store
	if err := database.DB.Where("user_id = ?", userID).First(&store).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "store not found",
		})
	}

	return c.JSON(store)
}

func UpdateProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id")

	var user entity.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	input := new(entity.User)
	if err := c.BodyParser(input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	// update field
	user.Name = input.Name
	user.Phone = input.Phone

	if err := database.DB.Save(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to update user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "update success",
		"data":    user,
	})
}

func GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id")

	var user entity.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		return c.Status(404).JSON(fiber.Map{
			"error": "user not found",
		})
	}

	return c.JSON(fiber.Map{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email,
		"phone": user.Phone,
		"role":  user.Role,
	})
}

func Login(c *fiber.Ctx) error {
	input := new(entity.User)

	if err := c.BodyParser(input); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	var user entity.User
	if err := database.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid email or password",
		})
	}

	token, err := jwt.GenerateToken(user.ID, user.Role)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to generate token",
		})
	}

	return c.JSON(fiber.Map{
		"message": "login success",
		"token":   token,
	})
}

func Register(c *fiber.Ctx) error {
	user := new(entity.User)

	if err := c.BodyParser(user); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid input",
		})
	}

	var existingUser entity.User
	if err := database.DB.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "email already used",
		})
	}

	if err := database.DB.Where("phone = ?", user.Phone).First(&existingUser).Error; err == nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "phone already used",
		})
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to hash password",
		})
	}

	user.Password = string(hashedPassword)

	if err := database.DB.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "failed to create user",
		})
	}

	store := entity.Store{
		Name:   user.Name + "'s Store",
		UserID: user.ID,
	}

	database.DB.Create(&store)

	return c.JSON(fiber.Map{
		"message": "register success",
	})
}
