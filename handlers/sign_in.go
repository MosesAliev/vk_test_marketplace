package handlers

import (
	"net/http"
	"vk_test_marketplace/database"
	"vk_test_marketplace/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func SignIn(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "некорректные данные"})

		return
	}

	result := database.DB.Db.Where("login = ? AND password = ?", user.Login, user.Password).Find(&model.User{})

	if result.RowsAffected == 0 {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Неправильный логин или пароль"})

		return
	}

	claims := jwt.MapClaims{
		"login": user.Login,
	}

	secret := []byte("secret_key")

	jwtToken := jwt.New(jwt.SigningMethodHS256)

	jwtToken.Claims = claims
	signedToken, _ := jwtToken.SignedString(secret)

	c.Writer.Header().Set("token", signedToken)
}
