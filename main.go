package main

import (
	"github.com/Minhajxdd/Ephemr/config"
	"github.com/Minhajxdd/Ephemr/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.Init()
	config.ConnectDB()

	router := gin.New()
	router.Use(gin.Logger())

	api := router.Group("/api/v1")
	routes.MainRouter(api)

	router.Run(":" + config.Cfg.Port)

}
