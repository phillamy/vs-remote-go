package router

import (
	"base/helloworld"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	helloworld.AddHelloWorld(r)

	return r
}
