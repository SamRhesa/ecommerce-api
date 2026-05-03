package middleware

import (
	"ecommerce-api/pkg/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtlib "github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")

	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "missing token",
		})
	}

	// format: Bearer TOKEN
	tokenString := strings.Split(authHeader, " ")
	if len(tokenString) != 2 {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid token format",
		})
	}

	token, err := jwtlib.Parse(tokenString[1], func(token *jwtlib.Token) (interface{}, error) {
		return jwt.SECRET_KEY, nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"error": "invalid token",
		})
	}

	claims := token.Claims.(jwtlib.MapClaims)

	// simpan ke context
	c.Locals("user_id", claims["user_id"])
	c.Locals("role", claims["role"])

	return c.Next()
}
