package models

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type DateOnly struct {
	time.Time
}

const dateLayout = "2006-01-02"

func (d *DateOnly) UnmarshalJSON(b []byte) error {
	s := string(b)
	s = s[1 : len(s)-1] // strip quotes
	t, err := time.Parse(dateLayout, s)
	if err != nil {
		return fmt.Errorf("invalid date format: use YYYY-MM-DD")
	}
	d.Time = t
	return nil
}

func (d DateOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.Time.Format(dateLayout))
}

const (
	StatusPending   = "pending"
	StatusConfirmed = "confirmed"
	StatusCancelled = "cancelled"
	StatusCompleted = "completed"
)

type Appointment struct {
	gorm.Model
	PatientID       uint      `gorm:"not null" json:"patient_id"`
	DoctorID        uint      `gorm:"not null" json:"doctor_id"`
	AppointmentDate time.Time `gorm:"type:date;not null" json:"appointment_date"`
	Status          string    `gorm:"type:varchar(20);not null;check:status IN ('pending','confirmed','cancelled','completed')" json:"status"`

	ScheduleID uint `gorm:"not null" json:"schedule_id"`

	// Relationships
	Schedule DoctorSchedule `gorm:"foreignKey:ScheduleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"schedule"`
	Doctor   DoctorList     `gorm:"foreignKey:DoctorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"doctor"`
	Patient  Patient        `gorm:"foreignKey:PatientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"patient"`
}
