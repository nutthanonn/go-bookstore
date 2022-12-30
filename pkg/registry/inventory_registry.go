package registry

import (
	ic "github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	ip "github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
	ir "github.com/nutthanonn/go-clean-architecture/pkg/interface/repository"
	ui "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
	up "github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	ur "github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

func (r *registry) NewInventoryController() ic.InventoryController {
	return ic.NewInventoryController(r.NewInventoryInteractor())
}

func (r *registry) NewInventoryInteractor() ui.InventoryInteractor {
	return ui.NewInventoryInteractor(r.NewInventoryPresenter(), r.NewInventoryRepository())
}

func (r *registry) NewInventoryRepository() ur.InventoryRepository {
	return ir.NewInventoryRepository(r.db)
}

func (r *registry) NewInventoryPresenter() up.InventoryPresenter {
	return ip.NewInventoryPresenter()
}
