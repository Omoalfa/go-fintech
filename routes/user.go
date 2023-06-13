package routes

import "github.com/gofiber/fiber/v2"

func SetupUserRoutes(app *fiber.App) {
	userRoutes := fiber.New()

	userRoutes.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go User Routes!")
	})

	app.Mount("api/v1/users", userRoutes)
}
