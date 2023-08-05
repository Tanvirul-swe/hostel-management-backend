package model

import "gorm.io/gorm"

type HostelFacilites struct {
	gorm.Model
	HostelId    int16  `form:"hostel_id" json:"hostel_id"`
	FeatureName string `form:"feature_name" json:"feature_name"`
}
