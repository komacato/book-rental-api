package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

// CreateLoan creates a new loan record
// @Summary Create a new loan
// @Description membuat/menyimpan data peminjaman buku
// @Tags Books
// @Accept json
// @Produce json
// @Success 201 {object} models.Loan
// @Router /loans [post]
func CreateLoan(c echo.Context) error {

	// Set loan date 
	loanDate := time.Now()

	// Calculate due date
	dueDate := loanDate.AddDate(0, 0, 7)
	
	// Return response
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message":   "loan created",
		"loan_date": loanDate,
		"due_date":  dueDate,
	})
}
