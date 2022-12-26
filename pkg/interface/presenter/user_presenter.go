package presenter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/presenter"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type userPresenter struct{}

func NewUserPresenter() presenter.UserPresenter {
	return &userPresenter{}
}

func (up *userPresenter) UserSuccessResponse(data *entities.User) *fiber.Map {
	user := User{
		ID:        data.ID,
		FirstName: data.FirstName,
		LastName:  data.LastName,
	}

	return &fiber.Map{
		"status": true,
		"data":   user,
		"error":  nil,
	}
}

func (up *userPresenter) UsersSuccessResponse(data *[]entities.User) *fiber.Map {
	var users []User

	for _, user := range *data {
		u := User{
			ID:        user.ID,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}
		users = append(users, u)
	}

	return &fiber.Map{
		"status": true,
		"data":   data,
		"error":  nil,
	}
}

func (up *userPresenter) UserErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status": false,
		"data":   nil,
		"error":  err.Error(),
	}
}

func (up *userPresenter) UserDeleteSuccessResponse() *fiber.Map {
	return &fiber.Map{
		"status": true,
		"data":   nil,
		"error":  nil,
	}
}
