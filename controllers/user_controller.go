package controllers

import (
	"net/http"

	"github.com/Minhajxdd/Ephemr/dto"
	"github.com/Minhajxdd/Ephemr/services"
	"github.com/Minhajxdd/Ephemr/shared"
	"github.com/gin-gonic/gin"
)

type SignupResponse struct {
	User  UserDTO `json:"user"`
	Token string  `json:"token"`
}

type UserDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type UserController struct {
	userService services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{userService: s}
}

func (ctrl *UserController) Signup(c *gin.Context) {
	var user dto.Signup
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := shared.Validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, token, err := ctrl.userService.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userResponse := UserDTO{
		ID:    string(createdUser.ID),
		Name:  createdUser.Name,
		Email: createdUser.Email,
	}

	c.JSON(http.StatusCreated, SignupResponse{
		User:  userResponse,
		Token: token,
	})
}