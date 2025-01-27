package contracts

import (
	"github.com/Gierdiaz/Book/internal/entities"
	"github.com/google/uuid"
)

type BookRepositoryInterface interface {
	Create(book *entities.Book) error
	GetAll() ([]entities.Book, error)
	GetById(id uuid.UUID) (*entities.Book, error)
	Update(book *entities.Book) error
	Delete(id uuid.UUID) error
}
