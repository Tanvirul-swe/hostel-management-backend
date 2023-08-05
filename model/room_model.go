package model

import "gorm.io/gorm"

type Rooms struct {
	gorm.Model
	RoomNo      string      `form:"room_no" json:"room_no"`
	RoomSize    float32     `form:"room_size" json:"room_size"`
	TotalSit    int16       `form:"total_sit" json:"total_sit"`
	Roomtype    string      `form:"room_type" json:"room_type"`
	Discription string      `form:"discription" json:"discription"`
	HostelId    int16       `form:"hostel_id" json:"hostel_id"`
	Latitude    float32     `form:"latitude" json:"latitude"`
	Longitude   float32     `form:"longitude" json:"longitude"`
	RoomImage   []RoomImage `gorm:"foreignKey:RoomId"`
}
