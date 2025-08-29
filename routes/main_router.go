package routes

import "github.com/gin-gonic/gin"

func MainRouter(router *gin.RouterGroup) {
	UserRoutes(router)
}
