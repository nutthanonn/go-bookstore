package registry

import (
	ic "github.com/nutthanonn/go-clean-architecture/pkg/interface/controller"
	ip "github.com/nutthanonn/go-clean-architecture/pkg/interface/presenter"
	ir "github.com/nutthanonn/go-clean-architecture/pkg/interface/repository"
	ui "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
	up "github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	ur "github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

func (r *registry) NewBookController() ic.BookController {
	return ic.NewBookController(r.NewBookInteractor())
}

func (r *registry) NewBookInteractor() ui.BookInteractor {
	return ui.NewBookInteractor(r.NewBookRepository(), r.NewBookPresenter())
}

func (r *registry) NewBookRepository() ur.BookRepository {
	return ir.NewBookRepository(r.db)
}

func (r *registry) NewBookPresenter() up.BookPresenter {
	return ip.NewBookPresenter()
}
