package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
)

type Book struct {
	Book_id      uuid.UUID `json:"book_id"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	Publish_year string    `json:"publish_year"`
	Price        float32   `json:"price"`
	Genre        string    `json:"genre"`
}

type bookPresenter struct{}

func NewBookPresenter() presenter.BookPresenter {
	return &bookPresenter{}
}

func (bp *bookPresenter) BookSuccessResponse(data *entities.Books) *fiber.Map {
	var book = Book{
		Book_id:      data.Book_id,
		Title:        data.Title,
		Author:       data.Author,
		Publish_year: data.Publish_year,
		Price:        data.Price,
		Genre:        data.Genre,
	}

	return &fiber.Map{
		"status": true,
		"error":  nil,
		"data":   book,
	}
}

func (bp *bookPresenter) BooksSuccessResponse(data *[]entities.Books) *fiber.Map {
	var books []Book

	for _, book := range *data {
		b := Book{
			Book_id:      book.Book_id,
			Title:        book.Title,
			Author:       book.Author,
			Publish_year: book.Publish_year,
			Price:        book.Price,
			Genre:        book.Genre,
		}

		books = append(books, b)
	}

	return &fiber.Map{
		"status": true,
		"error":  nil,
		"data":   books,
	}
}

func (bp *bookPresenter) BookErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status":  false,
		"message": err.Error(),
		"data":    nil,
	}
}

func (bp *bookPresenter) BookDeleteSuccessResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   nil,
		"error":  nil,
	}
}
