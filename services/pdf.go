package services

import (
	"doctor-on-demand/models"
	"fmt"
	"time"

	"github.com/jung-kurt/gofpdf"
)

type PDFService struct {
	OutputPath string // e.g., "./prescriptions"
}

func NewPDFService(outputPath string) *PDFService {
	return &PDFService{OutputPath: outputPath}
}

func (s *PDFService) GeneratePrescriptionPDF(prescription *models.Prescription) (string, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	//setting font here in pdf
	pdf.SetFont("Arial", "B", 16)

	//set header here in pdf
	pdf.Cell(0, 10, "MEDICAL PRESCRIPTION")
	pdf.Ln(12) //ln is basically a line break

	// Clinic/Hospital Info
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, "City Hospital")
	pdf.Ln(8)
	pdf.Cell(0, 10, "123 Health Street, Medical City")
	pdf.Ln(15)
	// Patient Information

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Patient:")
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, prescription.Patient.Name)
	pdf.Ln(8)

	pdf.Cell(40, 10, "Age/Gender:")
	pdf.Cell(0, 10, fmt.Sprintf("%d/%s", prescription.Patient.Age, prescription.Patient.Gender))
	pdf.Ln(8)

	pdf.Cell(40, 10, "Date:")
	pdf.Cell(0, 10, prescription.CreatedAt.Format("02-Jan-2006"))
	pdf.Ln(15)

	// Diagnosis
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "Diagnosis:")
	pdf.Ln(8)
	pdf.SetFont("Arial", "", 12)
	pdf.MultiCell(0, 8, prescription.Patient.Dignosis, "0", "L", false)
	pdf.Ln(10)

	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(0, 10, "Medications:")
	pdf.Ln(8)
	pdf.SetFont("Courier", "", 12)

	for _, med := range prescription.Medications {
		medLine := fmt.Sprintf("• %s %s - %s x %s", med.Name, med.Dosage, med.Frequency, med.Duration)
		if med.SpecialInstructions != "" {
			medLine += fmt.Sprintf(" (%s)", med.SpecialInstructions)
		}
		pdf.MultiCell(0, 8, medLine, "0", "L", false)
	}
	pdf.Ln(10)

	// Additional Instructions
	// if prescription.Medications.SpecialInstructions != "" {
	// 	pdf.SetFont("Arial", "B", 12)
	// 	pdf.Cell(0, 10, "Instructions:")
	// 	pdf.Ln(8)
	// 	pdf.SetFont("Arial", "", 12)
	// 	pdf.MultiCell(0, 8, prescription.Instructions, "0", "L", false)
	// 	pdf.Ln(10)
	// }

	// Doctor Information
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Doctor:")
	pdf.SetFont("Arial", "", 12)
	pdf.Cell(0, 10, prescription.Doctor.Name)
	pdf.Ln(8)

	// pdf.Cell(40, 10, "License No:")
	// pdf.Cell(0, 10, prescription.Doctor.LicenseNumber)
	// pdf.Ln(15)

	// Signature
	pdf.SetFont("Arial", "I", 12)
	pdf.Cell(0, 10, "Signature:________________________")
	pdf.Ln(8)
	pdf.Cell(0, 10, "Date: "+time.Now().Format("02-Jan-2006"))

	// Generate filename
	filename := fmt.Sprintf("%s/prescription_%d_%s.pdf",
		s.OutputPath,
		prescription.ID,
		time.Now().Format("20060102"))

	// Save file
	err := pdf.OutputFileAndClose(filename)
	if err != nil {
		return "", fmt.Errorf("failed to generate PDF: %v", err)
	}

	return filename, nil

}
