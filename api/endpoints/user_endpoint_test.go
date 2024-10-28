package endpoints

import (
	"context"
	"errors"
	"my_wallet/api/entities"
	"my_wallet/api/services"
	"testing"

	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMakeCreateUserEndpoint(t *testing.T) {

	testScenarios := []struct {
		testName        string
		endpoint        func(services.UserService, logrus.FieldLogger) endpoint.Endpoint
		service         services.UserService
		mockContext     context.Context
		mock            *serviceMock
		mockError       error
		mockLogger      logrus.FieldLogger
		configureMock   func(*serviceMock, entities.User, error)
		endpointRequest interface{}
		mockResponse    entities.User
		expectedOutput  CreateUserResponse
		expectedError   error
	}{
		{
			testName: "test MakeCreateUserEndpoint",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "5",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("CreateUser", mock.Anything, mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput: CreateUserResponse{
				ID: "5",
			},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   nil,
			endpointRequest: CreateUserRequest{Password: "12345678", Email: "alexer@gmail.com", Name: "Alexer Maestre"},
		},
		{
			testName: "test MakeCreateUserEndpoint with error Interface type wrong",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "5",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("CreateUser", mock.Anything, mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput:  CreateUserResponse{},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   ErrInterfaceWrong,
			endpointRequest: CreateUserResponse{},
		},
		{
			testName: "test MakeCreateUserEndpoint with error in the service",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "6",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("CreateUser", mock.Anything, mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput:  CreateUserResponse{},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   ErrInvalidCredentials,
			endpointRequest: CreateUserRequest{Password: "15678", Email: "alexer@gmail.com", Name: "Alexer Maestre"},
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			tt.endpoint = MakeCreateUserEndpoint
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse, tt.mockError)
			}
			ctx := context.TODO()

			// Act
			result, err := tt.endpoint(tt.mock, tt.mockLogger)(ctx, tt.endpointRequest)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func TestMakeGetUserEndpoint(t *testing.T) {

	testScenarios := []struct {
		testName        string
		endpoint        func(services.UserService, logrus.FieldLogger) endpoint.Endpoint
		service         services.UserService
		mockContext     context.Context
		mock            *serviceMock
		mockError       error
		mockLogger      logrus.FieldLogger
		configureMock   func(*serviceMock, entities.User, error)
		endpointRequest interface{}
		mockResponse    entities.User
		expectedOutput  GetUserResponse
		expectedError   error
	}{
		{
			testName: "test MakeGetUserEndpoint",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "5",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("GetUSer", mock.Anything, mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput:  GetUserResponse{User: entities.User{ID: "5"}},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   nil,
			endpointRequest: GetUserRequest{ID: "5"},
		},
		{
			testName: "test MakeGetUserEndpoint with error Interface type wrong",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "5",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("GetUSer", mock.Anything, mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput:  GetUserResponse{},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   ErrInterfaceWrong,
			endpointRequest: GetUserResponse{},
		},
		{
			testName: "test MakeGetUserEndpoint with error in the service",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "6",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("GetUSer", mock.Anything, mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput:  GetUserResponse{},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   errors.New("id is 6 "),
			endpointRequest: GetUserRequest{ID: "6"},
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			tt.endpoint = MakeGetUserEndpoint
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse, tt.mockError)
			}
			ctx := context.TODO()

			// Act
			result, err := tt.endpoint(tt.mock, tt.mockLogger)(ctx, tt.endpointRequest)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func TestMakeUpdateUserEndpoint(t *testing.T) {

	testScenarios := []struct {
		testName        string
		endpoint        func(services.UserService, logrus.FieldLogger) endpoint.Endpoint
		service         services.UserService
		mockContext     context.Context
		mock            *serviceMock
		mockError       error
		mockLogger      logrus.FieldLogger
		configureMock   func(*serviceMock, entities.User, error)
		endpointRequest interface{}
		mockResponse    entities.User
		expectedOutput  UpdateUserREsponse
		expectedError   error
	}{
		{
			testName: "test MakeUpdateUserEndpoint",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "5",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("UpdateUser", mock.Anything, mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput:  UpdateUserREsponse{User: entities.User{ID: "5"}},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   nil,
			endpointRequest: UpdateUserRequest{ID: "5", Password: "dasdasdasdad"},
		},
		{
			testName: "test MakeGetUserEndpoint with error Interface type wrong",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "5",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("UpdateUser", mock.Anything, mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput:  UpdateUserREsponse{},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   ErrInterfaceWrong,
			endpointRequest: UpdateUserREsponse{},
		},
		{
			testName: "test MakeGetUserEndpoint with error in the service",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID: "6",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("UpdateUser", mock.Anything, mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput:  UpdateUserREsponse{},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   ErrInvalidCredentials,
			endpointRequest: UpdateUserRequest{ID: "6", Password: "S23"},
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			tt.endpoint = MakeUpdateUserEndpoint
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse, tt.mockError)
			}
			ctx := context.TODO()

			// Act
			result, err := tt.endpoint(tt.mock, tt.mockLogger)(ctx, tt.endpointRequest)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func TestMakeSoftDeleteUserEndpoint(t *testing.T) {

	testScenarios := []struct {
		testName        string
		endpoint        func(services.UserService, logrus.FieldLogger) endpoint.Endpoint
		service         services.UserService
		mockContext     context.Context
		mock            *serviceMock
		mockError       error
		mockLogger      logrus.FieldLogger
		configureMock   func(*serviceMock, error)
		endpointRequest interface{}
		mockResponse    error
		expectedOutput  SoftDeleteUserResponse
		expectedError   error
	}{
		{
			testName:     "test MakeSoftDeleteUserEndpoint",
			mock:         &serviceMock{},
			mockResponse: nil,
			configureMock: func(m *serviceMock, mockResponse error) {
				m.On("SoftDeleteUser", mock.Anything, mock.Anything).Return(mockResponse)
			},
			expectedOutput:  SoftDeleteUserResponse{},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   nil,
			endpointRequest: SoftDeleteUserRequest{ID: "5"},
		},
		{
			testName:     "test MakeSoftDeleteUserEndpoint with error Interface type wrong",
			mock:         &serviceMock{},
			mockResponse: nil,
			configureMock: func(m *serviceMock, mockResponse error) {
				m.On("SoftDeleteUser", mock.Anything, mock.Anything).Return(mockResponse)
			},
			expectedOutput:  SoftDeleteUserResponse{},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   ErrInterfaceWrong,
			endpointRequest: GetUserRequest{ID: "5"},
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			tt.endpoint = MakeSoftDeleteUserEndpoint
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse)
			}
			ctx := context.TODO()

			// Act
			result, err := tt.endpoint(tt.mock, tt.mockLogger)(ctx, tt.endpointRequest)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func TestMakeDeleteUserEndpoint(t *testing.T) {

	testScenarios := []struct {
		testName        string
		endpoint        func(services.UserService, logrus.FieldLogger) endpoint.Endpoint
		service         services.UserService
		mockContext     context.Context
		mock            *serviceMock
		mockError       error
		mockLogger      logrus.FieldLogger
		configureMock   func(*serviceMock, error)
		endpointRequest interface{}
		mockResponse    error
		expectedOutput  DeleteUserResponse
		expectedError   error
	}{
		{
			testName:     "test MakeDeleteUserEndpoint",
			mock:         &serviceMock{},
			mockResponse: nil,
			configureMock: func(m *serviceMock, mockResponse error) {
				m.On("DeleteUser", mock.Anything, mock.Anything).Return(mockResponse)
			},
			expectedOutput:  DeleteUserResponse{},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   nil,
			endpointRequest: DeleteUserRequest{ID: "5"},
		},
		{
			testName:     "test MakeDeleteUserEndpoint with error Interface type wrong",
			mock:         &serviceMock{},
			mockResponse: nil,
			configureMock: func(m *serviceMock, mockResponse error) {
				m.On("DeleteUser", mock.Anything, mock.Anything).Return(mockResponse)
			},
			expectedOutput:  DeleteUserResponse{},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   ErrInterfaceWrong,
			endpointRequest: SoftDeleteUserRequest{ID: "5"},
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			tt.endpoint = MakeDeleteUserEndpoint
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse)
			}
			ctx := context.TODO()

			// Act
			result, err := tt.endpoint(tt.mock, tt.mockLogger)(ctx, tt.endpointRequest)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}

func TestMakeServerEndpoints(t *testing.T) {

	sm := &serviceMock{}
	testScenarios := []struct {
		testName       string
		mock           *serviceMock
		expectedOutput Endpoints
	}{
		{
			testName: "test MakeServerEndpoints",
			mock:     sm,
			expectedOutput: Endpoints{
				CreateUser:     MakeCreateUserEndpoint(sm, logrus.StandardLogger()),
				GetUser:        MakeGetUserEndpoint(sm, logrus.StandardLogger()),
				DeleteUser:     MakeDeleteUserEndpoint(sm, logrus.StandardLogger()),
				UpdateUser:     MakeUpdateUserEndpoint(sm, logrus.StandardLogger()),
				SoftDeleteUser: MakeSoftDeleteUserEndpoint(sm, logrus.StandardLogger()),
				Login:          MakeLoginEndpoint(sm, logrus.StandardLogger()),
			},
		},
	}

	for _, tt := range testScenarios {

		// Prepare
		t.Run(tt.testName, func(t *testing.T) {

			// Act
			result := MakeServerEndpoints(tt.mock, tt.mock, logrus.StandardLogger())

			// Assert
			assert.NotNil(t, result.CreateUser)
			assert.NotNil(t, result.SoftDeleteUser)
			assert.NotNil(t, result.DeleteUser)
			assert.NotNil(t, result.GetUser)
			assert.NotNil(t, result.UpdateUser)
			assert.NotNil(t, result.Login)
		})
	}
}

func TestMakeLoginUserEndpoint(t *testing.T) {

	testScenarios := []struct {
		testName        string
		endpoint        func(services.UserService, logrus.FieldLogger) endpoint.Endpoint
		service         services.UserService
		mockContext     context.Context
		mock            *serviceMock
		mockError       error
		mockLogger      logrus.FieldLogger
		configureMock   func(*serviceMock, entities.User, error)
		endpointRequest interface{}
		mockResponse    entities.User
		expectedOutput  LoginUserResponse
		expectedError   error
	}{
		{
			testName: "test MakeLoginEndpoint",
			mock:     &serviceMock{},
			mockResponse: entities.User{
				ID:    "5",
				Token: "33s",
			},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("Login", mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput: LoginUserResponse{
				StateLogin: true,
				Token:      "33s",
			},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   nil,
			endpointRequest: LoginUserRequest{Password: "12345678", Email: "alexer@gmail.com"},
		},
		{
			testName:     "test MakeLoginEndpoint with error Interface type wrong",
			mock:         &serviceMock{},
			mockResponse: entities.User{},
			configureMock: func(m *serviceMock, mockResponse entities.User, mockError error) {
				m.On("Login", mock.Anything, mock.Anything, mock.Anything).Return(mockResponse, mockError)
			},
			expectedOutput:  LoginUserResponse{},
			mockContext:     context.Background(),
			mockLogger:      logrus.StandardLogger(),
			expectedError:   ErrInterfaceWrong,
			endpointRequest: SoftDeleteUserRequest{ID: "5"},
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			tt.endpoint = MakeLoginEndpoint
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse, tt.mockError)
			}
			ctx := context.TODO()

			// Act
			result, err := tt.endpoint(tt.mock, tt.mockLogger)(ctx, tt.endpointRequest)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}
