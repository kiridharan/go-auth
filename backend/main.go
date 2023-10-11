package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kiridharan/controller"
	"github.com/kiridharan/initializers"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.DatabaseInit()
	initializers.AutoSync()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pings",
		})
	})

	r.POST("/signup", controller.SignUp)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port + "")
}
