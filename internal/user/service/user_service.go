package service

import repositories "github.com/Minhajxdd/Ephemr/internal/user/repository"

type UserService interface{}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return &userService{repo: r}
}
