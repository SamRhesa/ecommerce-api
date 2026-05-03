package middleware

import "github.com/gofiber/fiber/v2"

func AdminMiddleware(c *fiber.Ctx) error {
	role := c.Locals("role")

	if role != "admin" {
		return c.Status(403).JSON(fiber.Map{
			"error": "admin only",
		})
	}

	return c.Next()
}
