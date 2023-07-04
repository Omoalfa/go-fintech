package main

import (
	"fmt"
	"log"

	"github.com/Omoalfa/go-fintech/config"
	"github.com/Omoalfa/go-fintech/database"
	"github.com/Omoalfa/go-fintech/database/models"
	"github.com/Omoalfa/go-fintech/routes"
	"github.com/Omoalfa/go-fintech/validators"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//initialize user validators:::
	validators.UserValidators()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Go Fintec!")
	})

	config.SetUpConfig()
	fmt.Println(config.Config)
	// fmt.Println(config.Test)
	database.ConnectDB()

	db := database.GetDB()
	// if db != nil {
	//run auto migrate:::
	db.AutoMigrate(&models.User{})

	//Setup routes:::
	routes.SetupUserRoutes(app)

	//set up app in here::::
	app.Listen(":3000")
	// }

	log.Fatal("Unable to connect to DB")
}
