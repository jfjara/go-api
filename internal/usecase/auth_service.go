package usecase

import (
	"errors"

	"github.com/juanfran/mi-api/internal/domain"
	"github.com/juanfran/mi-api/internal/infrastructure/http/dto"
	"github.com/juanfran/mi-api/internal/infrastructure/logger"
	"github.com/juanfran/mi-api/internal/infrastructure/validation"
)

const USER_NOT_FOUND_MESSAGE = "user not found"
const USER_PASSWORD_INCORRECT_MESSAGE = "user or password incorrect"

type AuthService struct {
	userRepository     domain.UserRepository
	hasher             domain.PasswordHasher
	securityRepository domain.SecurityRepository
}

func NewAuthService(repo domain.UserRepository, hasher domain.PasswordHasher, securityRepository domain.SecurityRepository) *AuthService {
	return &AuthService{userRepository: repo, hasher: hasher, securityRepository: securityRepository}
}

func (s *AuthService) Login(userInfo dto.LoginRequest) (dto.LoginResponse, error) {
	logger.Log.Debug("Start user login", "userInfo", userInfo)

	if err := validation.Validate.Struct(userInfo); err != nil {
		return *dto.NewRegisterResponseError(err.Error()), errors.New("user information not valid")
	}

	user, err := s.userRepository.GetByUsername(userInfo.Username)
	if err != nil {
		logger.Log.Error(USER_NOT_FOUND_MESSAGE)
		return *dto.NewRegisterResponseError(err.Error()), errors.New(USER_NOT_FOUND_MESSAGE)
	}

	if !s.hasher.Compare(user.Password, userInfo.Password) {
		logger.Log.Error(USER_PASSWORD_INCORRECT_MESSAGE)
		return *dto.NewRegisterResponseError(USER_PASSWORD_INCORRECT_MESSAGE), errors.New(USER_PASSWORD_INCORRECT_MESSAGE)
	}

	tokenString, err := s.securityRepository.CreateToken(user)
	if err != nil {
		logger.Log.Error("User token not generated correctly")
		return *dto.NewRegisterResponseError(err.Error()), err
	}

	logger.Log.Debug("User logged correctly", "user", userInfo.Username)
	return *dto.NewRegisterResponse(tokenString), nil
}
