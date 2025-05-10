package models

import (
	"gorm.io/gorm"
)

type Medicine struct {
	gorm.Model
	Name     string  `gorm:"type:varchar(100);not null" json:"name"`
	Type     string  `gorm:"type:varchar(100);not null" json:"type"`
	Price    float64 `gorm:"type:decimal(10,2);not null" json:"price"`
	Stock    bool    `gorm:"type:boolean;not null" json:"stock"`
	Image    string  `gorm:"type:varchar(255)" json:"image"`
	Quantity int     `gorm:"type:int;not null" json:"quantity"`
}
