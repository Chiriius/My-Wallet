// api/endpoints/main.go
package endpoints

import (
	"context"
	"my_wallet/api/entities"
	"my_wallet/api/services"
	infraestructure_services "my_wallet/api/services/healtcheck"

	"github.com/go-kit/kit/endpoint"
	"github.com/sirupsen/logrus"
)

// LoginUserRequest represents the request to log in to the application
// @Description Valid user email and password
type LoginUserRequest struct {
	// @example "user@gmail.com"
	Email string `json:"email"` // User's email
	// @example "passwordExample"
	Password string `json:"password"` // User's password
}

// LoginUserResponse represents the response when logged in
// @Description Response when logged in successfully
type LoginUserResponse struct {
	StateLogin bool   `json:"Match,omitempty"` // Login state
	Token      string `json:"token,omitempty"` // Authentication token
	Err        string `json:"error,omitempty"` // Error message, if any
}

// CreateUserRequest represents the request to create a user
// @Description Data needed to create a user
type CreateUserRequest struct {
	// @example 1002842747
	DNI int `json:"dni"` // User's DNI
	// @example "CC"
	TypeDNI string `json:"type_dni"` // Type of DNI
	// @example "User"
	Name string `json:"name"` // User's name
	// @example "user@gmail.com"
	Email string `json:"email"` // User's email
	// @example "passwordExample"
	Password string `json:"password"` // User's password
	// @example "Cra 00 #00-00"
	Address string `json:"address"` // User's address
	// @example 30179423800
	Phone int `json:"phone"` // User's phone number
}

// CreateUserResponse represents the response when creating a user
// @Description Response when a new user is created
type CreateUserResponse struct {
	ID    string `json:"id,omitempty"`    // User ID
	Token string `json:"token,omitempty"` // Authentication token
	Err   string `json:"error,omitempty"` // Error message, if any
}

// UpdateUserRequest represents the request to update a user
// @Description Request to update an existing user
type UpdateUserRequest struct {
	ID      string `json:"id"`       // User ID
	TypeDNI string `json:"type_dni"` // Type of DNI
	// @example 1002842747
	DNI int `json:"dni"` // User's DNI
	// @example "John Doe"
	Name string `json:"name"` // User's name
	// @example "john@example.com"
	Email string `json:"email"` // User's email
	// @example "securepassword123"
	Password string `json:"password"` // New password for the user
	// @example "123 Main St, City"
	Address string `json:"address"` // User's address
	// @example 1234567890
	Phone int `json:"phone"` // User's phone number
}

// UpdateUserResponse represents the response when updating a user
// @Description Response when a user is updated
type UpdateUserREsponse struct {
	User entities.User `json:"user"`            // Updated user
	Err  string        `json:"error,omitempty"` // Error message, if any
}

// GetUserRequest represents the request to get a user by ID
// @Description Request to retrieve an existing user
type GetUserRequest struct {
	ID string `json:"id"` // User ID
}

// GetUserResponse represents the response when retrieving a user
// @Description Response when a user is retrieved
type GetUserResponse struct {
	User entities.User `json:"user"`            // Retrieved user
	Err  string        `json:"error,omitempty"` // Error message, if any
}

// DeleteUserRequest represents the request to delete a user
// @Description Request to delete an existing user
type DeleteUserRequest struct {
	ID string `json:"id"` // User ID
}

// DeleteUserResponse represents the response when deleting a user
// @Description Response when a user is deleted
type DeleteUserResponse struct {
	Err string `json:"error,omitempty"` // Error message, if any
}

// SoftDeleteUserRequest represents the request to soft delete a user
// @Description Request to logically delete an existing user
type SoftDeleteUserRequest struct {
	ID string `json:"id"` // User ID
}

// SoftDeleteUserResponse represents the response when soft deleting a user
// @Description Response when a user is soft deleted
type SoftDeleteUserResponse struct {
	Err string `json:"error,omitempty"` // Error message, if any
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
			logger.Errorln("Layer:user_endpoint", "Method:MakeLoginEndpoint", ErrInterfaceWrong)
			return LoginUserResponse{}, ErrInterfaceWrong
		}
		logger.Infoln(req.Email, " ", req.Password)
		state, user, err := s.Login(ctx, req.Email, req.Password)

		if err != nil {
			logger.Errorln("Layer:user_endpoint", "Method:MakeLoginEndpoint", err)
			return LoginUserResponse{}, ErrInvalidCredentials
		}
		return LoginUserResponse{StateLogin: state, Token: user.Token}, nil
	}
}

// @Summary Create User
// @Description Creates a new user
// @Accept json
// @Produce json
// @Param user body CreateUserRequest true "User"
// @Success 201 {object} CreateUserResponse
// @Failure 400 {object} ErrorResponse
// @Router /user [post]
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

// MakeGetUserEndpoint makes a Get User endpoint.
// @Summary Get User
// @Description Retrieve user information by ID
// @Param id path string true "User ID"
// @Success 200 {object} GetUserResponse
// @Failure 404 {object} ErrorResponse
// @Router /user/{id} [get]
func MakeGetUserEndpoint(s services.UserService, logger logrus.FieldLogger) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		var req GetUserRequest
		var ok bool = false

		if req, ok = request.(GetUserRequest); !ok {
			logger.Errorln("Layer:user_endpoint", "Method:MakeGetUserEndpoint", ErrInterfaceWrong)
			return GetUserResponse{}, ErrInterfaceWrong
		}
		user, err := s.GetUSer(ctx, req.ID)
		if err != nil {
			logger.Errorln("Layer:user_endpoint", "Method:MakeGetUserEndpoint", err)
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
			logger.Errorln("Layer:user_endpoint", "Method:MakeUpdateUserEndpoint", ErrInterfaceWrong)
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
			logger.Errorln("Layer: user_endpoint", "Method: MakeUpdateUserEndpoint", "Error:", err)
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
