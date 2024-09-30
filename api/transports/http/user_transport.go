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
		encodeGenericResponse,
	))
	m.Handle("/user/get", httpTransport.NewServer(
		endpoints.GetUser,
		decodeGetUserRequest,
		encodeGenericResponse,
	))
	return m
}

func encodeGenericResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {

	logrus.Infoln("User with Http:", response)
	return json.NewEncoder(w).Encode(response)
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
	req.ID = r.FormValue("id")

	return req, nil
}
