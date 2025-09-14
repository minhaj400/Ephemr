package service

import repositories "github.com/Minhajxdd/Ephemr/internal/user/repository"

// UserService defined method for business logic of users
type UserService interface{}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(r repositories.UserRepository) UserService {
	return &userService{repo: r}
}
