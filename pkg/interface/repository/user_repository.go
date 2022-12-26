package repository

import (
	"time"

	"github.com/google/uuid"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user *entities.User) (*entities.User, error) {

	user.Created_at = time.Now()
	user.Updated_at = time.Now()
	user.ID = uuid.New().String()

	if err := ur.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) ReadUser() (*[]entities.User, error) {
	var users []entities.User
	if err := ur.db.Find(&users).Error; err != nil {
		return nil, err
	}

	return &users, nil
}

func (ur *userRepository) UpdateUser(user *entities.User) (*entities.User, error) {

	user.Updated_at = time.Now()

	if err := ur.db.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (ur *userRepository) DeleteUser(ID string) error {
	if err := ur.db.Where("id = ?", ID).Delete(&entities.User{}).Error; err != nil {
		return err
	}

	return nil
}
