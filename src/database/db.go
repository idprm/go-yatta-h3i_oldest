package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func Connect() {
	dsn := "root:password@tcp(127.0.0.1:3306)/yatta_h3i?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		panic("Could not connect with the database!")
	}

	log.Println("Connected to database successfully")
	// DEBUG ON CONSOLE
	db.Logger = logger.Default.LogMode(logger.Info)
	// TODO: Add migrations

	Database = DbInstance{
		Db: db,
	}
}
