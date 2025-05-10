package repository

import (
	"context"
	"doctor-on-demand/models"

	"gorm.io/gorm"
)

type IDoctorRepository interface {
	CreateDoctor(ctx context.Context, doctor *models.DoctorList) error
	GetByID(ctx context.Context, id uint) (models.DoctorList, error)
	UpdateDoctor(ctx context.Context, id uint, doctor *models.DoctorList) error
	DeleteDoctor(ctx context.Context, id uint) error
	GetAll(ctx context.Context) ([]models.DoctorList, error)
	Count(ctx context.Context) (int64, error)
}

type DoctorRepository struct {
	db *gorm.DB
}

func NewDoctorRepository(db *gorm.DB) *DoctorRepository {
	return &DoctorRepository{db: db}
}

func (r *DoctorRepository) CreateDoctor(ctx context.Context, doctor *models.DoctorList) error {
	return r.db.WithContext(ctx).Create(doctor).Error
}

func (r *DoctorRepository) GetByID(ctx context.Context, id uint) (models.DoctorList, error) {
	var doctor models.DoctorList
	err := r.db.WithContext(ctx).First(&doctor, id).Error
	return doctor, err

}

func (r *DoctorRepository) UpdateDoctor(ctx context.Context, id uint, doctor *models.DoctorList) error {
	doctor.ID = id // GORM needs ID for Save
	return r.db.WithContext(ctx).Save(doctor).Error
}

func (r *DoctorRepository) DeleteDoctor(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.DoctorList{}, id).Error
}

func (r *DoctorRepository) GetAll(ctx context.Context) ([]models.DoctorList, error) {
	var doctors []models.DoctorList
	err := r.db.WithContext(ctx).Find(&doctors).Error
	return doctors, err
}

func (r *DoctorRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	// result := r.db.RowsAffected
	// fmt.Println(result)
	// return result, nil
	err := r.db.WithContext(ctx).Model(&models.DoctorList{}).Count(&count).Error
	return count, err
}
