package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kiridharan/controller"
)

func Routes(r *gin.Engine) {

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pings",
	// 	})
	// })

	r.POST("/signup", controller.SignUp)

}
