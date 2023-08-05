package model

import "gorm.io/gorm"

type HostelInfo struct {
	gorm.Model
	Name            string            `form:"name" json:"name"`
	HostelFacilites []HostelFacilites `gorm:"foreignKey:HostelId"`
}
