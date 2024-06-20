package models

type Book struct {
	ID             string `json:"id"`
	Title          string `json:"title"`
	Author         string `json:"author"`
	Published_date string `json:"published_date"`
}
