package models

import "gorm.io/gorm"

type Articles struct {
	gorm.Model
	Title string `gorm:"type:varchar(100);not null" json:"title"`
	Desc  string `gorm:"type:varchar(200);not null" json:"desc"`
	Image string `gorm:"type:text" json:"image"`
	Pin   bool   `gorm:"default:false" json:"pin"`
}
