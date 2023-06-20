package user_controller

import (
	"fmt"
	"time"

	"github.com/Omoalfa/go-fintech/api_response"
	"github.com/Omoalfa/go-fintech/database/models"
	my_utils "github.com/Omoalfa/go-fintech/utils"
	"github.com/Omoalfa/go-fintech/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *fiber.Ctx) error {
	c.Accepts("application/json")
	u := new(models.User)

	if err := c.BodyParser(u); err != nil {
		return api_response.ServerError(c)
	}

	err := validators.CreateUserValidator.Struct(u)
	if err != nil {
		fmt.Println(err)
		return api_response.BadRequest(c, err.Error())
	}

	if u.Username == "" {
		u.Username = u.FirstName + "_" + my_utils.RandStringBytes(5)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return api_response.ServerError(c)
	}

	u.Password = string(hash)

	fmt.Println(err, "validation")

	u.DBCreateUser()

	return api_response.SuccessCreated(c, u)
}

func GetUsers(c *fiber.Ctx) error {
	user := models.User{}

	users, err := models.DBGetUsers(&user)
	if err != nil {
		return api_response.ServerError(c)
	}

	fmt.Println(users)

	return api_response.Success(c, users)
}

func LoginUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return api_response.ServerError(c)
	}

	models.DBGetUserByEmail(user.Email, &user)

	claims := jwt.MapClaims{
		"ID":    user.ID,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println(err)
		return api_response.ServerError(c)
	}

	return api_response.Success(c, map[string]string{
		"email": user.Email,
		"token": t,
	})
}
