package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kiridharan/initializers"
	"github.com/kiridharan/routes"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.DatabaseInit()
	initializers.AutoSync()
}

func main() {
	env := os.Getenv("ENV")
	log.Println("ENV: ", env)
	port := os.Getenv("PORT")
	r := gin.Default()
	routes.Routes(r)
	if port == "" || env == "development" {
		port = "8080"
	}
	r.Run(":" + port + "")
}
