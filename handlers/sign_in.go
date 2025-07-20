package handlers

import (
	"net/http"
	"vk_test_marketplace/database"
	"vk_test_marketplace/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Авторизация пользователей
func SignIn(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)

	// Если пользователь неправильно заполнил форму авторизации, то в теле ответа будет ошибка
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "некорректные данные"})

		return
	}

	result := database.DB.Db.Where("login = ? AND password = ?", user.Login, user.Password).Find(&model.User{}) // Поиск пользователя в базе данных

	if result.RowsAffected == 0 {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Неправильный логин или пароль"})

		return
	}

	// Payload токена
	claims := jwt.MapClaims{
		"login": user.Login,
	}

	secret := []byte("secret_key") // Секретный ключ для токена

	jwtToken := jwt.New(jwt.SigningMethodHS256) // Генерация jwt-токена

	jwtToken.Claims = claims                        // Подписание токена
	signedToken, _ := jwtToken.SignedString(secret) //

	c.Writer.Header().Set("token", signedToken)
}
