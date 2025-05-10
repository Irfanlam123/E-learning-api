package repository

import (
	"context"
	"doctor-on-demand/models"
	"doctor-on-demand/services"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type IPrescriptionRepositoy interface {
	Generate(ctx context.Context, prescription *models.Prescription) (*models.Prescription, error)
}

type PrescriptionRepository struct {
	db         *gorm.DB
	pdfService *services.PDFService
}

func NewPrescriptionRepository(db *gorm.DB, pdfService *services.PDFService) *PrescriptionRepository {
	return &PrescriptionRepository{
		db:         db,
		pdfService: pdfService,
	}
}

func (r *PrescriptionRepository) Generate(ctx context.Context, prescription *models.Prescription) (*models.Prescription, error) {

	// First create the prescription record

	err := r.db.WithContext(ctx).Create(prescription).Error
	if err != nil {
		logrus.Info("eror generating odf")
	}

	// Then generate PDF
	pdfPath, err := r.pdfService.GeneratePrescriptionPDF(prescription)
	if err != nil {
		logrus.Info("Eror generating pdf")
	}

	// Update prescription with PDF path
	prescription.PDFPath = pdfPath
	err = r.db.WithContext(ctx).Save(prescription).Error

	return prescription, err
}
