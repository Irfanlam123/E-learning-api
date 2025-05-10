package repository

import (
	"context"
	"doctor-on-demand/models"

	"gorm.io/gorm"
)

type IPatientRepository interface {
	GetById(ctx context.Context, id uint) (models.Patient, error)
	Create(ctx context.Context, patient *models.Patient) error
	Update(ctx context.Context, id uint, patient *models.Patient) error
	Delete(ctx context.Context, id uint) error
	GetAll(ctx context.Context) ([]models.Patient, error)
	Count(ctx context.Context) error
}

type PatientRepository struct {
	db *gorm.DB
}

func NewPatientRepository(db *gorm.DB) *PatientRepository {
	return &PatientRepository{db: db}
}

func (r *PatientRepository) GetById(ctx context.Context, id uint) (models.Patient, error) {
	var patient models.Patient
	err := r.db.WithContext(ctx).First(&patient, id).Error
	return patient, err
}

func (r *PatientRepository) Create(ctx context.Context, patient *models.Patient) error {
	return r.db.WithContext(ctx).Create(patient).Error
}

func (r *PatientRepository) Update(ctx context.Context, id uint, patient *models.Patient) error {
	patient.ID = id
	return r.db.WithContext(ctx).Save(patient).Error
}

func (r *PatientRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Patient{}, id).Error
}
func (r *PatientRepository) GetAll(ctx context.Context) ([]models.Patient, error) {
	var patients []models.Patient
	err := r.db.WithContext(ctx).Find(&patients).Error
	return patients, err
}

func (r *PatientRepository) Count(ctx context.Context) error {
	var count int64
	return r.db.WithContext(ctx).Model(&models.Patient{}).Count(&count).Error

}
