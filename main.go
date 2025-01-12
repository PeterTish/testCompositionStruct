package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Создание экземпляра Library (библиотека)
	Gogolevka := &Library{}

	for {
		userAction := getActionSelection()

		switch userAction {
		case 1:
			NewBook(Gogolevka)
		case 2:
			NewReader(Gogolevka)
		}
	}

	content, _ := ReadFile("books.json")
	fmt.Println(content)

	//checkError(func() (string, error) {
	//	return BorrowBook(&Gogolevka, "re1", "bk1")
	//})
	//checkError(func() (string, error) {
	//	return BorrowBook(&Gogolevka, "re2", "bk1")
	//})
}

// Ввод значений пользователя для создания объектов
func promptData(prompt string) string {
	fmt.Printf("%s: ", prompt)

	userInput := bufio.NewScanner(os.Stdin)
	userInput.Scan()
	return userInput.Text()
}

// Выбор пользователем действия
func getActionSelection() int {
	var userAction int
	fmt.Println("1. Добавить книгу")
	fmt.Println("2. Добавить читателя")
	fmt.Print("Выберите действие: ")
	fmt.Scan(&userAction)

	return userAction
}
