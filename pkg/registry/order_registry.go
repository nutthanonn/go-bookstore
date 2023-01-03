package registry

import (
	ic "github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	ip "github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
	ir "github.com/nutthanonn/go-clean-architecture/pkg/interface/repository"
	ui "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
	up "github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	ur "github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

func (r *registry) NewOrderController() ic.OrderController {
	return ic.NewOrderController(r.NewOrderInteractor())
}

func (r *registry) NewOrderInteractor() ui.OrderInteractor {
	return ui.NewOrderInteractor(r.NewOrderPresenter(), r.NewOrderRepository())
}

func (r *registry) NewOrderRepository() ur.OrderRepository {
	return ir.NewOrderRepository(r.db)
}

func (r *registry) NewOrderPresenter() up.OrderPresenter {
	return ip.NewOrderPresenter()
}
