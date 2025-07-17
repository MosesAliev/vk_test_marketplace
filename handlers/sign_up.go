package handlers

import (
	"net/http"
	"vk_test_marketplace/database"
	"vk_test_marketplace/model"

	"github.com/gin-gonic/gin"
)

// Регистрация пользователей
func SignUp(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Некорректные данные"})

		return
	}

	result := database.DB.Db.Where("login = ?", user.Login).Find(&model.User{})

	if result.RowsAffected > 0 {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Пользователь с таким логином уже существует"})

		return
	}

	err = user.IsValidPassword()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})

		return
	}

	database.DB.Db.Save(&user)

	c.IndentedJSON(http.StatusOK, user)
}
