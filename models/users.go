package models

import (
	"gorm.io/gorm"
)

// User represents the structure of the user data for signup
type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(100);not null" json:"name"`
	Email    string `gorm:"type:varchar(100);unique;not null" json:"email"`
	Mobile   string `gorm:"type:varchar(15);unique;not null" json:"mobile"`
	Address  string `gorm:"type:varchar(255);not null" json:"address"`
	Password string `gorm:"type:varchar(255);not null" json:"password"`
	Image    string `gorm:"type:text" json:"image"`
	Role     string `gorm:"type:varchar(255)" json:"role"`
}

// TableName sets the table name for the User model
