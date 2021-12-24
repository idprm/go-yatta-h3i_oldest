package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"waki.mobi/go-yatta-h3i/src/models"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/yatta_h3i"), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!")
	}

}

func AutoMigrate() {
	DB.AutoMigrate(models.User{})
}
