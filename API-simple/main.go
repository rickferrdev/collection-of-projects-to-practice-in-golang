package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	Router().Run(":8080")
}

func Router() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	return router
}
