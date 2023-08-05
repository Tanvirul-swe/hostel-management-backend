package model

import "gorm.io/gorm"

type RoomImage struct {
	gorm.Model
	RoomId int16  `form:"room_id" json:"room_id"`
	Image  string `form:"image" json:"image"`
}
