package interfactor

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

type salesInteractor struct {
	salesRepository repository.SaleRepository
	salesPresenter  presenter.SalePresenter
}

type SalesInteractor interface {
	CreateSales(data *entities.Sales) (*entities.Sales, error)
}

func NewSalesInteractor(sr repository.SaleRepository, sp presenter.SalePresenter) SalesInteractor {
	return &salesInteractor{
		salesRepository: sr,
		salesPresenter:  sp,
	}
}

func (si *salesInteractor) CreateSales(data *entities.Sales) (*entities.Sales, error) {
	createdSale, err := si.salesRepository.CreateSale(data)
	if err != nil {
		return nil, err
	}

	return createdSale, nil
}
