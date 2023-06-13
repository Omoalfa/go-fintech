package validators

import (
	"errors"
	"reflect"

	"github.com/Omoalfa/go-fintech/database"
	"github.com/Omoalfa/go-fintech/database/models"
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

var (
	CreateUserValidator = validator.NewValidator()
	UpdateUserValidator = validator.NewValidator()
)

func emailExistValidatior(v interface{}, param string) error {
	st := reflect.ValueOf(v)
	if st.Kind() != reflect.String {
		return errors.New("invalid email")
	}
	db := database.GetDB()
	var user models.User
	result := db.Model(&models.User{}).Where("email = ?", st.String()).First(&user)

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("email already in use")
	}
	return nil
}

func UserValidators() {
	CreateUserValidator.SetTag("cuv")
	UpdateUserValidator.SetTag("uuv")
	CreateUserValidator.SetValidationFunc("eev", emailExistValidatior)
}
