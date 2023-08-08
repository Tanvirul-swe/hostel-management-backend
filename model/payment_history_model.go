package model

import "gorm.io/gorm"

type PaymentHistory struct {
	gorm.Model
	TransectionId string  `form:"transection_id" json:"transection_id"`
	UserId        int32   `form:"user_id" json:"user_id"`
	Amount        float64 `form:"amount" json:"amount"`
}
