package middleware

import (
	"fmt"
	"net/http"
	"os"
	"wereserve/models"
	"wereserve/repository"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)


func AuthMiddleware(userRepo *repository.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context){
		//Get token From cookies
		tokenString, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorize",
			})
		return
		}

		//Parse Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _,ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("SECRET")), nil
		})
		
		if err != nil || !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized,  gin.H{
				"error" : "Unauthorize",
			})
			return
		}


		// Ambil claims token dari token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorize",
			})
			return
		}

		// Ambil Id 
		UserID := uint(claims["sub"].(float64))

		// Cari UserID berdasarkan ID
		var user models.User
		result := userRepo.DB.First(&user, UserID)
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error" : "Unauthorized",
			})
			return
		}

		// Ambil data yang udah diambil di DB kirim ke context
		c.Set("user", user)
		c.Next()
	}
}