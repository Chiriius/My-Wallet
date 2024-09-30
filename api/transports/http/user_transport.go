package transports

import (
	"context"
	"encoding/json"
	"my_wallet/api/endpoints"
	"net/http"

	httpTransport "github.com/go-kit/kit/transport/http"
	"github.com/sirupsen/logrus"
)

func NewHTTPHandler(endpoints endpoints.Endpoints) http.Handler {

	m := http.NewServeMux()
	m.Handle("/user", httpTransport.NewServer(
		endpoints.CreateUser,
		decodeCreateUserRequest,
		encodeCreateUserResponse,
	))
	return m
}

func encodeCreateUserResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	logrus.Infoln("User with Http:", response)
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req endpoints.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}
