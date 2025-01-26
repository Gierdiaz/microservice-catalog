package dto

import (
	"time"

	"github.com/Gierdiaz/Book/internal/models"
	"github.com/google/uuid"
)

type BookDTO struct {
	ID        uuid.UUID  `json:"id,omitempty"`
	Name      string     `json:"name" binding:"required"`
	Title     string     `json:"title" binding:"required"`
	Author    string     `json:"author" binding:"required"`
	Genre     string     `json:"genre" binding:"required"`
	Price     int        `json:"price" binding:"required,gt=0"`
	Quantity  int        `json:"quantity" binding:"required,gt=0"`
	Year      int        `json:"year" binding:"required,gt=0"`
	Available bool       `json:"available"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}

// ToBookDTO converte um modelo Book em um BookDTO
func ToBookDTO(book *models.Book) BookDTO {
	return BookDTO{
		ID:        book.ID,
		Name:      book.Name,
		Title:     book.Title,
		Author:    book.Author,
		Genre:     book.Genre,
		Price:     book.Price,
		Quantity:  book.Quantity,
		Year:      book.Year,
		Available: book.Available,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
}

// ToBookDTOs converte uma lista de modelos Book para uma lista de BookDTO
func ToBookDTOs(books []models.Book) []BookDTO {
	var bookDTOs []BookDTO
	for _, book := range books {
		bookDTOs = append(bookDTOs, ToBookDTO(&book))
	}
	return bookDTOs
}

func (b *BookDTO) ToModel() (*models.Book, error) {
	// Garantindo que o ID seja um UUID válido
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &models.Book{
		ID:        id,
		Name:      b.Name,
		Title:     b.Title,
		Author:    b.Author,
		Genre:     b.Genre,
		Price:     b.Price,
		Quantity:  b.Quantity,
		Year:      b.Year,
		Available: b.Available,
	}, nil
}