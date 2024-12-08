package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID        uuid.UUID   `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Year      int    `json:"year"`
	Available bool   `json:"available"`
	CreatedAt time.Time `json:"created_at"`
}