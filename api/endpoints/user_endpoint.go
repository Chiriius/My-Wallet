package endpoints

import (
	"context"
	"my_wallet/api/entities"
	"my_wallet/api/services"
	infraestructure_services "my_wallet/api/services/healtcheck"

	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
)

type LoginUserRequest struct {
	Email    string
	Password string
}

type LoginUserResponse struct {
	StateLogin bool   `json:"Match,omitempty"`
	Token      string `json:"token,omitempty"`
	Err        string `json:"error,omitempty"`
}

type CreateUserRequest struct {
	DNI      int
	TypeDNI  string
	Name     string
	Email    string
	Password string
	Address  string
	Phone    int
}

type CreateUserResponse struct {
	ID    string `json:"id,omitempty"`
	Token string `json:"token,omitempty"`
	Err   string `json:"error,omitempty"`
}

type UpdateUserRequest struct {
	ID       string
	TypeDNI  string
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
	Login          endpoint.Endpoint
	HealthCheck    endpoint.Endpoint
}

func MakeServerEndpoints(s services.UserService, h infraestructure_services.HealtcheckService, logger logrus.FieldLogger) Endpoints {
	return Endpoints{
		CreateUser:     MakeCreateUserEndpoint(s, logger),
		GetUser:        MakeGetUserEndpoint(s, logger),
		DeleteUser:     MakeDeleteUserEndpoint(s, logger),
		UpdateUser:     MakeUpdateUserEndpoint(s, logger),
		SoftDeleteUser: MakeSoftDeleteUserEndpoint(s, logger),
		Login:          MakeLoginEndpoint(s, logger),
		HealthCheck:    MakeGetHealthCheckEndpoint(h, logger),
	}
}

func MakeLoginEndpoint(s services.UserService, logger logrus.FieldLogger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var req LoginUserRequest
		var ok bool = false

		if req, ok = request.(LoginUserRequest); !ok {
			return LoginUserResponse{}, ErrInterfaceWrong
		}
		logger.Infoln(req.Email, " ", req.Password)
		state, user, err := s.Login(ctx, req.Email, req.Password)

		if err != nil {
			return LoginUserResponse{}, ErrInvalidCredentials
		}
		return LoginUserResponse{StateLogin: state, Token: user.Token}, nil
	}
}

func MakeCreateUserEndpoint(s services.UserService, logger logrus.FieldLogger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var req CreateUserRequest
		var ok bool = false
		if req, ok = request.(CreateUserRequest); !ok {
			logger.Errorln("Layer:user_endpoint", "Method:MakeCreateUserEndpoint", "Error: Interface type wrong")
			return CreateUserResponse{}, ErrInterfaceWrong
		}
		user := entities.User{
			DNI:      req.DNI,
			TypeDNI:  req.TypeDNI,
			Name:     req.Name,
			Email:    req.Email,
			Password: req.Password,
			Address:  req.Address,
			Phone:    req.Phone,
			Enabled:  true,
		}
		serviceUser, err := s.CreateUser(ctx, user)
		if err != nil {
			logger.Errorln("Layer:user_endpoint", "Method:MakeCreateUserEndpoint", err)
			return CreateUserResponse{}, err
		}
		logger.Infoln("Layer:user_endpoint", "Method:MakeCreateUserEndpoint", "Response:", CreateUserResponse{ID: serviceUser.ID})
		return CreateUserResponse{ID: serviceUser.ID, Token: serviceUser.Token}, nil

	}
}

func MakeGetUserEndpoint(s services.UserService, logger logrus.FieldLogger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var req GetUserRequest
		var ok bool = false

		if req, ok = request.(GetUserRequest); !ok {
			return GetUserResponse{}, ErrInterfaceWrong
		}
		user, err := s.GetUSer(ctx, req.ID)
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
			return UpdateUserREsponse{}, ErrInterfaceWrong
		}
		user := entities.User{
			ID:       req.ID,
			TypeDNI:  req.TypeDNI,
			DNI:      req.DNI,
			Email:    req.Email,
			Name:     req.Name,
			Password: req.Password,
			Address:  req.Address,
			Phone:    req.Phone,
			Enabled:  true,
		}
		serviceUser, err := s.UpdateUser(ctx, user)
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
			return SoftDeleteUserResponse{}, ErrInterfaceWrong
		}
		erro := s.SoftDeleteUser(ctx, req.ID)
		logger.Infoln("Layer: user_endpoint ", "Method: MakeSoftDeleteUserEndpoint ", "Soft Delete user with id:%s sucessfully ", req.ID)
		return SoftDeleteUserResponse{}, erro
	}
}

func MakeDeleteUserEndpoint(s services.UserService, logger logrus.FieldLogger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var req DeleteUserRequest
		var ok bool = false

		if req, ok = request.(DeleteUserRequest); !ok {
			return DeleteUserResponse{}, ErrInterfaceWrong
		}
		erro := s.DeleteUser(ctx, req.ID)
		logger.Infoln("Layer: user_endpoint ", "Method: MakeDeleteUserEndpoint ", "Delete user with id:%s sucessfully ", req.ID)
		return DeleteUserResponse{}, erro

	}
}
