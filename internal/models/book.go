package models

type Book struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Genre     string `json:"genre"`
	Price     int    `json:"price"`
	Quantity  int    `json:"quantity"`
	Year      int    `json:"year"`
	Available bool   `json:"available"`
}