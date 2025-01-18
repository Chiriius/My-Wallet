package endpoints

import (
	"context"
	"errors"
	"my_wallet/api/entities"

	"github.com/stretchr/testify/mock"
)

type serviceMock struct {
	mock.Mock
}

func (s *serviceMock) CreateUser(ctx context.Context, user entities.User) (entities.User, error) {
	if len(user.Password) < 8 {
		return entities.User{}, ErrInvalidCredentials
	}
	r := s.Called(ctx, user)
	return r.Get(0).(entities.User), r.Error(1)

}

func (s *serviceMock) GetUSer(ctx context.Context, id string) (entities.User, error) {
	if id == "6" {
		return entities.User{}, errors.New("id is 6 ")
	}
	r := s.Called(ctx, id)
	return r.Get(0).(entities.User), r.Error(1)

}

func (s *serviceMock) UpdateUser(ctx context.Context, user entities.User) (entities.User, error) {
	if len(user.Password) < 8 {
		return entities.User{}, ErrInvalidCredentials
	}
	r := s.Called(ctx, user)
	return r.Get(0).(entities.User), r.Error(1)
}

func (s *serviceMock) SoftDeleteUser(ctx context.Context, id string) error {
	r := s.Called(ctx, id)
	return r.Error(0)
}

func (s *serviceMock) DeleteUser(ctx context.Context, id string) error {
	r := s.Called(ctx, id)
	return r.Error(0)
}

func (s *serviceMock) Login(ctx context.Context, email string, password string) (bool, entities.User, error) {
	r := s.Called(ctx, email, password)
	return true, r.Get(0).(entities.User), r.Error(1)

}
func (s *serviceMock) GetHealtcheck(ctx context.Context) (bool, error) {
	return true, nil
}
