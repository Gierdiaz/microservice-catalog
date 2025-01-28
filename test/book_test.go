package test

import (
	"testing"
	"time"

	"github.com/Gierdiaz/Book/internal/dto"
	"github.com/Gierdiaz/Book/internal/entities"
	"github.com/Gierdiaz/Book/internal/services"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockBookRepository é uma implementação mockada de BookRepositoryInterface
type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) Create(book *entities.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) GetAll() ([]entities.Book, error) {
	args := m.Called()
	return args.Get(0).([]entities.Book), args.Error(1)
}

func (m *MockBookRepository) GetById(id uuid.UUID) (*entities.Book, error) {
	args := m.Called(id)
	return args.Get(0).(*entities.Book), args.Error(1)
}

func (m *MockBookRepository) Update(book *entities.Book) error {
	args := m.Called(book)
	return args.Error(0)
}

func (m *MockBookRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateBook(t *testing.T) {
	mockRepo := new(MockBookRepository)
	service := services.NewBookService(mockRepo)

	book := &entities.Book{
		Id:        uuid.New(),
		Name:      "Test Book",
		Title:     "Learning Go",
		Author:    "John Doe",
		Genre:     "Programming",
		Price:     29.99,
		Quantity:  10,
		Year:      2024,
		Available: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("Create", book).Return(nil)

	err := service.CreateBook(&dto.BookDTO{
		Name:      "Test Book",
		Title:     "Learning Go",
		Author:    "John Doe",
		Genre:     "Programming",
		Price:     29.99,
		Quantity:  10,
		Year:      2024,
		Available: true,
	})

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetBooks(t *testing.T) {
	mockRepo := new(MockBookRepository)
	service := services.NewBookService(mockRepo)

	books := []entities.Book{
		{
			Id:        uuid.New(),
			Name:      "Test Book 1",
			Title:     "Go Basics",
			Author:    "Jane Doe",
			Genre:     "Programming",
			Price:     19.99,
			Quantity:  5,
			Year:      2022,
			Available: true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			Id:        uuid.New(),
			Name:      "Test Book 2",
			Title:     "Advanced Go",
			Author:    "John Doe",
			Genre:     "Programming",
			Price:     39.99,
			Quantity:  3,
			Year:      2023,
			Available: true,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	mockRepo.On("GetAll").Return(books, nil)

	result, err := service.GetBooks()

	assert.NoError(t, err)
	assert.Equal(t, len(books), len(result))
	mockRepo.AssertExpectations(t)
}

func TestGetBookById(t *testing.T) {
	mockRepo := new(MockBookRepository)
	service := services.NewBookService(mockRepo)

	book := &entities.Book{
		Id:        uuid.New(),
		Name:      "Test Book",
		Title:     "Learning Go",
		Author:    "John Doe",
		Genre:     "Programming",
		Price:     29.99,
		Quantity:  10,
		Year:      2024,
		Available: true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockRepo.On("GetById", book.Id).Return(book, nil)

	result, err := service.GetBookById(book.Id)

	assert.NoError(t, err)
	assert.Equal(t, book, result)
	mockRepo.AssertExpectations(t)
}

func TestDeleteBook(t *testing.T) {
	mockRepo := new(MockBookRepository)
	service := services.NewBookService(mockRepo)

	bookID := uuid.New()
	mockRepo.On("Delete", bookID).Return(nil)

	err := service.DeleteBook(bookID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}