package test

import (
	"testing"

	"github.com/Gierdiaz/Book/internal/dto"
	"github.com/Gierdiaz/Book/internal/entities"
	"github.com/google/uuid"

	"github.com/Gierdiaz/Book/internal/services"
)

// Mock do repositório de livros
type MockBookRepository struct {
	Books []entities.Book
}

// Implementação do método de criação no repositório mockado
func (repo *MockBookRepository) CreateBook(book *entities.Book) error {
	repo.Books = append(repo.Books, *book)
	return nil
}

// Implementação do método de obtenção de livro por ID no repositório mockado
func (repo *MockBookRepository) GetBookById(id uuid.UUID) (*entities.Book, error) {
	for _, book := range repo.Books {
		if book.Id == uuid.UUID(id) {
			return &book, nil
		}
	}
	return nil, nil
}

// Implementação do método de atualização no repositório mockado
func (repo *MockBookRepository) UpdateBook(book *entities.Book) error {
	for i, b := range repo.Books {
		if b.Id == book.Id {
			repo.Books[i] = *book
			return nil
		}
	}
	return nil
}

// Implementação do método de exclusão no repositório mockado
func (repo *MockBookRepository) DeleteBook(id uuid.UUID) error {
	for i, book := range repo.Books {
		if book.Id == id {
			repo.Books = append(repo.Books[:i], repo.Books[i+1:]...)
			return nil
		}
	}
	return nil
}

// Teste de criação de livro
func TestCreateBook(t *testing.T) {
	// Configurar repositório mockado e serviço
	repo := &MockBookRepository{}
	service := services.BookService{Repo: repo}

	// Criar livro
	bookDTO := dto.BookDTO{
		Name:      "Livro de Teste",
		Title:     "Título Teste",
		Author:    "Autor Teste",
		Genre:     "Ficção",
		Price:     19.99,
		Quantity:  10,
		Year:      2023,
		Available: true,
	}
	err := service.CreateBook(&bookDTO)
	if err != nil {
		t.Fatalf("Error creating book: %v", err)
	}

	// Validar criação no repositório mockado
	if len(repo.Books) != 1 {
		t.Fatalf("Expected 1 book in repository, Got: %d", len(repo.Books))
	}

	// Validar dados
	book := repo.Books[0]
	if book.Name != bookDTO.Name || book.Title != bookDTO.Title || book.Author != bookDTO.Author {
		t.Fatalf("Book data mismatch. Expected: %+v, Got: %+v", bookDTO, book)
	}
}

// Teste de recuperação de livro
func TestGetBookById(t *testing.T) {
	// Configurar repositório mockado e serviço
	repo := &MockBookRepository{}
	service := services.BookService{Repo: repo}

	// Criar livro
	bookDTO := dto.BookDTO{
		Name:      "Livro de Teste",
		Title:     "Título Teste",
		Author:    "Autor Teste",
		Genre:     "Ficção",
		Price:     19.99,
		Quantity:  10,
		Year:      2023,
		Available: true,
	}
	err := service.CreateBook(&bookDTO)
	if err != nil {
		t.Fatalf("Error creating book: %v", err)
	}

	// Recuperar o ID do livro
	book := repo.Books[0]

	// Recuperar o livro pelo ID
	gotBook, err := service.GetBookById(book.Id)
	if err != nil {
		t.Fatalf("Error retrieving book by ID: %v", err)
	}

	// Validar dados
	if gotBook.Name != bookDTO.Name || gotBook.Title != bookDTO.Title || gotBook.Author != bookDTO.Author {
		t.Fatalf("Retrieved book data mismatch. Expected: %+v, Got: %+v", bookDTO, gotBook)
	}
}

// Teste de atualização de livro
func TestUpdateBook(t *testing.T) {
	// Configurar repositório mockado e serviço
	repo := &MockBookRepository{}
	service := services.BookService{Repo: repo}

	// Criar livro
	bookDTO := dto.BookDTO{
		Name:      "Livro de Teste",
		Title:     "Título Teste",
		Author:    "Autor Teste",
		Genre:     "Ficção",
		Price:     19.99,
		Quantity:  10,
		Year:      2023,
		Available: true,
	}
	err := service.CreateBook(&bookDTO)
	if err != nil {
		t.Fatalf("Error creating book: %v", err)
	}

	// Recuperar o livro
	book := repo.Books[0]

	// Atualizar o livro
	bookDTO.Name = "Livro Atualizado"
	err = service.UpdateBook(&bookDTO)
	if err != nil {
		t.Fatalf("Error updating book: %v", err)
	}

	// Validar atualização no repositório mockado
	updatedBook := repo.Books[0]
	if updatedBook.Name != "Livro Atualizado" {
		t.Fatalf("Book not updated. Expected: 'Livro Atualizado', Got: '%s'", updatedBook.Name)
	}
}

// Teste de exclusão de livro
func TestDeleteBook(t *testing.T) {
	// Configurar repositório mockado e serviço
	repo := &MockBookRepository{}
	service := services.BookService{Repo: repo}

	// Criar livro
	bookDTO := dto.BookDTO{
		Name:      "Livro de Teste",
		Title:     "Título Teste",
		Author:    "Autor Teste",
		Genre:     "Ficção",
		Price:     19.99,
		Quantity:  10,
		Year:      2023,
		Available: true,
	}
	err := service.CreateBook(&bookDTO)
	if err != nil {
		t.Fatalf("Error creating book: %v", err)
	}

	// Recuperar o livro
	book := repo.Books[0]

	// Excluir o livro
	err = service.DeleteBook(book.Id)
	if err != nil {
		t.Fatalf("Error deleting book: %v", err)
	}

	// Validar exclusão no repositório mockado
	if len(repo.Books) != 0 {
		t.Fatalf("Book not deleted. Expected 0 books, Got: %d", len(repo.Books))
	}
}
