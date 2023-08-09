package model

import "gorm.io/gorm"

type Reviews struct {
	gorm.Model
	UserId   int32   `form:"user_id" json:"user_id"`
	HostelId int32   `form:"hostel_id" json:"hostel_id"`
	Rate     float32 `form:"rate" json:"rate"`
	Commant  string  `form:"commant" json:"commant"`
	User     User    `gorm:"foreignKey:UserId" json:"user"`
}
