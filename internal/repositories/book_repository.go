package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/Gierdiaz/Book/internal/models"
)

type BookRepository struct {
	DB *sqlx.DB
}

func (r *BookRepository) Create(book *models.Book) error {
	query := `
		INSERT INTO books (id, name, title, author, genre, price, quantity, year, available, created_at)
		VALUES (:id, :name, :title, :author, :genre, :price, :quantity, :year, :available, :created_at)
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

func (r *BookRepository) GetById(id string) (*models.Book, error) {
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
		SET name = :name, title = :title, author = :author, genre = :genre, price = :price, quantity = :quantity, year = :year, available = :available
		WHERE id = :id
	`
	_, err := r.DB.NamedExec(query, book)
	return err
}

func (r *BookRepository) Delete(id string) error {
	query := `DELETE FROM books WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}
