package endpoints

import (
	"context"
	"errors"
	"my_wallet/api/entities"
	"my_wallet/api/services"

	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
)

type CreateUserRequest struct {
	DNI      int
	Name     string
	Email    string
	Password string
	Address  string
	Phone    int
}

type CreateUserResponse struct {
	ID  string `json:"id,omitempty"`
	Err string `json:"error,omitempty"`
}

type GetUserRequest struct {
	ID string
}

type GetUserResponse struct {
	User entities.User
	Err  string `json:"error,omitempty"`
}

type Endpoints struct {
	CreateUser endpoint.Endpoint
	GetUser    endpoint.Endpoint
}

func MakeServerEndpoints(s services.UserService, logger logrus.FieldLogger) Endpoints {
	return Endpoints{
		CreateUser: MakeCreateUserEndpoint(s, logger),
		GetUser:    MakeGetUserEdpoint(s, logger),
	}
}

func MakeCreateUserEndpoint(s services.UserService, logger logrus.FieldLogger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var req CreateUserRequest
		var ok bool = false
		if req, ok = request.(CreateUserRequest); !ok {
			logger.Errorln("Layer:user_endpoint", "Method:MakeCreateUserEndpoint", "Error: Interface type wrong")
			return nil, errors.ErrUnsupported // aqui va el error personalizado de interfaz equivocada
		}
		user := entities.User{
			DNI:      req.DNI,
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
			Address:  req.Address,
			Phone:    req.Phone,
		}
		serviceUser, err := s.CreateUser(user)
		if err != nil {
			logger.Errorln("Layer:user_endpoint", "Method:MakeCreateUserEndpoint", err)
			return CreateUserResponse{}, err
		}
		logger.Infoln("Layer:user_endpoint", "Method:MakeCreateUserEndpoint", "Response:", CreateUserResponse{ID: serviceUser.ID})
		return CreateUserResponse{ID: serviceUser.ID}, nil

	}
}

func MakeGetUserEdpoint(s services.UserService, logger logrus.FieldLogger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var req GetUserRequest
		var ok bool = false

		if req, ok = request.(GetUserRequest); !ok {
			return nil, errors.ErrUnsupported // aqui va el error personalizado de interfaz equivocada
		}

		user, err := s.GetUSer(req.ID)

		if err != nil {
			return GetUserResponse{}, err
		}

		return GetUserResponse{User: user}, nil

	}

}
