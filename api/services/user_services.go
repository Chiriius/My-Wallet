package services

import (
	"context"
	"my_wallet/api/entities"
	repository_user "my_wallet/api/respository/user"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(use entities.User) (entities.User, error)
	GetUSer(id string) (entities.User, error)
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
		return entities.User{}, err
	}
	return s.repository.CreateUser(user, s.ctx)

}

func (s *userService) GetUSer(id string) (entities.User, error) {

	return s.repository.GetUser(id, s.ctx)
}
