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

	routes.AuthRouter(router)

	router.Run(":" + config.Cfg.Port)

}
