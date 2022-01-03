package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Connect() *gorm.DB {
	dsn := "root:password@tcp(127.0.0.1:3306)/yatta_h3i?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("unable to connect to database: %s", err.Error())
		panic("Could not connect with the database!")

	}

	return db
}
