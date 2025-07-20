package router

import (
	"vk_test_marketplace/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/sign_up", handlers.SignUp) // регистрация

	r.POST("/sign_in", handlers.SignIn) // авторизация

	r.POST("/post_ad", handlers.PostAd) // публикация объявления

	r.GET("/get_ads", handlers.GetAds) // отображение объявлений

	return r
}
