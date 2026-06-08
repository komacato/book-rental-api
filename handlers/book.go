package handlers

import (
	"net/http"
	"strconv"

	"book-rental-api/models"

	"github.com/labstack/echo/v4"
)

// Get all books
// @Summary Get all books
// @Description Get a list of all books from library
// @Tags Books
// @Accept json
// @Produce json
// @Success 200 {array} models.Book
// @Router /books [get]
func GetBooks(c echo.Context) error {
	var books = []models.Book{
		{
			ID:                1,
			Title:             "Harry Potter and the Sorcerer's Stone",
			Description:       "A young wizard's journey begins.",
			MinAgeRestriction: 10,
			CoverURL:          "https://example.com/harry-potter.jpg",
			GenreID:           1,
			AuthorID:          1,

			Author: models.Author{
				ID:        1,
				FirstName: "J.K.",
				LastName:  "Rowling",
			},
			Genre: models.Genre{
				ID:   1,
				Name: "Fantasy",
			},
		},
		{
			ID:                2,
			Title:             "1984",
			Description:       "A dystopian novel about totalitarianism.",
			CoverURL:          "https://example.com/1984.jpg",
			MinAgeRestriction: 16,
			GenreID:           2,
			AuthorID:          2,

			Author: models.Author{
				ID:        2,
				FirstName: "George",
				LastName:  "Orwell",
			},
			Genre: models.Genre{
				ID:   2,
				Name: "Dystopian",
			},
		},
		{
			ID:                3,
			Title:             "To Kill a Mockingbird",
			Description:       "A novel about racial injustice in the Deep South.",
			CoverURL:          "https://example.com/to-kill-a-mockingbird.jpg",
			MinAgeRestriction: 12,
			GenreID:           3,
			AuthorID:          3,

			Author: models.Author{
				ID:        3,
				FirstName: "Harper",
				LastName:  "Lee",
			},
			Genre: models.Genre{
				ID:   3,
				Name: "Classic",
			},
		},
	}

	// Get genre query parameter
	genre := c.QueryParam("genre")

	// Return all books if genre is not provided
	if genre == "" {
		return c.JSON(http.StatusOK, books)
	}

	// Convert genre to integer
	genreID, _ := strconv.Atoi(genre)

	// Filter books by genre
	var filteredBooks []models.Book
	for _, book := range books {
		if book.GenreID == uint(genreID) {
			filteredBooks = append(filteredBooks, book)
		}
	}

	// Return filtered books
	return c.JSON(http.StatusOK, filteredBooks)
}
