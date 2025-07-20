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

	// Если пользователь неправильно заполнил форму регистрации, то в теле ответа будет ошибка
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Некорректные данные"})

		return
	}

	result := database.DB.Db.Where("login = ?", user.Login).Find(&model.User{}) // Поиск пользователя с таким же логином

	if result.RowsAffected > 0 {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Пользователь с таким логином уже существует"})

		return
	}

	// Проверка корректности пароля
	err = user.IsValidPassword()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: err.Error()})

		return
	}

	database.DB.Db.Save(&user)

	c.IndentedJSON(http.StatusOK, user)
}
