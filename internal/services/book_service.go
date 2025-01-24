package services

import (
	"github.com/Gierdiaz/Book/internal/dto"
	"github.com/Gierdiaz/Book/internal/models"
	"github.com/Gierdiaz/Book/internal/repositories"
)

type BookService struct {
	Repo *repositories.BookRepository
}

func (s *BookService) GetBooks() ([]dto.BookDTO, error) {
	books, err := s.Repo.GetAll()
	if err != nil {
		return nil, err
	}
	return dto.ToBookDTOs(books), nil
}

func (s *BookService) GetBookById(id string) (*dto.BookDTO, error) {
	book, err := s.Repo.GetById(id)
	if err != nil {
		return nil, err
	}
	bookDTO := dto.ToBookDTO(book)
	return &bookDTO, nil
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
