package model

import "gorm.io/gorm"

type HostelInfo struct {
	gorm.Model
	Name            string            `form:"name" json:"name"`
	HostelFacilites []HostelFacilites `gorm:"foreignKey:HostelId" json:"feature"`
	Rooms           []Rooms           `gorm:"foreignKey:HostelId" json:"room"`
	Reviews         []Reviews         `gorm:"foreignKey:HostelId" json:"review"`
}
