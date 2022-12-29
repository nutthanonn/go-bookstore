package registry

import (
	ic "github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	ip "github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
	ir "github.com/nutthanonn/go-clean-architecture/pkg/interface/repository"
	ui "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
	up "github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	ur "github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

func (r *registry) NewCustomerController() ic.CustomerController {
	return ic.NewCustomerController(r.NewCustomerInteractor())
}

func (r *registry) NewCustomerInteractor() ui.CustomerInteractor {
	return ui.NewCustomerInteractor(r.NewCustomerPresenter(), r.NewCustomerRepository())
}

func (r *registry) NewCustomerRepository() ur.CustomerRepository {
	return ir.NewCustomerRepository(r.db)
}

func (r *registry) NewCustomerPresenter() up.CustomerPresenter {
	return ip.NewCustomerPresenter()
}
