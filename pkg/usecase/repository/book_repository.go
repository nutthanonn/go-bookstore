package repository

import "github.com/nutthanonn/go-clean-architecture/pkg/entities"

type BookRepository interface {
	CreateBook(book *entities.Books) (*entities.Books, error)
	ReadBook() (*[]entities.Books, error)
	ReadBookByID(ID string) (*entities.Books, error)
	UpdateBook(book *entities.Books) (*entities.Books, error)
	DeleteBook(ID string) error
}
