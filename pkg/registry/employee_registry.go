package registry

import (
	ic "github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	ip "github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
	ir "github.com/nutthanonn/go-clean-architecture/pkg/interface/repository"
	ui "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
	up "github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	ur "github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

func (r *registry) NewEmployeeController() ic.EmployeeController {
	return ic.NewEmployeeController(r.NewEmployeeInteractor())
}

func (r *registry) NewEmployeeInteractor() ui.EmployeeInteractor {
	return ui.NewEmployeeInteractor(r.NewEmployeePresenter(), r.NewEmployeeRepository())
}

func (r *registry) NewEmployeeRepository() ur.EmployeeRepository {
	return ir.NewEmployeeRepository(r.db)
}

func (r *registry) NewEmployeePresenter() up.EmployeePresenter {
	return ip.NewEmployeePresenter()
}
