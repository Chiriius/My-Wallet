package services

import (
	"context"
	"errors"
	"fmt"
	"my_wallet/api/entities"
	repository_user "my_wallet/api/respository/user"
	"my_wallet/api/utils"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(ctx context.Context, use entities.User) (entities.User, error)
	GetUSer(ctx context.Context, id string) (entities.User, error)
	DeleteUser(ctx context.Context, id string) error
	SoftDeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, user entities.User) (entities.User, error)
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

func (s *userService) CreateUser(ctx context.Context, user entities.User) (entities.User, error) {

	if err := s.validate.Struct(user); err != nil {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error:", err)
		return entities.User{}, err
	}
	phoneStr := fmt.Sprintf("%d", user.Phone)
	if len(user.Password) < 8 {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error: minimum password length 8")
		return entities.User{}, errors.New("minimum password length 8 ")

	}
	if len(phoneStr) != 10 {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error: Length of phone number 10")
		return entities.User{}, errors.New("Length of phone number 10")

	}
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	if !re.MatchString(user.Name) {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error: the name field must not contain special characters")
		return entities.User{}, errors.New("the name field must not contain special characters")
	}

	passwordHashed, err := utils.HashPassword(user.Password)
	if err != nil {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error: Error hashing the password")
		return entities.User{}, errors.New("Error hashing the password")
	}
	user.Password = passwordHashed
	s.logger.Info("Layer: user_services", "Method: CreateUser", "User:", user)

	return s.repository.CreateUser(user, ctx)

}

func (s *userService) GetUSer(ctx context.Context, id string) (entities.User, error) {

	return s.repository.GetUser(id, ctx)
}

func (s *userService) UpdateUser(ctx context.Context, user entities.User) (entities.User, error) {
	if err := s.validate.Struct(user); err != nil {
		s.logger.Errorln("Layer: user_services", "Method: UpdateUser", "Error:", err)
		return entities.User{}, err
	}
	phoneStr := fmt.Sprintf("%d", user.Phone)
	if len(user.Password) < 8 {
		s.logger.Errorln("Layer: user_services", "Method: UpdateUser", "Error: minimum password length 8")
		return entities.User{}, errors.New("Minimum password length 8 ")

	}
	if len(phoneStr) != 10 {
		s.logger.Errorln("Layer: user_services", "Method: UpdateUser", "Error: Length of phone number 10")
		return entities.User{}, errors.New("Length of phone number 10")

	}
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	if !re.MatchString(user.Name) {
		s.logger.Errorln("Layer: user_services", "Method: UpdateUser", "Error:the name field must not contain special characters")
		return entities.User{}, errors.New("the name field must not contain special characters")
	}
	passwordHashed, err := utils.HashPassword(user.Password)
	if err != nil {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error: Error hashing the password")
		return entities.User{}, errors.New("Error hashing the password")
	}
	user.Password = passwordHashed
	return s.repository.UpdateUser(user, ctx)
}

func (s *userService) SoftDeleteUser(ctx context.Context, id string) error {
	return s.repository.SoftDeleteUser(id, ctx)
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {

	return s.repository.DeleteUser(id, ctx)
}
