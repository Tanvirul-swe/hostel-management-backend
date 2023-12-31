package controllers

import (
	"example.com/main/services"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	// Call the service to get the posts
	services.CreateUser(c)

}

func UpdateUserInfo(c *gin.Context) {
	// Call the service to get the posts
	services.UpdateUserInfo(c)

}

func UserLogin(c *gin.Context) {
	// Call the service to get the posts
	services.UserLogin(c)

}

func Validate(c *gin.Context) {
	// Call the service to get the posts
	services.Validate(c)
}

// Get user profile details by user id
func GetUserProfile(c *gin.Context) {
	// Call the service to get the posts
	services.GetUserProfile(c)
}

// Get User Booked Room Details
func GetBookedRoomByUserId(c *gin.Context) {
	// Call the service to get the posts
	services.GetBookedRoomByUsersId(c)
}
