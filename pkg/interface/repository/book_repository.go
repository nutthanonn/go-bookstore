package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) repository.BookRepository {
	return &bookRepository{db}
}

func (br *bookRepository) CreateBook(book *entities.Books) (*entities.Books, error) {
	book.Book_id = uuid.New()
	book.Created_at = time.Now()
	book.Updated_at = time.Now()

	err := br.db.Create(book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (br *bookRepository) ReadBook() (*[]entities.Books, error) {
	var books []entities.Books
	err := br.db.Find(&books).Error
	if err != nil {
		return nil, err
	}

	return &books, nil
}

func (br *bookRepository) ReadBookByID(ID string) (*entities.Books, error) {
	var book entities.Books
	err := br.db.First(&book, ID).Error
	if err != nil {
		return nil, err
	}

	return &book, nil
}

func (br *bookRepository) UpdateBook(book *entities.Books, ID string) (*entities.Books, error) {
	var oldBook entities.Books
	book.Updated_at = time.Now()

	if err := br.db.First(&oldBook, ID).Error; err != nil {
		return nil, err
	}

	book.Book_id = oldBook.Book_id
	book.Created_at = oldBook.Created_at
	book.Updated_at = time.Now()

	err := br.db.Save(book).Error
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (br *bookRepository) DeleteBook(ID string) error {
	err := br.db.Delete(&entities.Books{}, ID).Error
	if err != nil {
		return err
	}

	return nil
}
