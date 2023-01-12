package repository

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
	"gorm.io/gorm"
)

type saleRepository struct {
	db *gorm.DB
}

func NewSaleRepository(db *gorm.DB) repository.SaleRepository {
	return &saleRepository{db}
}

func (sr *saleRepository) CreateSale(sale *entities.Sales) (*entities.Sales, error) {
	if err := sr.db.Create(sale).Error; err != nil {
		return nil, err
	}

	return sale, nil
}
