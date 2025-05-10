package repository

import (
	"context"
	"doctor-on-demand/models"

	"gorm.io/gorm"
)

type IMedincineRepository interface {
	GetById(ctx context.Context, id uint) (models.Medicine, error)
	Create(ctx context.Context, Medicine *models.Medicine) error
	Update(ctx context.Context, id uint, Medicine *models.Medicine) error
	Delete(ctx context.Context, id uint) error
	GetAll(ctx context.Context) ([]models.Medicine, error)
	Count(ctx context.Context) error
}

type MedicineRepository struct {
	db *gorm.DB
}

func NewMedicineRepository(db *gorm.DB) *MedicineRepository {
	return &MedicineRepository{db: db}
}

func (r *MedicineRepository) GetById(ctx context.Context, id uint) (models.Medicine, error) {
	var medcine models.Medicine
	err := r.db.WithContext(ctx).First(&medcine, id).Error
	return medcine, err
}

func (r *MedicineRepository) Create(ctx context.Context, Medicine *models.Medicine) error {
	return r.db.WithContext(ctx).Create(Medicine).Error
}

func (r *MedicineRepository) Update(ctx context.Context, id uint, medicine *models.Medicine) error {
	medicine.ID = id
	return r.db.WithContext(ctx).Save(medicine).Error
}

func (r *MedicineRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Patient{}, id).Error
}
func (r *MedicineRepository) GetAll(ctx context.Context) ([]models.Medicine, error) {
	var medicine []models.Medicine
	err := r.db.WithContext(ctx).Find(&medicine).Error
	return medicine, err
}

func (r *MedicineRepository) Count(ctx context.Context) error {
	var count int64
	return r.db.WithContext(ctx).Model(&models.Medicine{}).Count(&count).Error

}
