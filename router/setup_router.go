package router

import (
	"vk_test_marketplace/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/sign_up", handlers.SignUp)

	return r
}
