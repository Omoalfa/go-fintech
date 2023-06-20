package models

import (
	"github.com/Omoalfa/go-fintech/database"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email      string `json:"email" cuv:"required,email,eev"`
	Phone      string `json:"phone" cuv:"required,e164,pev"`
	FirstName  string `json:"firstName" cuv:"required"`
	LastName   string `json:"lastName" cuv:"required"`
	Avatar     string `json:"avatar"`
	Password   string `json:"password" cuv:"required,min=8"`
	IsVerified *bool  `json:"isVerified" gorm:"column:isVerified,default:false"`
	Username   string `json:"username" gorm:"unique,column:username" cuv:"uev"`
}

func (b *User) DBCreateUser() *User {
	db := database.GetDB()
	db.Create(b)
	return b
}

func DBUpdateUser(id int, user *User) (*User, error) {
	db := database.GetDB()
	tx := db.Where("ID = ?", id).Updates(user)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return user, nil
}

func DBDeleteUser(id int) error {
	db := database.GetDB()
	tx := db.Delete(&User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func DBGetUsers(query *User) ([]User, error) {
	db := database.GetDB()
	var users []User
	tx := db.Where(query).Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return users, nil
}

func DBGetUserByEmail(email string, user *User) {
	db := database.GetDB()
	db.Where("email = ?", email).First(user)
}
