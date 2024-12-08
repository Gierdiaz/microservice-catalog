package repositories

import (
	"database/sql"

	"github.com/Gierdiaz/Book/internal/models"
)


type BookRepository struct {
	DB *sql.DB
}

func (r *BookRepository) Create(book *models.Book) error {
	query := `INSERT INTO books (name, title, author, genre, price, quantity, year, available) VALUES (:id, :name, :title, :author, :genre, :price, :quantity, :year, :available)`
	_, err := r.DB.Exec(query, book.Name, book.Title, book.Author, book.Genre, book.Price, book.Quantity, book.Year, book.Available)
	return err
}

func (r *BookRepository) GetAll() ([]models.Book, error) {
	query := `SELECT id, name, title, author, genre, price, quantity, year, available FROM books`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var book models.Book
		err := rows.Scan(&book.ID, &book.Name, &book.Title, &book.Author, &book.Genre, &book.Price, &book.Quantity, &book.Year, &book.Available)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	return books, nil
}

func (r *BookRepository) GetById(id string) (*models.Book, error) {
	query := `SELECT id, name, title, author, genre, price, quantity, year, available FROM books WHERE id = :id`
	row := r.DB.QueryRow(query, id)
	var book models.Book
	err := row.Scan(&book.ID, &book.Name, &book.Title, &book.Author, &book.Genre, &book.Price, &book.Quantity, &book.Year, &book.Available)
	if err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) Update(book *models.Book) error {
	query := `UPDATE books SET name = :name, title = :title, author = :author, genre = :genre, price = :price, quantity = :quantity, year = :year, available = :available WHERE id = :id`
	_, err := r.DB.Exec(query, book.Name, book.Title, book.Author, book.Genre, book.Price, book.Quantity, book.Year, book.Available, book.ID)
	return err
}

func (r *BookRepository) Delete(id string) error {
	query := `DELETE FROM books WHERE id = :id`
	_, err := r.DB.Exec(query, id)
	return err
}