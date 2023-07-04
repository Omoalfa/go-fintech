package validators

import (
	"errors"
	"fmt"

	"github.com/Omoalfa/go-fintech/database"
	"github.com/Omoalfa/go-fintech/database/models"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

var (
	CreateUserValidator      = validator.New()
	UpdateUserValidator      = validator.New()
	ValidateVerificationCode = validator.New()
)

func emailExistValidation(fl validator.FieldLevel) bool {
	db := database.GetDB()
	var user models.User
	result := db.Model(&models.User{}).Where("email = ?", fl.Field()).First(&user)
	return errors.Is(result.Error, gorm.ErrRecordNotFound)
}

func emailExistValidationForValidation(fl validator.FieldLevel) bool {
	db := database.GetDB()
	var user models.User
	result := db.Model(&models.User{}).Where("email = ?", fl.Field()).First(&user)
	fmt.Println(result)
	return !errors.Is(result.Error, gorm.ErrRecordNotFound)
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

func verifyEmailPin(fl validator.FieldLevel) bool {
	email := fl.Parent().Interface().(*models.User)
	var user models.User

	models.DBGetUserByEmail(email.Email, &user)
	return user.VerificationCode == fl.Field().Interface()
}

func UserValidators() {
	CreateUserValidator.SetTagName("cuv")
	UpdateUserValidator.SetTagName("uuv")
	ValidateVerificationCode.SetTagName("vvc")
	CreateUserValidator.RegisterValidation("eev", emailExistValidation)
	ValidateVerificationCode.RegisterValidation("eev", emailExistValidationForValidation)
	ValidateVerificationCode.RegisterValidation("vp", verifyEmailPin)
	CreateUserValidator.RegisterValidation("pev", phoneExistValidation)
	CreateUserValidator.RegisterValidation("uev", usernameExistValidation)
	UpdateUserValidator.RegisterValidation("eev", emailExistValidation)
	UpdateUserValidator.RegisterValidation("pev", phoneExistValidation)
	UpdateUserValidator.RegisterValidation("uev", usernameExistValidation)

}
