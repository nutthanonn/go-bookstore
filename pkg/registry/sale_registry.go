package registry

import (
	ic "github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	ip "github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
	ir "github.com/nutthanonn/go-clean-architecture/pkg/interface/repository"
	ui "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
	up "github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	ur "github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

func (r *registry) NewSaleController() ic.SaleController {
	return ic.NewSaleController(r.NewSaleInteractor())
}

func (r *registry) NewSaleInteractor() ui.SalesInteractor {
	return ui.NewSalesInteractor(r.NewSaleRepository(), r.NewSalePresenter())
}

func (r *registry) NewSaleRepository() ur.SaleRepository {
	return ir.NewSaleRepository(r.db)
}

func (r *registry) NewSalePresenter() up.SalePresenter {
	return ip.NewSalePresenter()
}
