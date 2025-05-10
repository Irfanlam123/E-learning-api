package models

import "gorm.io/gorm"

type Prescription struct {
	gorm.Model
	PatientID   uint         `gorm:"not null" json:"patient_id"`
	DoctorID    uint         `gorm:"not null" json:"doctor_id"`
	PDFPath     string       `gorm:"type:text" json:"pdf_path"` // Add this field
	PatientName string       `gorm:"not null type:VARCHAR(30)" json:"patient_name"`
	DoctorName  string       `gorm:"not null type:VARCHAR(30)" json:"doctor_name"`
	Medications []Medication `gorm:"foreignKey:PrescriptionID" json:"medications"`
	Doctor      DoctorList   `gorm:"foreignKey:DoctorID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"doctor"`
	Patient     Patient      `gorm:"foreignKey:PatientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"patient"`
}

type Medication struct {
	gorm.Model
	PrescriptionID      uint   `gorm:"not null" json:"prescription_id"`
	Name                string `gorm:"not null" json:"name"`      // "Tab. Amlodipine 5mg"
	Dosage              string `gorm:"not null" json:"dosage"`    // "5mg"
	Form                string `gorm:"not null" json:"form"`      // "Tablet", "Syrup", etc.
	Frequency           string `gorm:"not null" json:"frequency"` // "OD" (once daily), "BD" (twice daily)
	Duration            string `gorm:"not null" json:"duration"`  // "30 days"
	SpecialInstructions string `gorm:"type:text" json:"special_instructions"`
	BeforeAfterFood     string `gorm:"type:varchar(10)" json:"before_after_food"` // "before", "after", "empty"
}
