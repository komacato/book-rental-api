package models

// Author represents a book author
type Author struct {
	ID        uint
	FirstName string
	LastName  string

	Books []Book `json:"-"`
}
