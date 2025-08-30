package routes

import (

	"github.com/Minhajxdd/Ephemr/config"
	"github.com/Minhajxdd/Ephemr/controllers"
	"github.com/Minhajxdd/Ephemr/repositories"
	"github.com/Minhajxdd/Ephemr/services"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.RouterGroup) {
	userController := controllers.NewUserController()
	auth := router.Group("/auth")
	{
		auth.POST("/signup", userController.Signup)
	}
}
