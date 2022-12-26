package repository

import (
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
)

type UserRepository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	ReadUser() (*[]entities.User, error)
	UpdateUser(user *entities.User) (*entities.User, error)
	DeleteUser(ID string) error
}
