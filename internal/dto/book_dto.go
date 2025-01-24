package dto

import "github.com/Gierdiaz/Book/internal/models"

type BookDTO struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Year      int    `json:"year"`
	Available bool   `json:"available"`
}

func ToBookDTO(b *models.Book) BookDTO {
	return BookDTO{
		ID:        b.ID.String(),
		Name:      b.Name,
		Title:     b.Title,
		Author:    b.Author,
		Genre:     b.Genre,
		Price:     b.Price,
		Quantity:  b.Quantity,
		Year:      b.Year,
		Available: b.IsAvailable(),
	}
}

func ToBookDTOs(books []models.Book) []BookDTO {
	dtos := make([]BookDTO, len(books))
	for i, b := range books {
		dtos[i] = ToBookDTO(&b)
	}
	return dtos
}
