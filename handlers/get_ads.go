package handlers

import (
	"math"
	"net/http"
	"strconv"
	"vk_test_marketplace/database"
	"vk_test_marketplace/model"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// Отобржение ленты объявлений
func GetAds(c *gin.Context) {
	page := 1
	pageParam, ok := c.GetQuery("page")

	if ok {
		var err error
		page, err = strconv.Atoi(pageParam)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "номером страницы должно быть целое число"})

			return
		}

		if page < 1 {
			c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "номер страницы не может быть меньше 1"})

			return
		}
	}

	maxPrice := math.MaxInt
	minPrice := 0
	minPriceStr, ok := c.GetQuery("min_price")

	if ok {
		var err error
		minPrice, err = strconv.Atoi(minPriceStr)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "минимальная цена должна быть целым числом"})

			return
		}
	}

	maxPriceStr, ok := c.GetQuery("max_price")

	if ok {
		var err error
		maxPrice, err = strconv.Atoi(maxPriceStr)

		if err != nil {
			c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "максимальная цена должна быть целым числом"})

			return
		}
	}

	var Ads []model.Ad
	param, ok := c.GetQuery("sort")

	order, _ := c.GetQuery("order")

	if ok {
		switch param {
		case "price":
			if order == "desc" {
				database.DB.Db.Limit(10).Offset((page-1)*10).Order("price desc").Where("price >= ? AND price <= ?", minPrice, maxPrice).Find(&Ads)
			} else {
				database.DB.Db.Limit(10).Offset((page-1)*10).Order("price").Where("price >= ? AND price <= ?", minPrice, maxPrice).Find(&Ads)
			}
		case "date":
			if order == "desc" {
				database.DB.Db.Limit(10).Offset((page-1)*10).Order("created_at desc").Where("price >= ? AND price <= ?", minPrice, maxPrice).Find(&Ads)
			} else {
				database.DB.Db.Limit(10).Offset((page-1)*10).Order("created_at").Where("price >= ? AND price <= ?", minPrice, maxPrice).Find(&Ads)
			}
		default:
			c.IndentedJSON(http.StatusBadRequest, model.ErrorResponse{Error: "отсортировать можно только по дате или цене"})

			return
		}
	} else {
		database.DB.Db.Limit(10).Offset((page-1)*10).Where("price >= ? AND price <= ?", minPrice, maxPrice).Find(&Ads)
	}

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
