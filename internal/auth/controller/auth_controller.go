package controller

import (
	"github.com/Minhajxdd/Ephemr/internal/auth/dto"
	services "github.com/Minhajxdd/Ephemr/internal/auth/service"
	"github.com/Minhajxdd/Ephemr/internal/config"
	"github.com/Minhajxdd/Ephemr/pkg/errs"
	"github.com/Minhajxdd/Ephemr/pkg/response"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	SignUp(ctx *gin.Context)
	ConfirmEmail(ctx *gin.Context)
	RefreshToken(ctx *gin.Context)
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

	response.Success(ctx, "Confirmed Email Successfully", nil)
}

func (c *authController) Login(ctx *gin.Context) {

}
