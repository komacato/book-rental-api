package models

import "time"

// Loan represents a book borrowing record
type Loan struct {
	ID uint `gorm:"primaryKey"`

	UserID uint `gorm:"not null"`
	BookID uint `gorm:"not null"`

	LoanDate time.Time `gorm:"not null"`
	DueDate  time.Time `gorm:"not null"`

	User User
	Book Book
}
