package transports

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"my_wallet/api/endpoints"
	infraestructure_repository "my_wallet/api/respository/healtcheck"
	repository_user "my_wallet/api/respository/user"
	"my_wallet/api/services"
	"strings"
	"time"

	"my_wallet/api/entities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCustomErrorEncoder(t *testing.T) {
	testScenarios := []struct {
		name           string
		err            error
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "ErrLenghtPassword",
			err:            services.ErrLenghtPassword,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"minimum password length 8"}`,
		},
		{
			name:           "ErrUserNotfound",
			err:            services.ErrUserNotfound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"Error not found user"}`,
		},
		{
			name:           "ErrLenghPhone",
			err:            services.ErrLenghPhone,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Length of phone number 10"}`,
		},
		{
			name:           "ErrNameSpecialCharacters",
			err:            services.ErrNameSpecialCharacters,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"the name field must not contain special characters"}`,
		},
		{
			name:           "ErrTypeDNI",
			err:            services.ErrTypeDNI,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"ID type must be CC or NIT"}`,
		},
		{
			name:           "ErrInvalidCredentials",
			err:            services.ErrInvalidCredentials,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid email or password"}`,
		},
		{
			name:           "ErrDisbledUser",
			err:            repository_user.ErrDisbledUser,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Disabled user"}`,
		},
		{
			name:           "ErrUserNotfound",
			err:            repository_user.ErrUserNotfound,
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"Error not found user"}`,
		},
		{
			name:           "ErrNotasks",
			err:            repository_user.ErrNotasks,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"No tasks were deleted"}`,
		},
		{
			name:           "ErrInvalidCredentialsEndpoints",
			err:            endpoints.ErrInvalidCredentials,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Invalid email or password"}`,
		},
		{
			name:           "ErrInterfaceWrong",
			err:            endpoints.ErrInterfaceWrong,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Request interface type wrong"}`,
		},
		{
			name:           "ErrInterfaceWrong",
			err:            infraestructure_repository.ErrLoadingDatabase,
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Database unavailable"}`,
		},
		{
			name:           "nil error",
			err:            nil,
			expectedStatus: http.StatusNoContent,
			expectedBody:   `{"error":""}`,
		},
		{
			name:           "Unknown error",
			err:            errors.New("some unknown error"),
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"Internal server error."}`,
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			w := httptest.NewRecorder()

			// Act
			CustomErrorEncoder(context.Background(), tt.err, w)
			actualBody := strings.TrimSpace(w.Body.String())
			expectedBody := strings.TrimSpace(tt.expectedBody)

			// Assert
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.Equal(t, expectedBody, actualBody)
		})
	}
}

func TestNewHTTPHandler(t *testing.T) {
	logger := logrus.New()
	mocks := new(mockEndpoints)

	endpointss := endpoints.Endpoints{
		CreateUser:     makeCreateUserEndpoint(mocks),
		GetUser:        makeGetUserEndpoint(mocks),
		DeleteUser:     makeDeleteUserEndpoint(mocks),
		UpdateUser:     makeUpdateUserEndpoint(mocks),
		SoftDeleteUser: makeSoftDeleteUserEndpoint(mocks),
		Login:          makeLoginEndpoint(mocks),
		HealthCheck:    makeHealthCheckEndpoint(mocks),
	}
	mocks.On("CreateUser", mock.Anything, mock.Anything).Return(endpoints.CreateUserResponse{ID: "1"}, nil)
	mocks.On("GetUser", mock.Anything, mock.Anything).Return(endpoints.GetUserResponse{User: entities.User{
		Address:      "",
		DNI:          0,
		Email:        "",
		Enabled:      false,
		Name:         "",
		Password:     "",
		Phone:        0,
		TypeDNI:      "",
		Created_at:   time.Time{},
		Update_at:    time.Time{},
		RefreshToken: "",
		Token:        "",
	}}, nil)
	mocks.On("UpdateUser", mock.Anything, mock.Anything).Return(endpoints.UpdateUserREsponse{User: entities.User{
		Address:      "",
		DNI:          0,
		Email:        "",
		Enabled:      true,
		Name:         "",
		Password:     "",
		Phone:        0,
		TypeDNI:      "",
		Created_at:   time.Time{},
		Update_at:    time.Time{},
		RefreshToken: "",
		Token:        "",
	}}, nil)
	mocks.On("HealthCheck", mock.Anything, mock.Anything).Return(endpoints.HealtcheckDbResponse{Database: "ok"}, nil)

	handler := NewHTTPHandler(endpointss, logger)

	testScenarios := []struct {
		name           string
		method         string
		url            string
		body           interface{}
		expectedCode   int
		expectedOutput string
		authorization  string
	}{
		{
			name:           "Create User Success",
			method:         http.MethodPost,
			url:            "/user",
			body:           map[string]string{"email": "testuser<@gmail.com", "password": "password23242"},
			expectedCode:   http.StatusCreated,
			expectedOutput: `{"id":"1"}`,
		},
		{
			name:           "Get User Success",
			method:         http.MethodGet,
			url:            "/user/1",
			body:           nil,
			expectedCode:   http.StatusOK,
			expectedOutput: `{"user":{"Address":"","DNI":0,"Email":"","Enabled":false,"Name":"","Password":"","Phone":0,"TypeDNI":"","created_at":"0001-01-01T00:00:00Z","refresh_token":"","token":"","updated_at":"0001-01-01T00:00:00Z"}}`,
		},
		{
			name:           "Update User Success",
			method:         http.MethodPut,
			url:            "/user/update/1",
			body:           map[string]string{"name": "", "email": "", "password": ""},
			expectedCode:   http.StatusOK,
			expectedOutput: `{"user":{"Address":"","DNI":0,"Email":"","Enabled":true,"Name":"","Password":"","Phone":0,"TypeDNI":"","created_at":"0001-01-01T00:00:00Z","refresh_token":"","token":"","updated_at":"0001-01-01T00:00:00Z"}}`,
		},
		{
			name:           "Health Check Success",
			method:         http.MethodGet,
			url:            "/healthcheck",
			body:           nil,
			expectedCode:   http.StatusOK,
			expectedOutput: `{"database":"ok"}`,
		},
	}

	for _, tt := range testScenarios {
		t.Run(tt.name, func(t *testing.T) {

			// Prepare
			var buf bytes.Buffer
			if tt.body != nil {
				if err := json.NewEncoder(&buf).Encode(tt.body); err != nil {
					t.Fatalf("failed to encode request body: %v", err)
				}
			}
			req := httptest.NewRequest(tt.method, tt.url, &buf)
			if tt.authorization != "" {
				req.Header.Set("Authorization", tt.authorization)
			}

			// Act
			w := httptest.NewRecorder()
			handler.ServeHTTP(w, req)
			res := w.Result()

			// Assert
			assert.Equal(t, tt.expectedCode, res.StatusCode)
			body, _ := io.ReadAll(w.Body)
			assert.JSONEq(t, tt.expectedOutput, string(body))
		})
	}
}

func makeCreateUserEndpoint(m *mockEndpoints) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(endpoints.CreateUserRequest)
		return m.CreateUser(ctx, req)
	}
}

func makeGetUserEndpoint(m *mockEndpoints) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(endpoints.GetUserRequest)
		return m.GetUser(ctx, req)
	}
}

func makeDeleteUserEndpoint(m *mockEndpoints) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(endpoints.DeleteUserRequest)
		return m.DeleteUser(ctx, req)
	}
}

func makeUpdateUserEndpoint(m *mockEndpoints) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(endpoints.UpdateUserRequest)
		return m.UpdateUser(ctx, req)
	}
}

func makeSoftDeleteUserEndpoint(m *mockEndpoints) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(endpoints.SoftDeleteUserRequest)
		return m.SoftDeleteUser(ctx, req)
	}
}

func makeLoginEndpoint(m *mockEndpoints) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(endpoints.LoginUserRequest)
		return m.Login(ctx, req)
	}
}

func makeHealthCheckEndpoint(m *mockEndpoints) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(endpoints.HealtcheckDbRequest)
		return m.HealthCheck(ctx, req)
	}
}
