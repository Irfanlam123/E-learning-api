package repository

import (
	"context"
	"doctor-on-demand/models"

	"gorm.io/gorm"
)

type ICourseRepository interface {
	Create(ctx context.Context, course *models.Course) error
	GetAll(ctx context.Context) ([]models.Course, error)
}

type CourseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) ICourseRepository {
	return &CourseRepository{
		db: db,
	}
}

func (r *CourseRepository) Create(ctx context.Context, courses *models.Course) error {
	return r.db.WithContext(ctx).Create(courses).Error
}
func (r *CourseRepository) GetAll(ctx context.Context) ([]models.Course, error) {
	var courses []models.Course
	err := r.db.WithContext(ctx).Find(&courses).Error
	return courses, err
}
