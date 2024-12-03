package main

func main() {
	// Создание экземпляра Library (библиотека)
	Gogolevka := Library{}

	// Создание экземпляров Book (книги)
	bookOne := NewBook("Три мушкетера ч.1", "Александр Дюма", "bk1", false)
	bookSecond := NewBook("Три мушкетера ч.2", "Александр Дюма", "bk2", true)
	bookThrid := NewBook("Три мушкетера ч.3", "Александр Дюма", "bk3", true)
	bookFour := NewBook("Три мушкетера ч.4", "Александр Дюма", "bk4", true)

	// Добавление созданных книг в библиотеку
	AddBooks(&Gogolevka, *bookOne, *bookSecond, *bookThrid, *bookFour)

	// Создание экземпляров Reader (читателя)
	readerOne := newReader("Иван", "re1")
	readerSecond := newReader("Константин", "re2")

	// Добавление читателей в библиотеку
	RegisterReaders(&Gogolevka, *readerOne, *readerSecond)

	GetReaders(Gogolevka.Readers)

	//checkError(func() (string, error) {
	//	return BorrowBook(&Gogolevka, "re1", "bk1")
	//})
}
