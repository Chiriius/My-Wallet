package services

import (
	"context"
	"errors"
	"fmt"
	"my_wallet/api/entities"
	repository_user "my_wallet/api/respository/user"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(use entities.User) (entities.User, error)
	GetUSer(id string) (entities.User, error)
	DeleteUser(id string) error
	SoftDeleteUser(id string) error
	UpdateUser(user entities.User) (entities.User, error)
}

type userService struct {
	ctx        context.Context
	repository repository_user.UserRepository
	logger     logrus.FieldLogger
	validate   *validator.Validate
}

func NewUserService(repo repository_user.UserRepository, logger logrus.FieldLogger, ctx context.Context) *userService {
	return &userService{
		ctx:        ctx,
		repository: repo,
		logger:     logger,
		validate:   validator.New(),
	}
}

func (s *userService) CreateUser(user entities.User) (entities.User, error) {

	if err := s.validate.Struct(user); err != nil {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error:", err)
		return entities.User{}, err
	}
	phoneStr := fmt.Sprintf("%d", user.Phone)
	if len(user.Password) < 8 {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error: minimum password length 8")
		return entities.User{}, errors.New("minimum password length 8 and phone length 10")

	}
	if len(phoneStr) != 10 {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error: Phone length 10")
		return entities.User{}, errors.New("minimum password length 8 and phone length 10")

	}
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	if !re.MatchString(user.Name) {
		return entities.User{}, errors.New("the name field must not contain special characters")
	}

	return s.repository.CreateUser(user, s.ctx)

}

func (s *userService) GetUSer(id string) (entities.User, error) {

	return s.repository.GetUser(id, s.ctx)
}

func (s *userService) UpdateUser(user entities.User) (entities.User, error) {
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	if !re.MatchString(user.Name) {
		return entities.User{}, errors.New("the name field must not contain special characters")
	}

	return s.repository.UpdateUser(user, s.ctx)
}

func (s *userService) SoftDeleteUser(id string) error {
	return s.repository.SoftDeleteUser(id, s.ctx)
}

func (s *userService) DeleteUser(id string) error {

	return s.repository.DeleteUser(id, s.ctx)
}
