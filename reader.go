package main

import (
	"errors"
)

type Reader struct {
	Name          string `json:"name"`
	ReaderID      string `json:"reader_id"`
	BorrowedBooks []Book `json:"borrowed_books"` // список книг, которые читатель взял на время
}

// Создание читателя
func NewReader(library *Library) *Reader {
	reader := &Reader{
		Name:     promptData("Введите имя"),
		ReaderID: promptData("Введите id"),
	}
	AddReader(library, *reader)
	return reader
}

// Позволяет читателю взять книгу, если она доступна
func BorrowBook(library *Library, readerId string, isbn string) (result string, err error) {
	var readerFound bool
	var bookFound bool

	for i, book := range library.Books {
		if book.ISBN == isbn {
			bookFound = true
			if book.Available {
				for j, reader := range library.Readers {
					if readerId == reader.ReaderID {
						readerFound = true
						library.Books[i].Available = false
						library.Readers[j].BorrowedBooks = append(library.Readers[j].BorrowedBooks, book)
						result = "Читатель " + reader.Name + " взял книгу " + book.Title
						return
					}
				}
				if !readerFound {
					err = errors.New("INVALID_READER")
					return
				}
			} else {
				err = errors.New("BOOK_NOT_AVAILABLE")
			}
		}
	}
	if !bookFound {
		err = errors.New("INVALID_BOOK")
	}
	return result, err
}

// Позволяет читателю вернуть книгу
func ReturnBook(library *Library, readerId string, isbn string) (result string, err error) {
	var readerFound bool
	var bookFound bool

	for i, reader := range library.Readers {
		if reader.ReaderID == readerId {
			readerFound = true
			for j, book := range library.Books {
				if book.ISBN == isbn {
					bookFound = true
					library.Readers[i].BorrowedBooks = append(library.Readers[i].BorrowedBooks[:j], library.Readers[j].BorrowedBooks[j+1:]...)
					for k, _ := range library.Books {
						library.Books[k].Available = true
						result = "Читатель " + reader.Name + " возвращает книгу " + book.Title
						return
					}
				}
				if !bookFound {
					err = errors.New("INVALID_BOOK")
					return
				}
			}
		}
	}
	if !readerFound {
		err = errors.New("INVALID_READER")
	}
	return result, err
}
