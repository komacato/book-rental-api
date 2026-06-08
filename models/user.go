package models

// User represents a customer
type User struct {
	ID uint `gorm:"primaryKey"`

	FirstName string `json:"first_name" gorm:"not null"`
	LastName  string `json:"last_name" gorm:"not null"`

	Email string `json:"email" gorm:"unique;not null"`

	Password string `json:"password" gorm:"not null"`

	Address string `json:"address" gorm:"not null"`

	DateOfBirth string `json:"date_of_birth" gorm:"not null"`

	Loans []Loan
}
