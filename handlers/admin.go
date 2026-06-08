package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// GetAuthors returns a list of authors
func GetAuthors(c echo.Context) error {
	return c.JSON(http.StatusOK, []map[string]interface{}{
		{
			"author": "J.K. Rowling",
			"book_count": 10,
		},
		{
			"author": "George Orwell",
			"book_count": 5,
		},
	})
}

// GetGenres returns a list of genres
func GetGenres(c echo.Context) error {
	return c.JSON(http.StatusOK, []map[string]interface{}{
		{
			"genre": "Fantasy",
			"loan_count": 10,
		},
		{
			"genre": "Science Fiction",
			"loan_count": 5,
		},
	})
}

// GetTopUsers returns top borrowers
func GetTopUsers(c echo.Context) error {
	return c.JSON(http.StatusOK, []map[string]interface{}{
		{
			"user": "Raden Komara",
			"loan_count": 10,
		},
		{
			"user": "Carmen Nyoman",
			"loan_count": 5,
		},
	})
}
