package main

type Book struct {
	Title     string `json:"title"`
	Author    string `json:"author"`
	ISBN      string `json:"ISBN"`      // уникальный идентификатор
	Available bool   `json:"available"` // доступна ли книга для выдачи
}

// Создание новой книги
func NewBook(library *Library) *Book {
	book := &Book{
		Title:     promptData("Введите наименование книги"),
		Author:    promptData("Введите автора книги"),
		ISBN:      promptData("Введите уникальный идентификатор"),
		Available: true,
	}
	AddBook(library, *book)
	return book
}
