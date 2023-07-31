package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string  `json:"name"`
	Email    string  `json:"email"`
	Password string  `json:"password"`
	Image    string  `json:"image"`
	StudentId   string  `json:"student_id"`
	Department    string  `json:"department"`
}
