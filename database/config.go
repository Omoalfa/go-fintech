package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB
var dsn = "host=localhost user=omoalfa password=OdunAyoMi dbname=fintech port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func ConnectDB() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Unable to connect to DB")
		return nil
	}

	fmt.Println("Successfully connect to database!")

	Db = db
	return db
}

func GetDB() *gorm.DB {
	return Db
}
