package services

import (
	"github.com/Gierdiaz/Book/internal/dto"

	"github.com/Gierdiaz/Book/internal/repositories"
	"github.com/Gierdiaz/Book/internal/validator"
	"github.com/google/uuid"
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

func (s *BookService) GetBookById(id uuid.UUID) (dto.BookDTO, error) {
	book, err := s.Repo.GetById(id)
	if err != nil {
		return dto.BookDTO{}, err 
	}
	return dto.ToBookDTO(book), nil 
}


// Para criação de Book, converte DTO para modelo, valida e cria
func (s *BookService) CreateBook(bookDTO *dto.BookDTO) error {
	// Valida o DTO
	if err := validator.ValidateBookDTO(bookDTO); err != nil {
		return err
	}

	book, err := bookDTO.ToModel()
	if err != nil {
		return err
	}

	// Cria o livro no repositório
	return s.Repo.Create(book)
}

func (s *BookService) UpdateBook(bookDTO *dto.BookDTO) error {
	// Valida o DTO
	if err := validator.ValidateBookDTO(bookDTO); err != nil {
		return err
	}

	book, err := bookDTO.ToModel()
	if err != nil {
		return err
	}

	// Atualiza o livro no repositório
	return s.Repo.Update(book)
}

func (s *BookService) DeleteBook(id uuid.UUID) error {
	return s.Repo.Delete(id)
}
