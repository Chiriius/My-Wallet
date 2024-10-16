package services

import (
	"context"
	"errors"

	"my_wallet/api/entities"
	repository_user "my_wallet/api/respository/user"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserService(t *testing.T) {
	testScenarios := []struct {
		testName       string
		mock           *userServiceMock
		mockResponse   entities.User
		mockContext    context.Context
		mockValidator  *validator.Validate
		mockLogger     logrus.FieldLogger
		mockError      error
		configureMock  func(*userServiceMock, entities.User, error)
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName: "TestCreateUserService",
			mock:     &userServiceMock{},
			mockResponse: entities.User{
				DNI:         34,
				Name:        "Alexer",
				Email:       "alexer@gmail.com",
				Password:    "12345678",
				Address:     "cra 22a",
				Phone:       1234567899,
				StateActive: true,
			},
			mockContext:   context.Background(),
			mockValidator: validator.New(),
			mockLogger:    logrus.StandardLogger(),

			mockError: nil,
			configureMock: func(m *userServiceMock, mockResponse entities.User, mockError error) {
				m.On("CreateUser", mock.Anything, mock.AnythingOfType("entities.User")).Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{
				DNI:         34,
				Name:        "Alexer",
				Email:       "alexer@gmail.com",
				Password:    "12345678",
				Address:     "cra 22a",
				Phone:       1234567899,
				StateActive: true,
			},
			expectedError: nil,
		},
		{
			testName: "testSpecialCharacters",
			mock:     &userServiceMock{},
			mockResponse: entities.User{
				DNI:         34,
				Name:        "Alexer@@",
				Email:       "alexer@gmail.com",
				Password:    "12345678",
				Address:     "cra 22a",
				Phone:       1234567899,
				StateActive: true,
			},
			mockContext:   context.Background(),
			mockValidator: validator.New(),
			mockLogger:    logrus.StandardLogger(),

			mockError: errors.New("the name field must not contain special characters"),
			configureMock: func(m *userServiceMock, mockResponse entities.User, mockError error) {
				m.On("CreateUser", mock.Anything, mock.AnythingOfType("entities.User")).Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{},
			expectedError:  errors.New("the name field must not contain special characters"),
		},
		{
			testName: "testLenghtPassword",
			mock:     &userServiceMock{},
			mockResponse: entities.User{
				DNI:         34,
				Name:        "Alexer",
				Email:       "alexer@gmail.com",
				Password:    "123456",
				Address:     "cra 22a",
				Phone:       1234567,
				StateActive: true,
			},
			mockContext:   context.Background(),
			mockValidator: validator.New(),
			mockLogger:    logrus.StandardLogger(),

			mockError: errors.New("minimum password length 8 "),
			configureMock: func(m *userServiceMock, mockResponse entities.User, mockError error) {
				m.On("CreateUser", mock.Anything, mock.AnythingOfType("entities.User")).Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{},
			expectedError:  errors.New("minimum password length 8 "),
		},
		{
			testName: "testLenghtPhone",
			mock:     &userServiceMock{},
			mockResponse: entities.User{
				DNI:         34,
				Name:        "Alexer",
				Email:       "alexer@gmail.com",
				Password:    "123456232323",
				Address:     "cra 22a",
				Phone:       123,
				StateActive: true,
			},
			mockContext:   context.Background(),
			mockValidator: validator.New(),
			mockLogger:    logrus.StandardLogger(),

			mockError: errors.New("Length of phone number 10"),
			configureMock: func(m *userServiceMock, mockResponse entities.User, mockError error) {
				m.On("CreateUser", mock.Anything, mock.AnythingOfType("entities.User")).Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{},
			expectedError:  errors.New("Length of phone number 10"),
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {
			// Prepare
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse, tt.mockError)
			}

			service := &userService{
				repository: tt.mock,
				ctx:        tt.mockContext,
				validate:   tt.mockValidator,
				logger:     tt.mockLogger,
			}
			// Act
			result, err := service.CreateUser(tt.mockContext, tt.mockResponse)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}
func TestUpdateUserService(t *testing.T) {
	testScenarios := []struct {
		testName       string
		mock           *userServiceMock
		mockResponse   entities.User
		mockContext    context.Context
		mockValidator  *validator.Validate
		mockLogger     logrus.FieldLogger
		mockError      error
		configureMock  func(*userServiceMock, entities.User, error)
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName: "TestCreateUserService",
			mock:     &userServiceMock{},
			mockResponse: entities.User{
				DNI:         34,
				Name:        "Alexer",
				Email:       "alexer@gmail.com",
				Password:    "12345678",
				Address:     "cra 22a",
				Phone:       1234567899,
				StateActive: true,
			},
			mockContext:   context.Background(),
			mockValidator: validator.New(),
			mockLogger:    logrus.StandardLogger(),

			mockError: nil,
			configureMock: func(m *userServiceMock, mockResponse entities.User, mockError error) {
				m.On("UpdateUser", mock.Anything, mock.AnythingOfType("entities.User")).Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{
				DNI:         34,
				Name:        "Alexer",
				Email:       "alexer@gmail.com",
				Password:    "12345678",
				Address:     "cra 22a",
				Phone:       1234567899,
				StateActive: true,
			},
			expectedError: nil,
		},
		{
			testName: "TestLenghtPassword",
			mock:     &userServiceMock{},
			mockResponse: entities.User{
				DNI:         34,
				Name:        "Alexer",
				Email:       "alexer@gmail.com",
				Password:    "12345",
				Address:     "cra 22a",
				Phone:       1234567899,
				StateActive: true,
			},
			mockContext:   context.Background(),
			mockValidator: validator.New(),
			mockLogger:    logrus.StandardLogger(),

			mockError: errors.New("Minimum password length 8 "),
			configureMock: func(m *userServiceMock, mockResponse entities.User, mockError error) {
				m.On("UpdateUser", mock.Anything, mock.AnythingOfType("entities.User")).Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{},
			expectedError:  errors.New("Minimum password length 8 "),
		},
		{
			testName: "TestLenghtPhone",
			mock:     &userServiceMock{},
			mockResponse: entities.User{
				DNI:         34,
				Name:        "Alexer",
				Email:       "alexer@gmail.com",
				Password:    "12345678",
				Address:     "cra 22a",
				Phone:       4242,
				StateActive: true,
			},
			mockContext:   context.Background(),
			mockValidator: validator.New(),
			mockLogger:    logrus.StandardLogger(),

			mockError: errors.New("Length of phone number 10"),
			configureMock: func(m *userServiceMock, mockResponse entities.User, mockError error) {
				m.On("UpdateUser", mock.Anything, mock.AnythingOfType("entities.User")).Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{},
			expectedError:  errors.New("Length of phone number 10"),
		},
		{
			testName: "TestSpecialCharactersName",
			mock:     &userServiceMock{},
			mockResponse: entities.User{
				DNI:         34,
				Name:        "Alexer@@",
				Email:       "alexer@gmail.com",
				Password:    "12345678",
				Address:     "cra 22a",
				Phone:       1234567899,
				StateActive: true,
			},
			mockContext:   context.Background(),
			mockValidator: validator.New(),
			mockLogger:    logrus.StandardLogger(),

			mockError: errors.New("the name field must not contain special characters"),
			configureMock: func(m *userServiceMock, mockResponse entities.User, mockError error) {
				m.On("UpdateUser", mock.Anything, mock.AnythingOfType("entities.User")).Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{},
			expectedError:  errors.New("the name field must not contain special characters"),
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {
			// Prepare
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockResponse, tt.mockError)
			}

			service := &userService{
				repository: tt.mock,
				ctx:        tt.mockContext,
				validate:   tt.mockValidator,
				logger:     tt.mockLogger,
			}
			// Act
			result, err := service.UpdateUser(tt.mockContext, tt.mockResponse)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}
}
func TestDeleteUserService(t *testing.T) {
	testScenarios := []struct {
		testName       string
		mock           *userServiceMock
		mockContext    context.Context
		mockValidator  *validator.Validate
		mockLogger     logrus.FieldLogger
		mockID         string
		mockError      error
		configureMock  func(*userServiceMock, string, context.Context, error)
		expectedOutput error
		expectedError  error
	}{
		{
			testName:      "TestDeleteUserService",
			mock:          &userServiceMock{},
			mockContext:   context.Background(),
			mockValidator: validator.New(),
			mockLogger:    logrus.StandardLogger(),
			mockID:        "1",
			mockError:     nil,
			configureMock: func(m *userServiceMock, id string, ctx context.Context, mockError error) {
				m.On("DeleteUser", ctx, id).Return(mockError)
			},
			expectedOutput: nil,
			expectedError:  nil,
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {
			// Prepare
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockID, tt.mockContext, tt.mockError)
			}

			service := &userService{
				repository: tt.mock,
				ctx:        tt.mockContext,
				validate:   tt.mockValidator,
				logger:     logrus.StandardLogger(),
			}
			// Act
			err := service.DeleteUser(tt.mockContext, tt.mockID)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, err)
		})
	}
}
func TestGetUserService(t *testing.T) {

	testScenarios := []struct {
		testName       string
		mock           *userServiceMock
		mockResponse   entities.User
		mockContext    context.Context
		mockValidator  *validator.Validate
		mockLogger     logrus.FieldLogger
		mockError      error
		configureMock  func(*userServiceMock, context.Context, entities.User, error)
		expectedOutput entities.User
		expectedError  error
	}{
		{
			testName: "TestGetUserService",
			mock:     &userServiceMock{},
			mockResponse: entities.User{
				ID: "3",
			},
			mockContext:   context.Background(),
			mockValidator: validator.New(),
			mockLogger:    logrus.StandardLogger(),
			mockError:     nil,
			configureMock: func(m *userServiceMock, ctx context.Context, mockResponse entities.User, mockError error) {
				m.On("GetUser", ctx, "3").Return(mockResponse, mockError)
			},
			expectedOutput: entities.User{
				ID: "3",
			},
			expectedError: nil,
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {

			// Prepare
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockContext, tt.mockResponse, tt.mockError)
			}

			service := &userService{
				repository: tt.mock,
			}

			// Act
			result, err := service.GetUSer(tt.mockContext, tt.mockResponse.ID)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, result)
		})
	}

}
func TestSoftDeleteUserService(t *testing.T) {
	testScenarios := []struct {
		testName       string
		mock           *userServiceMock
		mockContext    context.Context
		mockValidator  *validator.Validate
		mockLogger     logrus.FieldLogger
		mockID         string
		mockError      error
		configureMock  func(*userServiceMock, string, context.Context, error)
		expectedOutput error
		expectedError  error
	}{
		{
			testName:      "TestSoftDeleteUserService",
			mock:          &userServiceMock{},
			mockContext:   context.Background(),
			mockValidator: validator.New(),
			mockLogger:    logrus.StandardLogger(),
			mockID:        "1",
			mockError:     nil,
			configureMock: func(m *userServiceMock, id string, ctx context.Context, mockError error) {
				m.On("SoftDeleteUser", ctx, id).Return(mockError)
			},
			expectedOutput: nil,
			expectedError:  nil,
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.testName, func(t *testing.T) {
			// Prepare
			if tt.configureMock != nil {
				tt.configureMock(tt.mock, tt.mockID, tt.mockContext, tt.mockError)
			}

			service := &userService{
				repository: tt.mock,
				ctx:        tt.mockContext,
				validate:   tt.mockValidator,
				logger:     logrus.StandardLogger(),
			}
			// Act
			err := service.SoftDeleteUser(tt.mockContext, tt.mockID)

			// Assert
			assert.Equal(t, tt.expectedError, err)
			assert.Equal(t, tt.expectedOutput, err)
		})
	}
}
func TestNewUserService(t *testing.T) {

	sm := &userServiceMock{}
	testScenarios := []struct {
		testName       string
		mockLogger     logrus.FieldLogger
		mockContext    context.Context
		mockRepo       repository_user.UserRepository
		expectedOutput *userService
	}{
		{
			testName:    "test MakeServerEndpoints",
			mockLogger:  logrus.StandardLogger(),
			mockContext: context.Background(),
			mockRepo:    sm,
			expectedOutput: &userService{
				repository: sm,
			},
		},
	}

	for _, tt := range testScenarios {

		// Prepare
		t.Run(tt.testName, func(t *testing.T) {

			// Act
			result := NewUserService(tt.mockRepo, tt.mockLogger, tt.mockContext)

			// Assert
			assert.NotNil(t, result)
			assert.Equal(t, tt.expectedOutput.repository, result.repository)
		})
	}
}
