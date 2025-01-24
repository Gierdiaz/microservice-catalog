package handlers

import (
	"net/http"

	"github.com/Gierdiaz/Book/internal/dto"
	"github.com/Gierdiaz/Book/internal/models"
	"github.com/Gierdiaz/Book/internal/services"
	"github.com/gin-gonic/gin"
)

type BookHandler struct {
	Service *services.BookService
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.CreateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.ToBookDTO(&book))
}

func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.Service.GetBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBookById(c *gin.Context) {
	id := c.Param("id")
	book, err := h.Service.GetBookById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.Service.UpdateBook(&book); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToBookDTO(&book))
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")
	if err := h.Service.DeleteBook(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
