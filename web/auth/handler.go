package auth

import "github.com/gofiber/fiber/v2"

func authMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func SetupRoutes(api fiber.Router) {
	auth := api.Group("/auth", authMiddleware)

	auth.Post("/key", GetKey) // key
}
