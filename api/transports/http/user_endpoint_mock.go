package transports

import (
	"context"
	"my_wallet/api/endpoints"

	"github.com/stretchr/testify/mock"
)

type mockEndpoints struct {
	mock.Mock
}

func (m *mockEndpoints) CreateUser(ctx context.Context, request endpoints.CreateUserRequest) (response endpoints.CreateUserResponse, err error) {
	args := m.Called(ctx, request)
	return args.Get(0).(endpoints.CreateUserResponse), args.Error(1)
}

func (m *mockEndpoints) Login(ctx context.Context, request endpoints.LoginUserRequest) (response endpoints.LoginUserResponse, err error) {
	args := m.Called(ctx, request)
	return args.Get(0).(endpoints.LoginUserResponse), args.Error(1)
}

func (m *mockEndpoints) GetUser(ctx context.Context, request endpoints.GetUserRequest) (response endpoints.GetUserResponse, err error) {
	args := m.Called(ctx, request)
	return args.Get(0).(endpoints.GetUserResponse), args.Error(1)
}

func (m *mockEndpoints) DeleteUser(ctx context.Context, request endpoints.DeleteUserRequest) (response endpoints.DeleteUserResponse, err error) {
	args := m.Called(ctx, request)
	return args.Get(0).(endpoints.DeleteUserResponse), args.Error(1)
}

func (m *mockEndpoints) UpdateUser(ctx context.Context, request endpoints.UpdateUserRequest) (response endpoints.UpdateUserREsponse, err error) {
	args := m.Called(ctx, request)
	return args.Get(0).(endpoints.UpdateUserREsponse), args.Error(1)
}

func (m *mockEndpoints) SoftDeleteUser(ctx context.Context, request endpoints.SoftDeleteUserRequest) (response endpoints.SoftDeleteUserResponse, err error) {
	args := m.Called(ctx, request)
	return args.Get(0).(endpoints.SoftDeleteUserResponse), args.Error(1)
}

func (m *mockEndpoints) HealthCheck(ctx context.Context, request endpoints.HealtcheckDbRequest) (response endpoints.HealtcheckDbResponse, err error) {
	args := m.Called(ctx, request)
	return args.Get(0).(endpoints.HealtcheckDbResponse), args.Error(1)
}
