package main

func main() {
	Gogolevka := Library{}

	bookOne := newBook("Три мушкетера ч.1", "Александр Дюма", "bk1", false)
	bookSecond := newBook("Три мушкетера ч.2", "Александр Дюма", "bk2", true)
	bookThrid := newBook("Три мушкетера ч.3", "Александр Дюма", "bk3", true)
	bookFour := newBook("Три мушкетера ч.4", "Александр Дюма", "bk4", true)

	AddBooks(&Gogolevka, *bookOne, *bookSecond, *bookThrid, *bookFour)

	readerOne := newReader("Иван", "re1")
	readerSecond := newReader("Константин", "re2")

	registerReaders(&Gogolevka, *readerOne, *readerSecond)

	checkError(func() (string, error) {
		return BorrowBook(&Gogolevka, "re1", "bk1")
	})
}
