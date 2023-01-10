package controller

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/models"
	interactor "github.com/nutthanonn/go-clean-architecture/pkg/usecase/interactor"
)

type authController struct {
	AuthController interactor.AuthInteractor
}

type AuthController interface {
	SignIn(data *models.SignIn) (bool, error)
	CreateToken(data *entities.Customers) (*models.Token, error)
}

func NewAuthController(ai interactor.AuthInteractor) AuthController {
	return &authController{
		AuthController: ai,
	}
}

func (ac *authController) SignIn(data *models.SignIn) (bool, error) {
	return ac.AuthController.SignIn(data)
}

func (ac *authController) CreateToken(data *entities.Customers) (*models.Token, error) {
	return ac.AuthController.CreateToken(data)
}
