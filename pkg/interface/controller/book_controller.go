package controller

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	interactor "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
)

type bookController struct {
	bookController interactor.BookInteractor
}

type BookController interface {
	CreateBook(book *entities.Books) (*entities.Books, error)
	ReadBook() ([]*entities.Books, error)
	ReadBookByID(ID string) (*entities.Books, error)
	UpdateBook(book *entities.Books, ID string) (*entities.Books, error)
	DeleteBook(ID string) error
}

func NewBookController(bi interactor.BookInteractor) BookController {
	return &bookController{bi}
}

func (bc *bookController) CreateBook(book *entities.Books) (*entities.Books, error) {
	return bc.bookController.CreateBook(book)
}

func (bc *bookController) ReadBook() ([]*entities.Books, error) {
	return bc.bookController.ReadBook()
}

func (bc *bookController) ReadBookByID(ID string) (*entities.Books, error) {
	return bc.bookController.ReadBookByID(ID)
}

func (bc *bookController) UpdateBook(book *entities.Books, ID string) (*entities.Books, error) {
	return bc.bookController.UpdateBook(book, ID)
}

func (bc *bookController) DeleteBook(ID string) error {
	return bc.bookController.DeleteBook(ID)
}
