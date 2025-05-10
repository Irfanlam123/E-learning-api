package repository

import (
	"context"
	"doctor-on-demand/models"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// DoctorScheduleRepository defines the interface
type IDoctorScheduleRepository interface {
	Create(ctx context.Context, schedule models.DoctorSchedule) (models.DoctorSchedule, error)
	GetByID(ctx context.Context, id uint) (models.DoctorSchedule, error)
	Update(ctx context.Context, id uint, schedule models.DoctorSchedule) (models.DoctorSchedule, error)
	Delete(ctx context.Context, id uint) error
	GetByDoctorID(ctx context.Context, doctorID uint) ([]models.DoctorSchedule, error)
}

// doctorScheduleRepo implements the interface
type DoctorScheduleRepo struct {
	db *gorm.DB
}

// NewDoctorScheduleRepository creates a new repository instance
func NewDoctorScheduleRepository(db *gorm.DB) IDoctorScheduleRepository {
	return &DoctorScheduleRepo{db: db}
}

func (r *DoctorScheduleRepo) Create(ctx context.Context, schedule models.DoctorSchedule) (models.DoctorSchedule, error) {

	// Verify doctor exists
	var doctor models.DoctorList
	if err := r.db.WithContext(ctx).First(&doctor, schedule.DoctorID).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"doctor_id": schedule.DoctorID,
			"error":     err,
		}).Error("Doctor not found")
		return models.DoctorSchedule{}, fmt.Errorf("doctor with ID %d not found", schedule.DoctorID)
	}

	// Create the schedule
	if err := r.db.WithContext(ctx).Create(&schedule).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error":    err.Error(),
			"schedule": schedule,
		}).Error("Failed to create doctor schedule")
		return models.DoctorSchedule{}, fmt.Errorf("failed to create schedule: %w", err)
	}

	// Reload with associated data
	if err := r.db.WithContext(ctx).
		Preload("DoctorList").
		First(&schedule, schedule.ID).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"schedule_id": schedule.ID,
			"error":       err,
		}).Error("Failed to load schedule with doctor data")
		return models.DoctorSchedule{}, fmt.Errorf("failed to load created schedule: %w", err)
	}

	return schedule, nil
}

// Implement other interface methods similarly
func (r *DoctorScheduleRepo) GetByID(ctx context.Context, id uint) (models.DoctorSchedule, error) {
	var schedule models.DoctorSchedule
	err := r.db.WithContext(ctx).
		Preload("Doctor").
		First(&schedule, id).Error
	if err != nil {
		return models.DoctorSchedule{}, fmt.Errorf("failed to get schedule: %w", err)
	}
	return schedule, nil
}

func (r *DoctorScheduleRepo) Update(ctx context.Context, id uint, schedule models.DoctorSchedule) (models.DoctorSchedule, error) {
	// 1. First check if schedule exists
	var existingSchedule models.DoctorSchedule
	if err := r.db.WithContext(ctx).First(&existingSchedule, id).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
			"id":    id,
		}).Error("Schedule not found")
		return models.DoctorSchedule{}, fmt.Errorf("schedule not found")
	}

	// 2. Set the ID from parameter to ensure we update the correct record
	schedule.ID = id

	// 4. Perform the update
	if err := r.db.WithContext(ctx).Model(&models.DoctorSchedule{}).
		Where("id = ?", id).
		Updates(&schedule).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
			"id":    id,
		}).Error("Failed to update schedule")
		return models.DoctorSchedule{}, fmt.Errorf("failed to update schedule")
	}

	// 5. Fetch the updated record with relationships
	if err := r.db.WithContext(ctx).Preload("Doctor").First(&schedule, id).Error; err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err.Error(),
			"id":    id,
		}).Error("Failed to fetch updated schedule")
		return models.DoctorSchedule{}, fmt.Errorf("failed to fetch updated schedule")
	}

	return schedule, nil
}

func (r *DoctorScheduleRepo) Delete(ctx context.Context, id uint) error {
	// Implementation here
	return nil
}

func (r *DoctorScheduleRepo) GetByDoctorID(ctx context.Context, doctorID uint) ([]models.DoctorSchedule, error) {
	// Implementation here
	return []models.DoctorSchedule{}, nil
}
