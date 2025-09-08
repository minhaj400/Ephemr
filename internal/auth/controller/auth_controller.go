package controller

import (
	"github.com/Minhajxdd/Ephemr/internal/auth/dto"
	services "github.com/Minhajxdd/Ephemr/internal/auth/service"
	"github.com/Minhajxdd/Ephemr/pkg/errs"
	"github.com/Minhajxdd/Ephemr/pkg/response"
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
		response.HandleError(ctx, errs.FromValidation(err))
		return
	}

	newUser, err := c.authService.SignUp(&body)

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	response.Success(ctx, "Verification Mail Sent", newUser)
}
