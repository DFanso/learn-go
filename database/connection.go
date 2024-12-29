package database

import (
	"fmt"
	"log"

	"github.com/dfanso/learn-go/models"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	var err error

	DB, err = gorm.Open(sqlite.Open("application.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate both Book and User schemas
	err = DB.AutoMigrate(&models.Book{}, &models.User{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	fmt.Println("Database connection successfully opened and migrated")
}
