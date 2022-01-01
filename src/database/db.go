package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(dsn string) {

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!")
	}
}
