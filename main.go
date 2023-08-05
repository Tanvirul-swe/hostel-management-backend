package main

import (
	"example.com/main/constants"
	controllers "example.com/main/controller"
	"example.com/main/database"
	"example.com/main/middleware"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	// database.DB.AutoMigrate(&model.User{})
	// database.DB.AutoMigrate(&model.User{})
	r := gin.Default()
	r.POST("/"+constants.ApiVersion+"/create-user", controllers.CreateUser)
	r.PUT("/"+constants.ApiVersion+"/update-user/:Id", controllers.UpdateUserInfo)
	r.POST("/"+constants.ApiVersion+"/user-login", controllers.UserLogin)
	r.GET("/"+constants.ApiVersion+"/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/"+constants.ApiVersion+"/user-profile/:id", controllers.GetUserProfile)

	r.Run()
}
