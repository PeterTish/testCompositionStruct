package main

import "fmt"

func main() {
	// Создание экземпляра Library (библиотека)
	Gogolevka := &Library{}
	//Pushkinskaya := &Library{}

	for {
		userLibrary := getLibrarySelection()
		userAction := getActionSelection()

		if userLibrary == 1 {
			switch userAction {
			case 1:
				NewBook(Gogolevka)
			case 2:
				NewReader(Gogolevka)
			}
		}
	}

	//checkError(func() (string, error) {
	//	return BorrowBook(&Gogolevka, "re1", "bk1")
	//})
	//checkError(func() (string, error) {
	//	return BorrowBook(&Gogolevka, "re2", "bk1")
	//})
}

func promptData(prompt string) string {
	fmt.Printf("%s: ", prompt)
	var userInput string
	fmt.Scan(&userInput)
	return userInput
}

func getLibrarySelection() int {
	var userLibrary int

	fmt.Println("1. Gogolevka")
	fmt.Println("2. Pushkinskaya")
	fmt.Print("Выберите библиотеку: ")
	fmt.Scan(&userLibrary)

	return userLibrary
}

func getActionSelection() int {
	var userAction int
	fmt.Println("1. Добавить книгу")
	fmt.Println("2. Добавить читателя")
	fmt.Println("Выберите действие: ")
	fmt.Scan(&userAction)

	return userAction
}
