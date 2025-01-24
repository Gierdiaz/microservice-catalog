package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Genre     string    `json:"genre"`
	Price     int       `json:"price"`
	Quantity  int       `json:"quantity"`
	Year      int       `json:"year"`
	Available bool      `json:"available"`
	CreatedAt time.Time `json:"created_at"`
}

func NewBook(name, title, author, genre string, price, quantity, year int) (*Book, error) {
	if name == "" || title == "" || author == "" {
		return nil, errors.New("name, title, and author are required")
	}

	if price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}

	if quantity < 0 {
		return nil, errors.New("quantity cannot be negative")
	}

	if year < 0 {
		return nil, errors.New("year cannot be negative")
	}

	return &Book{
		ID:        uuid.New(),
		Name:      name,
		Title:     title,
		Author:    author,
		Genre:     genre,
		Price:     price,
		Quantity:  quantity,
		Year:      year,
		Available: quantity > 0,
		CreatedAt: time.Now(),
	}, nil
}

func (b *Book) IsAvailable() bool {
	return b.Quantity > 0
}

func (b *Book) ApplyDiscount(percent int) error {
	if percent < 0 || percent > 100 {
		return errors.New("percent must be between 0 and 100")
	}

	discount := (b.Price * percent) / 100
	b.Price -= discount

	return nil
}
