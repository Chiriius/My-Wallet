package services

import (
	"context"
	"fmt"
	"my_wallet/api/entities"
	repository_user "my_wallet/api/respository/user"
	"my_wallet/api/utils"
	"my_wallet/api/utils/jwt"
	"regexp"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type UserService interface {
	CreateUser(ctx context.Context, use entities.User) (entities.User, error)
	GetUSer(ctx context.Context, id string) (entities.User, error)
	DeleteUser(ctx context.Context, id string) error
	SoftDeleteUser(ctx context.Context, id string) error
	UpdateUser(ctx context.Context, user entities.User) (entities.User, error)
	Login(ctx context.Context, email string, password string) (bool, entities.User, error)
}

type userService struct {
	ctx        context.Context
	repository repository_user.UserRepository
	logger     logrus.FieldLogger
	validate   *validator.Validate
}

func NewUserService(repo repository_user.UserRepository, logger logrus.FieldLogger, ctx context.Context) *userService {
	return &userService{
		ctx:        ctx,
		repository: repo,
		logger:     logger,
		validate:   validator.New(),
	}
}

func (s *userService) CreateUser(ctx context.Context, user entities.User) (entities.User, error) {

	if err := s.validate.Struct(user); err != nil {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error:", err)
		fmt.Println("error:", err)
		return entities.User{}, ErrValidation
	}
	phoneStr := fmt.Sprintf("%d", user.Phone)
	if len(user.Password) < 8 {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error:", ErrLenghtPassword)
		return entities.User{}, ErrLenghtPassword

	}
	if len(phoneStr) != 10 {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error:", ErrLenghPhone)
		return entities.User{}, ErrLenghPhone

	}
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	if !re.MatchString(user.Name) {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error:", ErrNameSpecialCharacters)
		return entities.User{}, ErrNameSpecialCharacters
	}
	if user.TypeDNI != "CC" && user.TypeDNI != "NIT" {
		s.logger.Errorln("Layer: user_services", "Method: UpdateUser", "Error:", ErrTypeDNI)
		return entities.User{}, ErrTypeDNI
	}

	passwordHashed, err := utils.HashPassword(user.Password)
	if err != nil {
		s.logger.Errorln("Layer: user_services", "Method: CreateUser", "Error:", ErrHashingPassword)
		return entities.User{}, ErrHashingPassword
	}
	user.Password = passwordHashed
	s.logger.Info("Layer: user_services", "Method: CreateUser", "User:", user)

	user.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.Update_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	token, refreshToken, _ := jwt.GenerateToken(user.Email, s.logger)
	user.Token = token
	user.RefreshToken = refreshToken

	return s.repository.CreateUser(user, ctx)

}

func (s *userService) GetUSer(ctx context.Context, id string) (entities.User, error) {

	return s.repository.GetUser(id, ctx)
}

func (s *userService) UpdateUser(ctx context.Context, user entities.User) (entities.User, error) {
	if err := s.validate.Struct(user); err != nil {
		s.logger.Errorln("Layer: user_services", "Method: UpdateUser", "Error:", err)
		return entities.User{}, err
	}
	phoneStr := fmt.Sprintf("%d", user.Phone)
	if len(user.Password) < 8 {
		s.logger.Errorln("Layer: user_services", "Method: UpdateUser", "Error:", ErrLenghtPassword)
		return entities.User{}, ErrLenghtPassword

	}
	if len(phoneStr) != 10 {
		s.logger.Errorln("Layer: user_services", "Method: UpdateUser", "Error:", ErrLenghPhone)
		return entities.User{}, ErrLenghPhone

	}
	if user.TypeDNI != "CC" && user.TypeDNI != "NIT" {
		s.logger.Errorln("Layer: user_services", "Method: UpdateUser", "Error:", ErrTypeDNI)
		return entities.User{}, ErrTypeDNI
	}
	re := regexp.MustCompile(`^[a-zA-Z\s]+$`)
	if !re.MatchString(user.Name) {
		s.logger.Errorln("Layer: user_services", "Method: UpdateUser", "Error:", ErrNameSpecialCharacters)
		return entities.User{}, ErrNameSpecialCharacters
	}
	passwordHashed, err := utils.HashPassword(user.Password)
	if err != nil {

		s.logger.Errorln("Layer: user_services", "Method: UpdateUser", "Error:", ErrHashingPassword)
		return entities.User{}, ErrHashingPassword
	}
	user.Password = passwordHashed
	return s.repository.UpdateUser(user, ctx)
}

func (s *userService) SoftDeleteUser(ctx context.Context, id string) error {
	return s.repository.SoftDeleteUser(id, ctx)
}

func (s *userService) DeleteUser(ctx context.Context, id string) error {

	return s.repository.DeleteUser(id, ctx)
}

func (s *userService) Login(ctx context.Context, email string, password string) (bool, entities.User, error) {
	user, _ := s.repository.GetUserByEmail(email, ctx)
	s.logger.Infoln(user)
	loginState := true

	if utils.CheckPasswordHash(password, user.Password) != true {

		s.logger.Errorln("Layer: user_services", "Method: Login", "Error:", ErrInvalidCredentials)

		return false, entities.User{}, ErrInvalidCredentials

	}

	token, refreshToken, _ := jwt.GenerateToken(user.Email, s.logger)
	user.Token = token
	user.RefreshToken = refreshToken

	userr, _ := s.repository.UpdateUserToken(user, ctx)

	user.Created_at = userr.Created_at
	user.Update_at = userr.Update_at

	return loginState, user, nil
}
