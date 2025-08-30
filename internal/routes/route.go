package routes

import (
	"github.com/Minhajxdd/Ephemr/internal/app"
	"github.com/Minhajxdd/Ephemr/internal/user"
	"github.com/gin-gonic/gin"
)

func Setup(api *gin.RouterGroup, c *app.Container) {
	user.RegisterUserRoutes(api, c.UserModule.Controller)
}
