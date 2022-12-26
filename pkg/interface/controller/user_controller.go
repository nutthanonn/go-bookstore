package controller

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	interactor "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
)

type userController struct {
	userInteractor interactor.UserInteractor
}

type UserController interface {
	CreateUser(user *entities.User) (*entities.User, error)
	ReadUser() (*[]entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(ID string) error
}

func NewUserController(ui interactor.UserInteractor) UserController {
	return &userController{ui}
}

func (uc *userController) CreateUser(user *entities.User) (*entities.User, error) {
	return uc.userInteractor.CreateUser(user)
}

func (uc *userController) ReadUser() (*[]entities.User, error) {
	return uc.userInteractor.ReadUser()
}

func (uc *userController) UpdateUser(user *entities.User) (*entities.User, error) {
	return uc.userInteractor.UpdateUser(user)
}

func (uc *userController) DeleteUser(ID string) error {
	return uc.userInteractor.DeleteUser(ID)
}
