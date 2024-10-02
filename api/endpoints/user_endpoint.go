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

type UpdateUserRequest struct {
	ID       string
	DNI      int
	Name     string
	Email    string
	Password string
	Address  string
	Phone    int
}

type UpdateUserREsponse struct {
	User entities.User
	Err  string `json:"error,omitempty"`
}

type GetUserRequest struct {
	ID string
}

type GetUserResponse struct {
	User entities.User
	Err  string `json:"error,omitempty"`
}

type DeleteUserRequest struct {
	ID string
}

type DeleteUserResponse struct {
	Err string `json:"error,omitempty"`
}

type SoftDeleteUserRequest struct {
	ID string
}

type SoftDeleteUserResponse struct {
	Err string `json:"error,omitempty"`
}
type Endpoints struct {
	CreateUser     endpoint.Endpoint
	GetUser        endpoint.Endpoint
	DeleteUser     endpoint.Endpoint
	UpdateUser     endpoint.Endpoint
	SoftDeleteUser endpoint.Endpoint
}

func MakeServerEndpoints(s services.UserService, logger logrus.FieldLogger) Endpoints {
	return Endpoints{
		CreateUser:     MakeCreateUserEndpoint(s, logger),
		GetUser:        MakeGetUserEdpoint(s, logger),
		DeleteUser:     MakeDeleteUserEndpoint(s, logger),
		UpdateUser:     MakeUpdateUserEndpoint(s, logger),
		SoftDeleteUser: MakeSoftDeleteUserEndpoint(s, logger),
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
			State:    true,
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

func MakeUpdateUserEndpoint(s services.UserService, logger logrus.FieldLogger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var req UpdateUserRequest
		var ok bool = false

		if req, ok = request.(UpdateUserRequest); !ok {
			return nil, errors.ErrUnsupported // aqui va el error personalizado de interfaz equivocada
		}

		if err != nil {
			logger.Errorln(err.Error())
			return UpdateUserREsponse{}, err
		}

		user := entities.User{
			ID:       req.ID,
			DNI:      req.DNI,
			Email:    req.Email,
			Name:     req.Name,
			Password: req.Password,
			Address:  req.Address,
			Phone:    req.Phone,
			State:    true,
		}

		serviceUser, err := s.UpdateUser(user)
		if err != nil {
			logger.Errorln("Layer: user_endpoint", "Method: MakeUpdateUserEndpoint", "Error:", err)
			return UpdateUserREsponse{}, err
		}
		logger.Infoln("Updated user with id:%s sucessfully ", serviceUser.ID)
		return UpdateUserREsponse{User: serviceUser}, nil
	}
}

func MakeSoftDeleteUserEndpoint(s services.UserService, logger logrus.FieldLogger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var req SoftDeleteUserRequest
		var ok bool = false

		if req, ok = request.(SoftDeleteUserRequest); !ok {
			return nil, errors.New("")
		}

		erro := s.SoftDeleteUser(req.ID)
		if err != nil {
			return DeleteUserResponse{}, err
		}
		logger.Infoln("Layer: user_endpoint ", "Method: MakeSoftDeleteUserEndpoint ", "Soft Delete user with id:%s sucessfully ", req.ID)
		return DeleteUserResponse{}, erro
	}
}

func MakeDeleteUserEndpoint(s services.UserService, logger logrus.FieldLogger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var req DeleteUserRequest
		var ok bool = false

		if req, ok = request.(DeleteUserRequest); !ok {
			return nil, errors.New("ss")
		}
		erro := s.DeleteUser(req.ID)
		if err != nil {
			return DeleteUserResponse{}, err
		}
		logger.Infoln("Layer: user_endpoint ", "Method: MakeDeleteUserEndpoint ", "Delete user with id:%s sucessfully ", req.ID)
		return DeleteUserResponse{}, erro

	}
}
