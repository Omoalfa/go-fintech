package validators

import (
	"errors"

	"github.com/Omoalfa/go-fintech/database"
	"github.com/Omoalfa/go-fintech/database/models"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

var (
	CreateUserValidator = validator.New()
	UpdateUserValidator = validator.New()
)

func emailExistValidation(fl validator.FieldLevel) bool {
	db := database.GetDB()
	var user models.User
	result := db.Model(&models.User{}).Where("email = ?", fl.Field()).First(&user)
	return errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func phoneExistValidation(fl validator.FieldLevel) bool {
	db := database.GetDB()
	var user models.User
	result := db.Model(&models.User{}).Where("phone = ?", fl.Field()).First(&user)
	return errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func usernameExistValidation(fl validator.FieldLevel) bool {
	val := fl.Field()
	if val.String() == "" {
		return true
	}
	db := database.GetDB()
	var user models.User
	result := db.Model(&models.User{}).Where("username = ?", fl.Field()).First(&user)
	return errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func UserValidators() {
	CreateUserValidator.SetTagName("cuv")
	UpdateUserValidator.SetTagName("uuv")
	CreateUserValidator.RegisterValidation("eev", emailExistValidation)
	CreateUserValidator.RegisterValidation("pev", phoneExistValidation)
	CreateUserValidator.RegisterValidation("uev", usernameExistValidation)
	UpdateUserValidator.RegisterValidation("eev", emailExistValidation)
	UpdateUserValidator.RegisterValidation("pev", phoneExistValidation)
	UpdateUserValidator.RegisterValidation("uev", usernameExistValidation)

}
