package services

import (
	"net/http"

	"example.com/main/constants"
	"example.com/main/database"
	"example.com/main/model"
	"github.com/gin-gonic/gin"
)

// ********Payment Start*******
func CreatePayment(c *gin.Context) {
	var payment model.PaymentHistory
	// Bind the PostForm data to the User model
	if err := c.ShouldBind(&payment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Check if title and body fields are not empty

	if payment.TransectionId == "" {
		c.JSON(400, gin.H{
			"message":     "Transecction Id field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	if payment.UserId == 0 {
		c.JSON(400, gin.H{
			"message":     "User Id field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}
	if payment.Amount == 0 {
		c.JSON(400, gin.H{
			"message":     "Amount  field is required",
			"status_code": http.StatusBadRequest,
		})
		return

	}

	result := database.DB.Create(&payment)

	// Return result as JSON response with status code 400 if there is an error
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message":     "Faill to create payment",
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

//*********End Payment*********

// Get All payment History by user id
func GetPaymentHistoryByUserId(c *gin.Context) {
	var history []model.PaymentHistory
	id := c.Param("id")
	result := database.DB.Where("user_id = ?", id).Find(&history).Where("UserId = ?", id)
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
		"data":        history,
	})
}
