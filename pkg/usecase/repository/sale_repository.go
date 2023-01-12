package repository

import "github.com/nutthanonn/go-clean-architecture/pkg/entities"

type SaleRepository interface {
	CreateSale(data *entities.Sales) (*entities.Sales, error)
}
