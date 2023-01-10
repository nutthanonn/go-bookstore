package registry

import (
	ic "github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	ip "github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
	ir "github.com/nutthanonn/go-clean-architecture/pkg/interface/repository"
	ui "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
	up "github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	ur "github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

func (r *registry) NewAuthController() ic.AuthController {
	return ic.NewAuthController(r.NewAuthInteractor())
}

func (r *registry) NewAuthInteractor() ic.AuthController {
	return ui.NewAuthInteractor(r.NewAuthRepository(), r.NewAuthPresenter())
}

func (r *registry) NewAuthRepository() ur.AuthRepository {
	return ir.NewAuthRepository(r.db)
}

func (r *registry) NewAuthPresenter() up.AuthPresenter {
	return ip.NewAuthPresenter()
}
