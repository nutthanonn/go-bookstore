package repository

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/models"
)

type AuthRepository interface {
	SignIn(data *models.SignIn) (bool, error)
	CreateToken(data *entities.Customers) (*models.Token, error)
}
