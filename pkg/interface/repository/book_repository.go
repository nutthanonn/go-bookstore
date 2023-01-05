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

	inventory := &entities.Inventories{
		Book_id:  book.Book_id,
		Quantity: 1,
	}

	if err := br.db.Model(&inventory).Create(&inventory).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (br *bookRepository) ReadBook() ([]*entities.Books, error) {
	var books []*entities.Books
	if err := br.db.Preload("Inventory_id").Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

func (br *bookRepository) ReadBookByID(ID string) (*entities.Books, error) {
	var book *entities.Books

	if err := br.db.Preload("Inventory_id").First(&book, "book_id = ?", ID).Error; err != nil {
		return nil, err
	}

	return book, nil
}

func (br *bookRepository) UpdateBook(book *entities.Books, ID string) (*entities.Books, error) {
	if err := br.db.Model(&book).Where("employee_id = ?", ID).Updates(book).Error; err != nil {
		return nil, err
	}
	return book, nil
}

func (br *bookRepository) DeleteBook(ID string) error {
	if err := br.db.Where("book_id = ?", ID).Delete(&entities.Books{}).Error; err != nil {
		return err
	}

	return nil
}
