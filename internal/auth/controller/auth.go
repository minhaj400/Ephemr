package controller

import (
	"github.com/Minhajxdd/Ephemr/internal/auth/dto"
	services "github.com/Minhajxdd/Ephemr/internal/auth/service"
	"github.com/Minhajxdd/Ephemr/internal/config"
	"github.com/Minhajxdd/Ephemr/pkg/errs"
	"github.com/Minhajxdd/Ephemr/pkg/response"
	"github.com/gin-gonic/gin"
)

// AuthController defines the methods for handling user authentication.
type AuthController interface {
	// SignUp handles user registration.
	SignUp(ctx *gin.Context)

	// ConfirmEmail handles the confirmation of a user's email address.
	ConfirmEmail(ctx *gin.Context)

	// RefreshToken handles the renewal of access tokens using a refresh token.
	RefreshToken(ctx *gin.Context)

	// Login handles user authentication.
	Login(ctx *gin.Context)
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

func (c *authController) Login(ctx *gin.Context) {
	var body dto.LoginRequest
	if err := ctx.ShouldBindJSON(&body); err != nil {
		response.HandleError(ctx, errs.FromValidation(err))
		return
	}

	var device = ctx.GetHeader("user-agent")
	var ipAddress = ctx.ClientIP()

	token, refreshToken, err := c.authService.Login(&body, device, ipAddress)

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	ctx.SetCookie("access_token", token, 900, "/", config.Cfg.HostName, false, false)

	ctx.SetCookie("refresh_token", refreshToken, 7*24*3600, "/", config.Cfg.HostName, false, true)

	response.Success(ctx, "Logged In Successfully", nil)
}

func (c *authController) ConfirmEmail(ctx *gin.Context) {
	var params dto.ConfirmEmailRequest
	if err := ctx.ShouldBindUri(&params); err != nil {
		response.HandleError(ctx, errs.FromValidation(err))
		return
	}

	var device = ctx.GetHeader("user-agent")
	var ipAddress = ctx.ClientIP()

	token, refreshToken, err := c.authService.ConfirmEmail(&params, device, ipAddress)

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	ctx.SetCookie("access_token", token, 900, "/", config.Cfg.HostName, false, false)

	ctx.SetCookie("refresh_token", refreshToken, 7*24*3600, "/", config.Cfg.HostName, false, true)

	response.Success(ctx, "Confirmed Email Successfully", nil)
}

func (c *authController) RefreshToken(ctx *gin.Context) {

	var device = ctx.GetHeader("user-agent")
	var ipAddress = ctx.ClientIP()
	var refreshToken, err = ctx.Cookie("refresh_token")

	if err != nil {
		response.HandleError(ctx, errs.InternalError(err))
		return
	}

	token, refreshToken, err := c.authService.RefreshToken(refreshToken, device, ipAddress)

	if err != nil {
		response.HandleError(ctx, err)
		return
	}

	ctx.SetCookie("access_token", token, 900, "/", config.Cfg.HostName, false, false)

	ctx.SetCookie("refresh_token", refreshToken, 7*24*3600, "/", config.Cfg.HostName, false, true)

	response.Success(ctx, "Refreshed Successfully", nil)
}
