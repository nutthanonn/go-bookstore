package interfactor

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	CreateUser(user *entities.User) (*entities.User, error)
	ReadUser() (*[]entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(ID string) error
}

func NewUserInteractor(ur repository.UserRepository, up presenter.UserPresenter) UserInteractor {
	return &userInteractor{
		UserRepository: ur,
		UserPresenter:  up,
	}
}

func (ui *userInteractor) CreateUser(user *entities.User) (*entities.User, error) {
	createdUser, err := ui.UserRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return createdUser, nil
}

func (ui *userInteractor) ReadUser() (*[]entities.User, error) {
	users, err := ui.UserRepository.ReadUser()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (ui *userInteractor) UpdateUser(user *entities.User) (*entities.User, error) {
	updatedUser, err := ui.UserRepository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (ui *userInteractor) DeleteUser(ID string) error {
	err := ui.UserRepository.DeleteUser(ID)
	if err != nil {
		return err
	}

	return nil
}
