package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `form:"name" json:"name"`
	Email      string `form:"email" json:"email"`
	Password   string `form:"password" json:"password"`
	Image      string `form:"image" json:"image"`
	StudentId  string `form:"student_id" json:"student_id"`
	Department string `form:"department" json:"department"`
}
