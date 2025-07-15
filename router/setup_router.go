package router

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	var r = gin.Default()
	return r
}
