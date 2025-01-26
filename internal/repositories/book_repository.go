package repositories

import (
	"time"

	"github.com/Gierdiaz/Book/internal/models"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type BookRepository struct {
	DB *sqlx.DB
}

func (r *BookRepository) Create(book *models.Book) error {
	query := `
		INSERT INTO books (id, name, title, author, genre, price, quantity, year, available, created_at, updated_at)
		VALUES (:id, :name, :title, :author, :genre, :price, :quantity, :year, :available, :created_at, :updated_at)
	`
	_, err := r.DB.NamedExec(query, book)
	return err
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
	query := `
		SELECT id, name, title, author, genre, price, quantity, year, available, created_at
		FROM books
	`
	var books []models.Book
	err := r.DB.Select(&books, query)
	return books, err
}

func (r *BookRepository) GetById(id uuid.UUID) (*models.Book, error) {
	query := `
		SELECT id, name, title, author, genre, price, quantity, year, available, created_at
		FROM books WHERE id = $1
	`
	var book models.Book
	err := r.DB.Get(&book, query, id)
	return &book, err
}

func (r *BookRepository) Update(book *models.Book) error {
	query := `
		UPDATE books
		SET name = :name, title = :title, author = :author, genre = :genre, price = :price,
			quantity = :quantity, year = :year, available = :available, updated_at = :updated_at
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
