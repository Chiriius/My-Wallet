package services

import (
	"context"
	"my_wallet/api/entities"

	"github.com/stretchr/testify/mock"
)

type userServiceMock struct {
	mock.Mock
}

func (m *userServiceMock) CreateUser(user entities.User, ctx context.Context) (entities.User, error) {
	r := m.Called(ctx, user)

	return r.Get(0).(entities.User), r.Error(1)
}
func (m *userServiceMock) DeleteUser(id string, ctx context.Context) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *userServiceMock) GetUser(id string, ctx context.Context) (entities.User, error) {
	r := m.Called(ctx, id)
	return r.Get(0).(entities.User), r.Error(1)
}
func (m *userServiceMock) UpdateUser(user entities.User, ctx context.Context) (entities.User, error) {
	r := m.Called(ctx, user)
	return r.Get(0).(entities.User), r.Error(1)
}
func (m *userServiceMock) SoftDeleteUser(id string, ctx context.Context) error {
	r := m.Called(ctx, id)
	return r.Error(0)
}

func (m *userServiceMock) GetUserByEmail(email string, ctx context.Context) (entities.User, error) {

	r := m.Called(ctx, email)
	return r.Get(0).(entities.User), r.Error(1)

}

func (m *userServiceMock) UpdateUserToken(userUpr entities.User, ctx context.Context) (entities.User, error) {
	r := m.Called(ctx, userUpr)
	return r.Get(0).(entities.User), r.Error(1)
}
