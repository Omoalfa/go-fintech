package routes

import (
	user_controller "github.com/Omoalfa/go-fintech/controller"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func SetupUserRoutes(app *fiber.App) {
	userRoutes := fiber.New()

	userRoutes.Post("/verify", user_controller.VerifyEmail)
	userRoutes.Get("/verify", user_controller.ResendVerificationMail)
	userRoutes.Get("/", user_controller.GetUsers)
	userRoutes.Post("/", user_controller.CreateUser)
	userRoutes.Post("/login", user_controller.LoginUser)
	userRoutes.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
		ErrorHandler: func(c *fiber.Ctx, err error) error {

			return c.Status(fiber.StatusUnauthorized).JSON(map[string]any{
				"status":  fiber.StatusUnauthorized,
				"message": "Unauthorized user",
				"errors":  err.Error(),
			})
		},
	}))
	userRoutes.Get("/protected", func(c *fiber.Ctx) error {
		user := c.Locals("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		return c.SendString(email)
	})

	app.Mount("api/v1/users", userRoutes)
}
