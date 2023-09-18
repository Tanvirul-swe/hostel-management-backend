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
	// database.DB.AutoMigrate(&model.HostelFacilites{})
	// database.DB.AutoMigrate(&model.HostelInfo{})
	// database.DB.AutoMigrate(&model.RoomImage{})
	// database.DB.AutoMigrate(&model.Rooms{})
	// database.DB.AutoMigrate(&model.ReservedRooms{})
	// database.DB.AutoMigrate(&model.PaymentHistory{})
	// database.DB.AutoMigrate(&model.Reviews{})

	r := gin.Default()
	r.POST("/"+constants.ApiVersion+"/create-review", controllers.CreateHostelReview)
	r.POST("/"+constants.ApiVersion+"/create-room", controllers.CreateRoom)
	r.GET("/"+constants.ApiVersion+"/rooms/:id", controllers.GetSingleRoom)
	r.POST("/"+constants.ApiVersion+"/create-hostel-feature", controllers.CreateHostelFacilites)
	r.GET("/"+constants.ApiVersion+"/hostel-feature", controllers.GetHostelFacilites)
	r.POST("/"+constants.ApiVersion+"/create-hostel", controllers.CreateHostel)
	r.GET("/"+constants.ApiVersion+"/hostel-info", controllers.GetAllHostelInfo)
	r.POST("/"+constants.ApiVersion+"/create-user", controllers.CreateUser)
	r.PUT("/"+constants.ApiVersion+"/update-user/:id", controllers.UpdateUserInfo)
	r.POST("/"+constants.ApiVersion+"/user-login", controllers.UserLogin)
	r.GET("/"+constants.ApiVersion+"/validate", middleware.RequireAuth, controllers.Validate)
	r.GET("/"+constants.ApiVersion+"/user-profile/:id", controllers.GetUserProfile)
	r.POST("/"+constants.ApiVersion+"/reserve-room", controllers.ReservedRoom)
	r.POST("/"+constants.ApiVersion+"/create-payment", controllers.CreatePayment)
	r.GET("/"+constants.ApiVersion+"/payment-history/:id", controllers.GetPaymentHistory)
	r.POST("/"+constants.ApiVersion+"/update-feature/:id", controllers.UpdateHostelFacilites)
	r.POST("/"+constants.ApiVersion+"/delete-feature/:id", controllers.DeleteHostelFacilites)
	r.GET("/"+constants.ApiVersion+"/hostel-rooms/:id", controllers.GetAllRoomsByHostelId)
	r.GET("/"+constants.ApiVersion+"/booked-room/:id", controllers.GetBookedRoomByUserId)

	r.Run()
}
