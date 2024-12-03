package main

import "fmt"

func checkError(fn func() error) {
	if err := fn(); err != nil {
		switch error.Error() {
		case "INVALID_READER":
			fmt.Println("Пользователь не существует")
		case "INVALID_BOOK":
			fmt.Println("Книга не существует")
		}
	} else {
		fmt.Println(fn)
	}
}
