package handlers

import (
	"log"
	"net/http"
	"strings"
	"vk_test_marketplace/database"
	"vk_test_marketplace/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Размещение нового объявления
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

	if len(ad.Title) > 50 {
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "слишком большая длина заголовка"})

		return
	}

	parts := strings.Split(ad.Image, ".")

	if len(parts) == 0 {
		log.Println(len(parts))

		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "недопустимый формат изображения"})

		return
	}

	if parts[len(parts)-1] != "png" && parts[len(parts)-1] != "PNG" && parts[len(parts)-1] != "jpg" && parts[len(parts)-1] != "JPG" &&
		parts[len(parts)-1] != "jpeg" && parts[len(parts)-1] != "JPEG" {
		log.Println(parts[len(parts)-1])
		c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "недопустимый формат изображения"})

		return
	}

	claims, _ := jwtToken.Claims.(jwt.MapClaims)

	ad.UserLogin = claims["login"].(string)

	database.DB.Db.Save(&ad)

	c.IndentedJSON(http.StatusOK, ad)
}
