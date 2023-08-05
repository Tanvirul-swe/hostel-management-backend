package services

import (
	"net/http"

	"example.com/main/constants"
	"example.com/main/database"
	"example.com/main/model"
	"github.com/gin-gonic/gin"
)

// ********Hostel Start*******
func CreateHostel(c *gin.Context) {
	var hostel model.HostelInfo
	// Bind the PostForm data to the User model
	if err := c.ShouldBind(&hostel); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if title and body fields are not empty

	if hostel.Name == "" {
		c.JSON(400, gin.H{
			"message":     "Hostel  name field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	result := database.DB.Create(&hostel)

	// Return result as JSON response with status code 400 if there is an error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Faill to create hostel",
			"status_code": http.StatusInternalServerError,
		})
		return
	}
	//Return response as JSON with status code 201
	c.JSON(http.StatusCreated, gin.H{
		"message":     constants.CreateSuccessfully,
		"status_code": http.StatusCreated,
	})
}

// Get All hostel list
func GetAllHostelInfo(c *gin.Context) {
	var hostelList []model.HostelInfo
	//Get all the hostel list
	result := database.DB.Model(model.HostelInfo{}).Preload("HostelFacilites").Preload("Rooms").Find(&hostelList)
	// Return result as JSON response with status code 400 if there is an error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Fetching Failed",
			"status_code": http.StatusInternalServerError,
		})
		return
	}
	//Return response as JSON with status code 200
	c.JSON(http.StatusOK, gin.H{
		"message":     constants.Successfully,
		"status_code": 200,
		"data":        hostelList,
	})
}

// Create Hostel Facilites
func CreateHostelFacilites(c *gin.Context) {
	var facilites model.HostelFacilites
	// Bind the PostForm data to the Hostel Facilites model
	if err := c.ShouldBind(&facilites); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if title and body fields are not empty

	if facilites.FeatureName == "" {
		c.JSON(400, gin.H{
			"message":     "Feature  field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	result := database.DB.Create(&facilites)

	// Return result as JSON response with status code 400 if there is an error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Faill to create Facilites",
			"status_code": http.StatusInternalServerError,
		})
		return
	}

	//Return response as JSON with status code 201
	c.JSON(http.StatusCreated, gin.H{
		"message":     constants.CreateSuccessfully,
		"status_code": http.StatusCreated,
	})
}

// Get All hostel  feature list
func GetAllHostelFeatureInfo(c *gin.Context) {
	var facilites []model.HostelFacilites
	//Get all the hostel list
	result := database.DB.Find(&facilites)
	// Return result as JSON response with status code 400 if there is an error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Fetching Failed",
			"status_code": http.StatusInternalServerError,
		})
		return
	}
	//Return response as JSON with status code 200
	c.JSON(http.StatusOK, gin.H{
		"message":     constants.Successfully,
		"status_code": 200,
		"data":        facilites,
	})
}

//********Hostel End *********

//**********Room Start*******

func CreateHostelRoom(c *gin.Context) {
	var room model.Rooms
	// Bind the PostForm data to the User model
	if err := c.ShouldBind(&room); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if title and body fields are not empty

	if room.RoomNo == "" {
		c.JSON(400, gin.H{
			"message":     "Room name field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}
	if room.RoomSize == 0 {
		c.JSON(400, gin.H{
			"message":     "Room size field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	if room.TotalSit == 0 {
		c.JSON(400, gin.H{
			"message":     "Total site field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}
	if room.Roomtype == "" {
		c.JSON(400, gin.H{
			"message":     "Room Type field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	if room.Discription == "" {
		c.JSON(400, gin.H{
			"message":     "Discription field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	if room.HostelId == 0 {
		c.JSON(400, gin.H{
			"message":     "Discription field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}
	// var hostel model.HostelInfo
	// hasHostel := database.DB.Find(&hostel,room.HostelId);
	// if room.HostelId ==0 {
	// 	c.JSON(400, gin.H{
	// 		"message":     "Discription field is required",
	// 		"status_code": http.StatusBadRequest,
	// 	})
	// 	return

	// }

	if room.Latitude == 0 {
		c.JSON(400, gin.H{
			"message":     "Latitude field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}
	if room.Longitude == 0 {
		c.JSON(400, gin.H{
			"message":     "Longitude field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	result := database.DB.Create(&room)

	// Return result as JSON response with status code 400 if there is an error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Faill to create room",
			"status_code": http.StatusInternalServerError,
		})
		return
	}
	//Return response as JSON with status code 201
	c.JSON(http.StatusCreated, gin.H{
		"message":     constants.CreateSuccessfully,
		"status_code": http.StatusCreated,
	})
}

// Get All Room details list
func GetSingleRoomInfo(c *gin.Context) {
	id := c.Param("id")
	var roomInfo model.HostelInfo
	//Get all the hostel list
	result := database.DB.Model(model.HostelInfo{}).Preload("HostelFacilites").Preload("Rooms").Where("ID = ?", id).Find(&roomInfo)

	// Return result as JSON response with status code 400 if there is an error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Fetching Failed",
			"status_code": http.StatusInternalServerError,
		})
		return
	}
	if len(roomInfo.Rooms) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message":     "Room id not found",
			"status_code": 200,
		})
		return

	}
	//Return response as JSON with status code 200
	c.JSON(http.StatusOK, gin.H{
		"message":     constants.Successfully,
		"status_code": 200,
		"data":        roomInfo,
	})
}

//********Room End*******
