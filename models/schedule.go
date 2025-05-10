package models

import (
	"time"

	"gorm.io/gorm"
)

type DoctorSchedule struct {
	gorm.Model
	DoctorID    uint      `gorm:"not null" json:"doctor_id"`
	DayOfWeek   string    `gorm:"type:varchar(10);not null" json:"day"`
	IsAvailable bool      `gorm:"not null" json:"is_available"`
	Date        time.Time `gorm:"type:date;not null" json:"date"`        // Date without time
	StartTime   time.Time `gorm:"type:time; not null" json:"start_time"` // Time in HH:MM:SS format
	EndTime     time.Time `gorm:"type:time;not null" json:"end_time"`    // Time in HH:MM:SS format

	Doctor DoctorList `gorm:"foreignKey:DoctorID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"doctor"`
}
