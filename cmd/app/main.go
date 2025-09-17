package main

import (
	"github.com/Minhajxdd/Ephemr/internal/app"
	"github.com/Minhajxdd/Ephemr/internal/config"
	"github.com/Minhajxdd/Ephemr/internal/database"
	"github.com/Minhajxdd/Ephemr/internal/middleware"
	"github.com/Minhajxdd/Ephemr/internal/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.Init()
	database.ConnectDB()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.CORS(config.Cfg.Allowed_Origins))

	c := app.NewContainer()

	api := router.Group("/api/v1")
	routes.Setup(api, c)

	router.Run(":" + config.Cfg.Port)

}
