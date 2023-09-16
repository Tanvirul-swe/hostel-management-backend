package controllers

import (
	"example.com/main/services"
	"github.com/gin-gonic/gin"
)

// ************Hostel Section**********
func CreateHostel(c *gin.Context) {
	// Create Hostel call create function in service class
	services.CreateHostel(c)

}

// Hostel review create
func CreateHostelReview(c *gin.Context) {
	// Create review call create function in service class
	services.CreateHostelReview(c)

}

func GetAllHostelInfo(c *gin.Context) {
	// Create Hostel call create function in service class
	services.GetAllHostelInfo(c)

}

func CreateHostelFacilites(c *gin.Context) {
	// Create Hostel facilites call create function in service class
	services.CreateHostelFacilites(c)

}

// Update Hostel facilites call create function in service class
func UpdateHostelFacilites(c *gin.Context) {
	services.UpdateHostelFacilites(c)
}

// Delete Hostel facilites call create function in service class
func DeleteHostelFacilites(c *gin.Context) {
	services.DeleteHostelFacilites(c)
}

func GetHostelFacilites(c *gin.Context) {
	// Create Hostel facilites call create function in service class
	services.GetAllHostelFeatureInfo(c)

}

//**********End Hostel Section*******

//******Start Hostel Room Section*****

func CreateRoom(c *gin.Context) {
	services.CreateHostelRoom(c)

}

func GetSingleRoom(c *gin.Context) {
	services.GetSingleRoomInfo(c)

}

// Get All Rooms By Hostel Id
func GetAllRoomsByHostelId(c *gin.Context) {
	services.GetAllRoomsById(c)

}

//********End Hostel Room Section*********

// ***********Reserved Room*******
func ReservedRoom(c *gin.Context) {
	services.ReservedRoom(c)

}
