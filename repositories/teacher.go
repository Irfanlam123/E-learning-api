package repository

import (
	"context"
	"doctor-on-demand/models"

	"gorm.io/gorm"
)

type ITeacherRepository interface {
	CreateTeacher(ctx context.Context, teacher *models.Teacher) error
	GetByID(ctx context.Context, id uint) (models.Teacher, error)
	UpdateTeacher(ctx context.Context, id uint, teacher *models.Teacher) error
	DeleteTeacher(ctx context.Context, id uint) error
	GetAll(ctx context.Context) ([]models.Teacher, error)
	Count(ctx context.Context) (int64, error)
}

type TeacherRepository struct {
	db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) *TeacherRepository {
	return &TeacherRepository{db: db}
}

func (r *TeacherRepository) CreateTeacher(ctx context.Context, teacher *models.Teacher) error {
	return r.db.WithContext(ctx).Create(teacher).Error
}

func (r *TeacherRepository) GetByID(ctx context.Context, id uint) (models.Teacher, error) {
	var doctor models.Teacher
	err := r.db.WithContext(ctx).First(&doctor, id).Error
	return doctor, err

}

func (r *TeacherRepository) UpdateTeacher(ctx context.Context, id uint, teacher *models.Teacher) error {
	teacher.ID = id // GORM needs ID for Save
	return r.db.WithContext(ctx).Save(teacher).Error
}

func (r *TeacherRepository) DeleteTeacher(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Teacher{}, id).Error
}

func (r *TeacherRepository) GetAll(ctx context.Context) ([]models.Teacher, error) {
	var doctors []models.Teacher
	err := r.db.WithContext(ctx).Find(&doctors).Error
	return doctors, err
}

func (r *TeacherRepository) Count(ctx context.Context) (int64, error) {
	var count int64
	// result := r.db.RowsAffected
	// fmt.Println(result)
	// return result, nil
	err := r.db.WithContext(ctx).Model(&models.Teacher{}).Count(&count).Error
	return count, err
}
