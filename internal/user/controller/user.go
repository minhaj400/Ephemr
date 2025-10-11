package controller

import "github.com/Minhajxdd/Ephemr/internal/user/service"

// UserController defines the methods for handling user related queries.
type UserController interface{}

type userController struct {
	srv service.UserService
}

func NewUserService(srv service.UserService) UserController {
	return &userController{srv: srv}
}
