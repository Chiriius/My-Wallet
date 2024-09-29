package services

import (
	"context"
	"my_wallet/api/entities"
	repository_user "my_wallet/api/respository/user"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(use entities.User) (entities.User, error)
}

type userService struct {
	ctx        context.Context
	repository repository_user.UserRepository
	logger     logrus.FieldLogger
}

func NewUserService(repo repository_user.UserRepository, logger logrus.FieldLogger, ctx context.Context) *userService {
	return &userService{
		ctx:        ctx,
		repository: repo,
		logger:     logger,
	}
}

func (s *userService) CreateUser(user entities.User) (entities.User, error) {
	user.UID = uuid.NewString()
	return s.repository.CreateUser(user, s.ctx)

}
