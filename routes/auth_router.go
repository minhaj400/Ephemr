package routes

import (
	"github.com/Minhajxdd/Ephemr/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("auth/signup", controllers.Signup())
}
