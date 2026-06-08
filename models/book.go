package models

// Book represents a book in the system
type Book struct {
	ID                uint   `gorm:"primaryKey"`
	Title             string `gorm:"not null"`
	Description       string `gorm:"not null"`
	MinAgeRestriction int    `gorm:"not null"`
	CoverURL          string `gorm:"not null"`

	AuthorID uint `gorm:"not null"`
	GenreID  uint `gorm:"not null"`

	Author Author
	Genre  Genre
}
