package handlers

import (
	"net/http"
	"vk_test_marketplace/database"
	"vk_test_marketplace/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func PostAd(c *gin.Context) {
	signedToken := c.GetHeader("token")

	secret := []byte("secret_key")

	jwtToken, err := jwt.Parse(signedToken, func(t *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "некорректный токен"})

		return
	}

	if !jwtToken.Valid {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "некорректный токен"})

		return
	}

	var ad model.Ad
	err = c.BindJSON(&ad)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "некорректные данные"})

		return
	}

	database.DB.Db.Save(&ad)
}
