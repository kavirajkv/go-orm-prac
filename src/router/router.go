package router

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	// Health Check
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	// Mount User Routes
	MountUser(app.Group("/user"))
}
