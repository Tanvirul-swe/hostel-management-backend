package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `form:"name" json:"name"`
	Email       string `form:"email" json:"email"`
	Password    string `form:"password" json:"password"`
	Image       string `form:"image" json:"image"`
	StudentId   string `form:"student_id" json:"student_id"`
	Department  string `form:"department" json:"department"`
	Phone       string `form:"phone" json:"phone"`
	DateOfBirth string `form:"date_of_birth" json:"date_of_birth"`
	Gender      string `form:"gender" json:"gender"`
}
