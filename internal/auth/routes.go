package auth

import (
	"github.com/Minhajxdd/Ephemr/internal/auth/controller"
	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.RouterGroup, c controller.AuthController) {
	auth := r.Group("/auth")
	{
		auth.POST("/signup", c.SignUp)
		auth.GET("/confirm-email/:userId/:token", c.ConfirmEmail)
	}
}
