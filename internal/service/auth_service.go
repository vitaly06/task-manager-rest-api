package service

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/vitaly06/task-manager-rest-api/internal/domain"
	"github.com/vitaly06/task-manager-rest-api/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService struct {
	userR *repository.UserRepository
}

var EmailAlreadyExists = errors.New("Пользователь с данной почтой уже существует")
var EmailNotRegisted = errors.New("Пользователь с данной почтой не зарегистрирован")
var PasswordsNotMatch = errors.New("Пароли не совпадают")

func NewAuthService(userR *repository.UserRepository) *AuthService {
	return &AuthService{
		userR: userR,
	}
}

func (s *AuthService) SignUp(email, password string) error {
	_, err := s.userR.GetByEmail(email)

	if err == nil {
		return EmailAlreadyExists
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	hashedPassword, err := hashPassword(password)

	if err != nil {
		return err
	}

	// create user
	user := domain.User{
		ID:        uuid.New(),
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}

	if err := s.userR.Create(&user); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) SignIn(email, password string) error {
	checkUser, err := s.userR.GetByEmail(email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return EmailNotRegisted
		}

		return err
	}

	if !ComparePasswords(password, checkUser.Password) {
		return PasswordsNotMatch
	}

	return nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(bytes), err
}

func ComparePasswords(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))

	return err == nil
}
