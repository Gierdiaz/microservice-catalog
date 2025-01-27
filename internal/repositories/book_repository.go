package repositories

import (
	"fmt"
	"time"

	"github.com/Gierdiaz/Book/internal/contracts"
	"github.com/Gierdiaz/Book/internal/entities"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BookRepository struct {
	DB *sqlx.DB
}

func NewBookRepository(db *sqlx.DB) contracts.BookRepositoryInterface {
	return &BookRepository{DB: db}
}

func (r *BookRepository) Create(book *entities.Book) error {
	query := `
		INSERT INTO books (id, name, title, author, genre, price, quantity, year, available, created_at, updated_at)
		VALUES (:id, :name, :title, :author, :genre, :price, :quantity, :year, :available, NOW(), NOW())
	`
	_, err := r.DB.NamedExec(query, book)
	return err
}

func (r *BookRepository) GetAll() ([]entities.Book, error) {
	query := `
		SELECT id, name, title, author, genre, price, quantity, year, available, created_at
		FROM books
	`
	var books []entities.Book
	err := r.DB.Select(&books, query)
	return books, err
}

func (r *BookRepository) GetById(id uuid.UUID) (*entities.Book, error) {
	query := `
		SELECT id, name, title, author, genre, price, quantity, year, available, created_at
		FROM books WHERE id = $1
	`
	var book entities.Book
	err := r.DB.Get(&book, query, id)

	if err != nil {
		return nil, fmt.Errorf("error fetching book with id %s: %v", id, err)
	}

	if book.Id == uuid.Nil {
		return nil, fmt.Errorf("book with id %s not found", id)
	}

	return &book, err
}

func (r *BookRepository) Update(book *entities.Book) error {
	query := `
		UPDATE books
		SET name = :name, title = :title, author = :author, genre = :genre, price = :price,
			quantity = :quantity, year = :year, available = :available, updated_at = NOW()
		WHERE id = :id
	`
	_, err := r.DB.NamedExec(query, book)
	return err
}

func (r *BookRepository) Delete(id uuid.UUID) error {
	query := `
		UPDATE books
		SET deleted_at = :deleted_at
		WHERE id = :id AND deleted_at IS NULL
	`
	_, err := r.DB.NamedExec(query, map[string]interface{}{"id": id, "deleted_at": time.Now()})
	return err
}
