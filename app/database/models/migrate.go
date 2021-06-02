package models

import (
	"fmt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Sample{})
	fmt.Println("Auto Migration has beed processed")
}
