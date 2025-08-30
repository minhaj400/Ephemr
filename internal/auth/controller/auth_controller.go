package controller

import (
	"fmt"
	"net/http"

	"github.com/Minhajxdd/Ephemr/internal/auth/dto"
	services "github.com/Minhajxdd/Ephemr/internal/auth/service"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	SignUp(ctx *gin.Context)
}

type authController struct {
	authService services.AuthService
}

func NewAuthController(a services.AuthService) AuthController {
	return &authController{authService: a}
}

func (c *authController) SignUp(ctx *gin.Context) {
	var body dto.SignUpRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser, token, err := c.authService.SignUp(&body)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.Header("Authorization", fmt.Sprintf("Bearer %s", token))

	ctx.JSON(http.StatusCreated, gin.H{
		"data": newUser,
	})
}
