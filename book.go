package main

type Book struct {
	Title     string
	Author    string
	ISBN      string // уникальный идентификатор
	Available bool   // доступна ли книга для выдачи
}

// Создание новой книги
func newBook(title string, author string, isbn string, available bool) *Book {
	return &Book{
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		Available: true,
	}
}
