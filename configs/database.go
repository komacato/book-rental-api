package configs

import (
	"book-rental-api/models"
	"fmt"
	"os"

	// PostgreSQL driver for GORM
	"gorm.io/driver/postgres"

	// GORM ORM package
	"gorm.io/gorm"
)

// Global database connection variable
var DB *gorm.DB

// ConnectDB creates a connection to the database
func ConnectDB() {

	// Read database connection string from .env file
	dsn := os.Getenv("DATABASE_URL")

	// Open PostgreSQL connection using GORM
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	// Check if connection failed
	if err != nil {

		fmt.Println("Database Error:", err)
		panic("failed to connect database")
	}

	// Store database connection in global variable
	DB = db

	// Automatically create and update database tables
	DB.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Author{},
		&models.Genre{},
		&models.Loan{},
	)

	// Print success message
	fmt.Println("Database connected")
}
