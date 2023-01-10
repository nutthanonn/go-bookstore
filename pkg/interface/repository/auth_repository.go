package repository

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nutthanonn/go-clean-architecture/api/infrastructure/datastore"
	"github.com/nutthanonn/go-clean-architecture/pkg/entities"
	"github.com/nutthanonn/go-clean-architecture/pkg/models"
	"github.com/nutthanonn/go-clean-architecture/pkg/usecase/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) repository.AuthRepository {
	return &authRepository{db}
}

func (ar *authRepository) SignIn(data *models.SignIn) (bool, error) {
	var user *entities.Customers

	if err := ar.db.Model(&user).First(&user, "email = ?", data.Email).Error; err != nil {
		return false, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return false, err
	}

	return true, nil
}

func (ar *authRepository) CreateToken(data *entities.Customers) (*models.Token, error) {
	datastore.LoadEnv()
	TOKEN_SECRET_KEY := os.Getenv("TOKEN_SECRET_KEY")

	var msgToken models.Token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["sub"] = data.Email
	claims["exp"] = time.Now().Add(time.Minute * 60).Unix()
	claims["iat"] = time.Now().Unix()

	t, err := token.SignedString([]byte(TOKEN_SECRET_KEY))

	if err != nil {
		return nil, err
	}

	msgToken.AccessToken = t

	return &msgToken, nil
}
