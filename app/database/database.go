package database

import (
	"app/database/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func Initialize() (*gorm.DB, error) {
	dbConfig := os.Getenv("DB_CONFIG")
	fmt.Println(dbConfig)
	db, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to database")
	models.Migrate(db)

	return db, err
}
