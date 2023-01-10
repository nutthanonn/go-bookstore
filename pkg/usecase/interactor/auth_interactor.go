package interfactor

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/models"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

type authInteractor struct {
	authRepository repository.AuthRepository
	authPresenter  presenter.AuthPresenter
}

type AuthInteractor interface {
	SignIn(data *models.SignIn) (bool, error)
	CreateToken(data *entities.Customers) (*models.Token, error)
}

func NewAuthInteractor(ar repository.AuthRepository, ap presenter.AuthPresenter) AuthInteractor {
	return &authInteractor{
		authRepository: ar,
		authPresenter:  ap,
	}
}

func (ai *authInteractor) SignIn(data *models.SignIn) (bool, error) {
	check, err := ai.authRepository.SignIn(data)

	if err != nil {
		return false, nil
	}

	return check, nil
}

func (ai *authInteractor) CreateToken(data *entities.Customers) (*models.Token, error) {
	token, err := ai.authRepository.CreateToken(data)

	if err != nil {
		return nil, err
	}

	return token, nil
}
