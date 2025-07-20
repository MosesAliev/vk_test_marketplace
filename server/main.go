package main

import (
	"vk_test_marketplace/database"
	"vk_test_marketplace/router"
)

func main() {
	database.ConnectDB()

	r := router.SetupRouter()

	r.Run()
}
