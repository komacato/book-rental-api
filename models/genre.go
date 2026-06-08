package models

// Genre represents a book category
type Genre struct {
	ID   uint
	Name string

	Books []Book `json:"-"`
}
