package transports

import (
	"context"
	"encoding/json"
	"errors"
	"my_wallet/api/endpoints"
	repository_user "my_wallet/api/respository/user"
	"my_wallet/api/services"
	"my_wallet/api/utils/jwt"
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/sirupsen/logrus"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func NewHTTPHandler(endpoints endpoints.Endpoints, logger logrus.FieldLogger) http.Handler {

	m := http.NewServeMux()
	m.Handle("/user", httpTransport.NewServer(
		endpoints.CreateUser,
		decodeCreateUserRequest,
		encodeCreateUserResponse,
		httpTransport.ServerErrorEncoder(CustomErrorEncoder),
	))
	m.Handle("/user/login", httpTransport.NewServer(
		endpoints.Login,
		decodeLoginUserRequest,
		encodeLoginUserResponse,
		httpTransport.ServerErrorEncoder(CustomErrorEncoder),
	))
	m.Handle("/user/{id}", httpTransport.NewServer(
		endpoints.GetUser,
		decodeGetUserRequest,
		encodeGetUserResponse,
		httpTransport.ServerErrorEncoder(CustomErrorEncoder),
	))
	m.Handle("/user/delete/{id}", jwt.JWTMiddleware(httpTransport.NewServer(
		endpoints.DeleteUser,
		decodeDeleteUserRequest,
		encodeDeleteUserResponse,
		httpTransport.ServerErrorEncoder(CustomErrorEncoder),
	)))
	m.Handle("/user/update/{id}", httpTransport.NewServer(
		endpoints.UpdateUser,
		decodeUpdateRequest,
		encodeUpdateUserResponse,
		httpTransport.ServerErrorEncoder(CustomErrorEncoder),
	))
	m.Handle("/user/soft/{id}", jwt.JWTMiddleware(httpTransport.NewServer(
		endpoints.SoftDeleteUser,
		decodeSoftDeleteUserRequest,
		encodeSoftDeleteUserResponse,
		httpTransport.ServerErrorEncoder(CustomErrorEncoder),
	)))
	return m
}

func CustomErrorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	var statusCode int
	var errorMessage string

	switch {
	case errors.Is(err, services.ErrLenghtPassword):
		statusCode = http.StatusBadRequest
		errorMessage = services.ErrLenghtPassword.Error()
	case errors.Is(err, services.ErrLenghPhone):
		statusCode = http.StatusBadRequest
		errorMessage = services.ErrLenghPhone.Error()
	case errors.Is(err, services.ErrNameSpecialCharacters):
		statusCode = http.StatusBadRequest
		errorMessage = services.ErrNameSpecialCharacters.Error()
	case errors.Is(err, services.ErrTypeDNI):
		statusCode = http.StatusBadRequest
		errorMessage = services.ErrTypeDNI.Error()
	case errors.Is(err, services.ErrUserNotfound):
		statusCode = http.StatusNotFound
		errorMessage = services.ErrUserNotfound.Error()
	case errors.Is(err, services.ErrInvalidCredentials):
		statusCode = http.StatusBadRequest
		errorMessage = services.ErrInvalidCredentials.Error()
	case errors.Is(err, repository_user.ErrDisbledUser):
		statusCode = http.StatusBadRequest
		errorMessage = repository_user.ErrDisbledUser.Error()
	case errors.Is(err, repository_user.ErrUserNotfound):
		statusCode = http.StatusNotFound
		errorMessage = repository_user.ErrUserNotfound.Error()
	case errors.Is(err, nil):
		statusCode = http.StatusNoContent
		errorMessage = ""
	case errors.Is(err, repository_user.ErrNotasks):
		statusCode = http.StatusBadRequest
		errorMessage = repository_user.ErrNotasks.Error()
	case errors.Is(err, endpoints.ErrInvalidCredentials):
		statusCode = http.StatusBadRequest
		errorMessage = endpoints.ErrInvalidCredentials.Error()
	case errors.Is(err, endpoints.ErrInterfaceWrong):
		statusCode = http.StatusBadRequest
		errorMessage = endpoints.ErrInterfaceWrong.Error()

	default:
		statusCode = http.StatusInternalServerError
		errorMessage = "Internal server error."
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Error: errorMessage})
}

func encodeLoginUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	w.WriteHeader(http.StatusAccepted)
	return json.NewEncoder(w).Encode(response)
}

func encodeCreateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}
func encodeGetUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if response == nil {
		w.WriteHeader(http.StatusNotFound)
		return nil
	}
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func encodeDeleteUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func encodeUpdateUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}

func encodeSoftDeleteUserResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	w.WriteHeader(http.StatusNoContent)
	return nil
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}
	return req, nil
}
func decodeLoginUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.LoginUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.GetUserRequest
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	req.ID = r.PathValue("id")

	return req, nil
}

func decodeDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.DeleteUserRequest

	req.ID = r.PathValue("id")

	return req, nil
}

func decodeSoftDeleteUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.SoftDeleteUserRequest

	req.ID = r.PathValue("id")

	return req, nil
}

func decodeUpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.UpdateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	req.ID = r.PathValue("id")
	if req.ID == "" {
		req.ID = r.PathValue("id")
	}
	return req, err
}
