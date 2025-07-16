package handlers

import (
	"net/http"
	"vk_test_marketplace/database"
	"vk_test_marketplace/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Отобржение ленты объявлений
func GetAds(c *gin.Context) {
	var Ads []model.Ad
	database.DB.Db.Find(&Ads)

	signedToken := c.GetHeader("token")

	if len(signedToken) > 0 {
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

		claims, _ := jwtToken.Claims.(jwt.MapClaims)

		authorizedUser := claims["login"].(string)

		for i := range Ads {
			if Ads[i].UserLogin == authorizedUser {
				Ads[i].IsYours = true
			}
		}

	}

	c.IndentedJSON(http.StatusOK, Ads)
}
