package services

import (
	"net/http"

	"example.com/main/constants"
	"example.com/main/database"
	"example.com/main/model"
	"github.com/gin-gonic/gin"
)

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
	result := database.DB.Model(model.HostelInfo{}).Preload("HostelFacilites").Find(&hostelList)
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
