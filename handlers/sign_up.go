package handlers

import (
	"crypto/aes"
	"log"
	"net/http"
	"vk_test_marketplace/database"
	"vk_test_marketplace/model"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user model.User
	err := c.BindJSON(&user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "string"})

		return
	}

	result := database.DB.Db.Where("login = ?", user.Login).Find(&model.User{})

	if result.RowsAffected > 0 {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "Пользователь с таким логином уже существует"})

		return
	}

	block, err := aes.NewCipher([]byte("secret_key_12345"))

	if err != nil {
		log.Println(err)

		c.IndentedJSON(http.StatusInternalServerError, model.ErrorResponse{Error: "string"})

		return
	}

	unencryptedPassword := user.Password
	encryptedPassword := make([]byte, len([]byte(user.Password)))

	block.Encrypt(encryptedPassword, []byte(user.Password))

	user.Password = string(encryptedPassword)

	database.DB.Db.Save(&user)

	user.Password = unencryptedPassword
	c.IndentedJSON(http.StatusOK, user)
}
