package model

import "gorm.io/gorm"

type ReservedRooms struct {
	gorm.Model
	RoomId int32 `form:"room_id" json:"room_id"`
	UserId int32 `form:"user_id" json:"user_id"`
}
