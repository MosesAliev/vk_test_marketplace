package router

import (
	"vk_test_marketplace/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/sign_up", handlers.SignUp)

	r.POST("/sign_in", handlers.SignIn)

	r.POST("/post_ad", handlers.PostAd)

	r.GET("/get_ads", handlers.GetAds)

	return r
}
