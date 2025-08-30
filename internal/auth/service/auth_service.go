package service

import (
	"errors"
	"strconv"

	"github.com/Minhajxdd/Ephemr/internal/auth/dto"
	"github.com/Minhajxdd/Ephemr/internal/user/model"
	repositories "github.com/Minhajxdd/Ephemr/internal/user/repository"
	"github.com/Minhajxdd/Ephemr/pkg/crypto"
	"github.com/Minhajxdd/Ephemr/pkg/jwt"
)

type AuthService interface {
	SignUp(user *dto.SignUpRequest) (*model.User, string, error)
}

type authService struct {
	userRepo   repositories.UserRepository
	hasher     crypto.PasswordHasher
	jwtManager jwt.TokenManager
}

func NewAuthService(r repositories.UserRepository, h crypto.PasswordHasher, j jwt.TokenManager) AuthService {
	return &authService{userRepo: r, hasher: h, jwtManager: j}
}

func (s *authService) SignUp(user *dto.SignUpRequest) (*model.User, string, error) {
	existingUser, err := s.userRepo.FindByEmail(user.Email)

	if err != nil {
		return nil, "", nil
	}
	if existingUser != nil {
		return nil, "", errors.New("user with email already exists")
	}

	hashedPassword, err := s.hasher.Hash(user.Password)
	if err != nil {
		return nil, "", errors.New("invalid password")
	}

	newUser := &model.User{
		Name:     user.FullName,
		Email:    user.Email,
		Password: hashedPassword,
	}

	err = s.userRepo.Create(newUser)

	if err != nil {
		return nil, "", errors.New("something went wrong")
	}

	newUser.Password = ""

	claims := &jwt.Claims{
		UserID: strconv.FormatUint(uint64(newUser.ID), 10),
		Role:   "user",
	}

	token, err := s.jwtManager.Generate(*claims)

	if err != nil {
		return nil, "", errors.New("somethign went wrong")
	}

	return newUser, token, nil
}
