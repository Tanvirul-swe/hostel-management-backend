package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"example.com/main/constants"
	"example.com/main/database"
	"example.com/main/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(c *gin.Context) {
	// Get token from request header
	tokenString,err := c.Cookie("Authorization")
	fmt.Println(tokenString)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	// Parse token string to jwt.Token struct
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	// Return error if there is an error
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
	// Check if token is valid
	if token.Valid {
		// Get claims from token
		claims := token.Claims.(jwt.MapClaims)
		// Get user id from claims
		userId := claims["user_id"].(string)

		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Find user by id in database using GORM ORM and store result in user variable
		var user model.User
		result := database.DB.Where("id = ?", userId).First(&user)
		// Return result as JSON response with status code 400 if there is an error
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message":     constants.StatusInternalServerError,
				"status_code": http.StatusInternalServerError,
			})
			return
		}

		// Set user variable in context
		c.Set("user", user)
		c.Next()

	}
	// Return error if token is not valid
	c.JSON(http.StatusUnauthorized, gin.H{
		"message":     "Unauthorized Access - Token is not valid",
		"status_code": http.StatusUnauthorized,
	})
}
