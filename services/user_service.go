package services

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"example.com/main/constants"
	"example.com/main/database"
	"example.com/main/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
	var user model.User
	// Bind the PostForm data to the User model
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if title and body fields are not empty

	if user.Email == "" {
		c.JSON(400, gin.H{
			"message":     "Email Field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}
	if user.Password == "" {
		c.JSON(400, gin.H{
			"message":     "Password Field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	hasUser := database.DB.Where("email = ?", user.Email).First(&user)
	if hasUser.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email Already Exists",
		})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     constants.UserCreateFailedMessage,
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	user.Password = string(hash)
	result := database.DB.Create(&user)

	// Return result as JSON response with status code 400 if there is an error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     constants.UserCreateFailedMessage,
			"status_code": http.StatusInternalServerError,
		})
		return
	}
	//Return response as JSON with status code 201
	c.JSON(http.StatusCreated, gin.H{
		"message":     constants.UserCreateSuccessMessage,
		"status_code": http.StatusCreated,
	})
}

// UserLogin function to login user and return token if user exists in database

func UserLogin(c *gin.Context) {
	var user model.User
	// Bind the PostForm data to the User model
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	password := user.Password
	// Check email and password fields are not empty
	// If empty return status code 400
	// And return message as JSON response
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "Email Field is required",
			"status_code": http.StatusBadRequest,
		})
		return
	}

	if user.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "Password Field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	// Find user by email and password in database using GORM ORM and store result in user variable
	result := database.DB.Where("email = ?", user.Email).First(&user)

	// Return result as JSON response with status code 400 if there is an error
	if result.Error != nil && result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message":     constants.UserNotFoundMessage,
			"status_code": http.StatusNotFound,
		})
		return
	}

	// Compare password from request with password from database using bcrypt package and store result in err variable
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	fmt.Println(err)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":     "Invalid Credentials",
			"status_code": http.StatusUnauthorized,
		})
		return
	}

	// Create claims with multiple fields populated
	claims := MyCustomClaims{
		"bar",
		jwt.RegisteredClaims{
			// A usual scenario is to set the expiration time relative to the current time
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "auth-tanvir",
			Subject:   "acces token",
			ID:        "1",
			Audience:  []string{"somebody_else"},
		},
		user.Email,
		strconv.Itoa(int(user.ID)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	userToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	fmt.Printf("User Token: %v\n", userToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message":     "Unauthorized User",
			"status_code": http.StatusUnauthorized,
		})
	}
	// Set cookie with token value and set SameSite to None and Secure to true for cross site request and https request respectively
	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("Authorization", userToken, 3600, "/", "localhost", false, true)
	//Return response as JSON with status code 201
	c.JSON(http.StatusOK, gin.H{
		"message":     constants.UserLoginSuccessMessage,
		"status_code": http.StatusOK,
		"user":        user,
		"token":       userToken,
	})
}

type MyCustomClaims struct {
	Foo string `json:"foo"`

	jwt.RegisteredClaims

	// This field is also added in the payload, even though it's not defined in
	// the MyCustomClaims struct
	Email  string `json:"email"`
	UserId string `json:"user_id"`
}

func Validate(c *gin.Context) {

	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"user":    user,
	})

}

// Get user profile function to get user profile from database using GORM ORM
func GetUserProfile(c *gin.Context) {
	var user model.User
	// Get user id from request params
	id := c.Param("id")
	// Find user by id in database using GORM ORM and store result in user variable
	fmt.Println("User ID: ", id)
	result := database.DB.First(&user, id)
	// Return result as JSON response with status code 400 if there is an error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     constants.UserNotFoundMessage,
			"status_code": http.StatusBadRequest,
		})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message":     constants.UserNotFoundMessage,
			"status_code": http.StatusNotFound,
		})
		return
	}

	//Return response as JSON with status code 200
	c.JSON(http.StatusOK, gin.H{
		"message":     "User Profile",
		"status_code": http.StatusOK,
		"user":        user,
	})
}

func UpdateUserInfo(c *gin.Context) {

	//Get post id from request
	userId := c.Param("Id")

	//Get model values from request
	var user model.User
	//Bind JSON body to model

	c.BindJSON(&user)

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":     "AuthorName Field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	//Update Post
	result := database.DB.Model(&user).Where("id = ?", userId).Updates(&user).Preload("Categorys").First(&user, userId)

	// Return result as JSON response with status code 400 if there is an error or post not found in database
	// RowsAffected is 0 if no record found

	if result.Error != nil && result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status_code": http.StatusNotFound,
			"message":     "User Id Not Found",
		})
		return
	}
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": result.Error.Error(),
		})
		return
	}

	//Return response as JSON with status code 200
	c.JSON(http.StatusAccepted, gin.H{
		"message":     "Updated Successfully",
		"status_code": http.StatusAccepted,
		"user":        user,
	})

}
