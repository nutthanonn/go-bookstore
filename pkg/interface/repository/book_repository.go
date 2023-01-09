package repository

import (
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

	if err := br.db.Create(&book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (br *bookRepository) ReadBook() ([]*entities.Books, error) {
	var books []*entities.Books
	if err := br.db.Preload("Inventory").Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (br *bookRepository) ReadBookByID(ID string) (*entities.Books, error) {
	var book *entities.Books

	if err := br.db.Preload("Inventory").First(&book, "book_id = ?", ID).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (br *bookRepository) UpdateBook(book *entities.Books, ID string) (*entities.Books, error) {
	if err := br.db.Model(&book).Where("book_id = ?", ID).Updates(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (br *bookRepository) DeleteBook(ID string) error {
	br.db.Begin() // start transaction

	if err := br.db.Where("book_id = ?", ID).Delete(&entities.OrderDetails{}).Error; err != nil {
		return err
	}

	if err := br.db.Where("book_id = ?", ID).Delete(&entities.Inventories{}).Error; err != nil {
		return err
	}

	if err := br.db.Where("book_id = ?", ID).Delete(&entities.Books{}).Error; err != nil {
		return err
	}

	br.db.Commit() // commit transaction

	return nil
}
