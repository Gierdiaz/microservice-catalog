package services

import (
	"github.com/Gierdiaz/Book/internal/models"
	"github.com/Gierdiaz/Book/internal/repositories"
)

type BookService struct {
	Repo *repositories.BookRepository
}

func (s *BookService) GetBooks() ([]models.Book, error) {
	return s.Repo.GetAll()
}

func (s *BookService) GetBookById(id string) (*models.Book, error) {
	return s.Repo.GetById(id)
}

func (s *BookService) CreateBook(book *models.Book) error {
	return s.Repo.Create(book)
}

func (s *BookService) UpdateBook(book *models.Book) error {
	return s.Repo.Update(book)
}

func (s *BookService) DeleteBook(id string) error {
	return s.Repo.Delete(id)
}