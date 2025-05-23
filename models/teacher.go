package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100);not null" json:"name"`
	Age    int    `gorm:"not null" json:"age"`
	Gender string `gorm:"type:varchar(10)" json:"gender"`

	Phone string `gorm:"type:varchar(15);not null;unique" json:"phone"` // string is preferred for phone numbers
	Email string `gorm:"type:varchar(100);unique;not null" json:"email"`
}
