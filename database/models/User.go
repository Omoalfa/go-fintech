package models

import (
	"fmt"

	"github.com/Omoalfa/go-fintech/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string `json:"email" gorm:"unique"`
	Phone      string `json:"string" gorm:"unique"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Avatar     string `json:"avatar"`
	Password   string `json:"password"`
	IsVerified *bool  `json:"isVerified"`
}

var db = database.GetDB()

func (u *User) BeforeSave() (err error) {
	fmt.Println("before save")
	fmt.Println(u.Password)
	if u.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
		if err != nil {
			return err
		}
		u.Password = string(hash)
	}
	return
}

func DBCreateUser(user *User) *User {
	db.Create(&user)
	return user
}

func DBUpdateUser(id int, user *User) *User {
	db.Where("ID = ?", id).Updates(&user)
	return user
}

func DBDeleteUser(id int) {
	db.Delete(&User{}, id)
}

func DBGetUsers(query *User) []User {
	var users []User
	db.Where(query).Find(&users)
	return users
}
