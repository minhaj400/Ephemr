package services

import (
	"errors"

	"github.com/Minhajxdd/Ephemr/dto"
	"github.com/Minhajxdd/Ephemr/models"
	"github.com/Minhajxdd/Ephemr/repositories"
	"github.com/Minhajxdd/Ephemr/utils"
)

type UserService interface {
	CreateUser(user dto.Signup) (*models.User, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return &userService{userRepo: r}
}

func (s *userService) CreateUser(user dto.Signup) (*models.User, error) {
	existingUser, err := s.userRepo.FindByEmail(*user.Email)

	if err != nil {
		return nil, nil
	}
	if existingUser != nil {
		return nil, errors.New("user with email already exists")
	}

	hashedPassword, err := utils.HashPassword(*user.Password)
	if err != nil {
		return nil, errors.New("invalid password")
	}

	newUser := &models.User{
		Name:     *user.First_name + *user.Last_name,
		Email:    *user.Email,
		Password: hashedPassword,
	}

	err = s.userRepo.Create(newUser)

	if err != nil {
		return nil, errors.New("something went wrong")
	}

	newUser.Password = ""

	return newUser, nil
}
