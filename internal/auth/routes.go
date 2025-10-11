package auth

import (
	"github.com/Minhajxdd/Ephemr/internal/auth/controller"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup, c controller.AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", c.SignUp)
		auth.POST("/login", c.Login)
		auth.GET("/confirm-email/:userId/:token", c.ConfirmEmail)
		auth.GET("/refresh", c.RefreshToken)
	}
}
