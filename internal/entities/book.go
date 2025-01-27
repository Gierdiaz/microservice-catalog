package entities

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	Id        uuid.UUID    `db:"id"`
	Name      string       `db:"name"`
	Title     string       `db:"title"`
	Author    string       `db:"author"`
	Genre     string       `db:"genre"`
	Price     int          `db:"price"`
	Quantity  int          `db:"quantity"`
	Year      int          `db:"year"`
	Available bool         `db:"available"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
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
		Id:        uuid.New(),
		Name:      name,
		Title:     title,
		Author:    author,
		Genre:     genre,
		Price:     price,
		Quantity:  quantity,
		Year:      year,
		Available: quantity > 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

func (b *Book) IsDeleted() bool {
	return b.DeletedAt.Valid
}
