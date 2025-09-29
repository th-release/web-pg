package web

import (
	"cth.release/web-pg/web/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type ServerConfig struct {
	App *fiber.App
}

func InitServer() *ServerConfig {
	app := fiber.New()

	server := &ServerConfig{
		App: app,
	}

	server.SetupRoutes(app)
	return server
}

func (s *ServerConfig) SetupRoutes(app *fiber.App) *fiber.App {
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE",
		AllowHeaders: "Origin, Content-Type, Database-Auth",
	}))

	api := app.Group("/api", EmptyMiddleware)
	api.Get("/health", HealthHandler)

	auth.SetupRoutes(api)

	return app
}

func EmptyMiddleware(c *fiber.Ctx) error {
	return c.Next()
}

func HealthHandler(c *fiber.Ctx) error {
	return c.Status(200).SendString("200")
}
