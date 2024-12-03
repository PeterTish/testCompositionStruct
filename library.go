package main

import "fmt"

type Library struct {
	Books   []Book   // список всех книг в библиотеке
	Readers []Reader // список всех читателей
}

// Добавляет книги в библиотеку
func AddBooks(library *Library, books ...Book) {
	for _, book := range books {
		library.Books = append(library.Books, book)
		fmt.Printf("Книга '%s' добавлена в общий каталог\n", book.Title)
	}
}

// Регистрация нового читателя
func registerReaders(library *Library, readers ...Reader) {
	for _, reader := range readers {
		library.Readers = append(library.Readers, reader)
		fmt.Printf("Читатель %s добавлен в библиотеку\n", reader.Name)
	}
}
