package main

import (
    "os"

    "github.com/joho/godotenv"
    "github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"

    "book-rental-api/configs"
    "book-rental-api/handlers"
    "book-rental-api/middleware"
    "book-rental-api/models"
	_ "book-rental-api/docs"
)

// @title Book Rental API
// @version 1.0
// @description Book Rental REST API
// @host localhost:8080
// @BasePath /

func main() {

	// Load environment variables
	godotenv.Load()

	// Initialize database connection
	configs.ConnectDB()

	// Automatically create and update database tables
	configs.DB.AutoMigrate(
		&models.User{},
		&models.Book{},
		&models.Author{},
		&models.Genre{},
		&models.Loan{},
	)
	// Create Echo instance
	e := echo.New()

	// User routes
	e.POST("/users/register", handlers.RegisterUser)
	e.POST("/users/login", handlers.LoginUser)

	// Admin routes
	e.GET("/admin/authors", handlers.GetAuthors)
	e.GET("/admin/genres", handlers.GetGenres)
	e.GET("/admin/users", handlers.GetTopUsers)

	// Swagger documentation route
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Protected routes
	auth := e.Group("")
	auth.Use(middleware.JWTMiddleware)

	auth.GET("/users/me", handlers.GetUserProfile)
	auth.GET("/books", handlers.GetBooks)
	auth.POST("/loans", handlers.CreateLoan)

	// Book routes
	e.GET("/books", handlers.GetBooks)

	// Loan routes
	e.POST("/loans", handlers.CreateLoan)

	// Start server 
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	e.Logger.Fatal(e.Start(":" + port))
}