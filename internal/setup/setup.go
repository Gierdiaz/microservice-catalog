package setup

import (
	"github.com/jmoiron/sqlx"

	"github.com/Gierdiaz/Book/internal/handlers"
	"github.com/Gierdiaz/Book/internal/repositories"
	"github.com/Gierdiaz/Book/internal/services"
)

func SetupBook(db *sqlx.DB) *handlers.BookHandler {
	repo := &repositories.BookRepository{DB: db}
	service := &services.BookService{Repo: repo}
	handler := &handlers.BookHandler{Service: service}
	return handler
}
