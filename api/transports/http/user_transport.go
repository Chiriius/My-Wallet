package transports

import (
	"context"
	"encoding/json"
	"my_wallet/api/endpoints"
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/sirupsen/logrus"
)

func NewHTTPHandler(endpoints endpoints.Endpoints, logger logrus.FieldLogger) http.Handler {

	m := http.NewServeMux()
	m.Handle("/user", httpTransport.NewServer(
		endpoints.CreateUser,
		decodeCreateUserRequest,
		encodeCreateUserResponse,
	))
	m.Handle("/user/{id}", httpTransport.NewServer(
		endpoints.GetUser,
		decodeGetUserRequest,
		encodeGetUserResponse,
	))
	m.Handle("/user/delete/{id}", httpTransport.NewServer(
		endpoints.DeleteUser,
		decodeDeleteUserRequest,
		encodeDeleteUserResponse,
	))
	m.Handle("/user/update/{id}", httpTransport.NewServer(
		endpoints.UpdateUser,
		decodeUpdateRequest,
		encodeUpdateUserResponse,
	))
	m.Handle("/user/soft/{id}", httpTransport.NewServer(
		endpoints.SoftDeleteUser,
		decodeSoftDeleteUserRequest,
		encodeSoftDeleteUserResponse,
	))
	return m
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
