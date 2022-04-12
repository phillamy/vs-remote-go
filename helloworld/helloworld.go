package helloworld

import (
	"github.com/gin-gonic/gin"
)

func AddHelloWorld(r *gin.Engine) {
	r.GET("/HelloWorld", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!!",
		})
	})
}
