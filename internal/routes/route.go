package routes

import (
	"github.com/Minhajxdd/Ephemr/internal/app"
	"github.com/Minhajxdd/Ephemr/internal/auth"
	"github.com/Minhajxdd/Ephemr/internal/user"
	"github.com/gin-gonic/gin"
)

func Setup(api *gin.RouterGroup, c *app.Container) {
	auth.RegisterAuthRoutes(api, *c.AuthModule.AuthController)
	user.RegisterUserRoutes(api, c.UserModule.UserController)
}
