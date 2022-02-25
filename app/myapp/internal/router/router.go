package router

import "github.com/gofiber/fiber/v2"

// InitializeRouters initializes the router
func InitializeRouters(app *fiber.App) {
	// healthcheck
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("ok")
	})
}
