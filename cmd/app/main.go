package main

import (
	"github.com/Minhajxdd/Ephemr/config"
	"github.com/Minhajxdd/Ephemr/internal/app"
	"github.com/Minhajxdd/Ephemr/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.Init()
	config.ConnectDB()

	router := gin.New()
	router.Use(gin.Logger())

	c := app.NewContainer()

	api := router.Group("/api/v1")
	routes.Setup(api, c)

	router.Run(":" + config.Cfg.Port)

}
