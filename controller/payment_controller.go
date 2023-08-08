package controllers

import (
	"example.com/main/services"
	"github.com/gin-gonic/gin"
)

func CreatePayment(c *gin.Context) {
	// Create Payment call create function in service class
	services.CreatePayment(c)

}

func GetPaymentHistory(c *gin.Context) {
	// Get Payment history  function in service class
	services.GetPaymentHistoryByUserId(c)

}
