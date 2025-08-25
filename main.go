package main

import (
	"log"
	"os"

	"github.com/Minhajxdd/Ephemr/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env: ", err)
	}

	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "8000"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRouter(router)

	router.Run(":" + PORT)

}
