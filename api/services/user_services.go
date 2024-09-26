package services

import (
	"my_wallet/api/entities"
	repository_user "my_wallet/api/respository/user"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(use entities.User) (entities.User, error)
}

type userService struct {
	repository repository_user.MongoUserRepositoy
	logger     logrus.FieldLogger
}

func NewUserService(repo repository_user.MongoUserRepositoy, logger logrus.FieldLogger) *userService {
	return &userService{
		repository: repo,
		logger:     logger,
	}
}

func (s *userService) CreateUser(user entities.User) (entities.User, error) {
	user.UID = uuid.NewString()
	return s.repository.CreateUser(user)

}
