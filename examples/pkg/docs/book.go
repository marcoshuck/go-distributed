package docs

import lorem "github.com/drhodes/golorem"

type Book struct {
	Title string
	Description string
	Content string
}

func NewBook() Book {
	return Book{
		Title:       lorem.Sentence(3, 10),
		Description: lorem.Sentence(50, 150),
		Content:     lorem.Paragraph(150, 500),
	}
}

func GenerateBooks() []Book {
	var books []Book
	for i := 0; i < 1500; i++ {
		books = append(books, NewBook())
	}
	return books
}