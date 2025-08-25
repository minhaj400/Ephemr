package controllers

import (
	"net/http"

	"github.com/Minhajxdd/Ephemr/dto"
	"github.com/Minhajxdd/Ephemr/shared"
	"github.com/gin-gonic/gin"
)

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user dto.Signup
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := shared.Validate.Struct(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

	}
}
