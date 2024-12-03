package main

import "fmt"

// Функция отлова ошибки
func checkError(fn func() (string, error)) {
	result, err := fn()
	if err != nil {
		switch err.Error() {
		case "INVALID_READER":
			fmt.Println("Пользователь не существует")
		case "INVALID_BOOK":
			fmt.Println("Книга не существует")
		case "BOOK_NOT_AVAILABLE":
			fmt.Println("Книга недоступна")
		}
	} else {
		fmt.Println(result)
	}
}
