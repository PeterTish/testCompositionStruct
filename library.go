package main

import (
	"fmt"
)

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

// Получение всех книг
func GetBooks(books []Book) {
	for _, book := range books {
		var availableBook string

		if book.Available {
			availableBook = "Доступна"
		} else {
			availableBook = "Недоступна"
		}
		fmt.Printf("\nНаименование: %s\nАвтор: %s\nДоступность книги: %s\n-----", book.Title, book.Author, availableBook)
	}
}

// Регистрация нового читателя
func RegisterReaders(library *Library, readers ...Reader) {
	for _, reader := range readers {
		library.Readers = append(library.Readers, reader)
		fmt.Printf("Читатель %s добавлен в библиотеку\n", reader.Name)
	}
}

// Функция получения всех пользователей библиотеки
func GetReaders(readers []Reader) {
	for _, reader := range readers {
		fmt.Println(reader.Name)
	}
}
