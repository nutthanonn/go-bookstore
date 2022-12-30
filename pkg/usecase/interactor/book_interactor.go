package interfactor

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

type bookInteractor struct {
	BookRepository repository.BookRepository
	BookPresenter  presenter.BookPresenter
}

type BookInteractor interface {
	CreateBook(book *entities.Books) (*entities.Books, error)
	ReadBook() ([]*entities.Books, error)
	ReadBookByID(ID string) (*entities.Books, error)
	UpdateBook(book *entities.Books, ID string) (*entities.Books, error)
	DeleteBook(ID string) error
}

func NewBookInteractor(br repository.BookRepository, bp presenter.BookPresenter) BookInteractor {
	return &bookInteractor{
		BookRepository: br,
		BookPresenter:  bp,
	}
}

func (bi *bookInteractor) CreateBook(book *entities.Books) (*entities.Books, error) {
	createdBook, err := bi.BookRepository.CreateBook(book)
	if err != nil {
		return nil, err
	}

	return createdBook, nil
}

func (bi *bookInteractor) ReadBook() ([]*entities.Books, error) {
	books, err := bi.BookRepository.ReadBook()
	if err != nil {
		return nil, err
	}

	return books, nil
}

func (bi *bookInteractor) ReadBookByID(ID string) (*entities.Books, error) {
	book, err := bi.BookRepository.ReadBookByID(ID)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (bi *bookInteractor) UpdateBook(book *entities.Books, ID string) (*entities.Books, error) {
	updatedBook, err := bi.BookRepository.UpdateBook(book, ID)
	if err != nil {
		return nil, err
	}

	return updatedBook, nil
}

func (bi *bookInteractor) DeleteBook(ID string) error {
	err := bi.BookRepository.DeleteBook(ID)
	if err != nil {
		return err
	}

	return nil
}
