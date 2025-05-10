package repository

import (
	"context"
	"doctor-on-demand/models"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

var (
	ErrScheduleNotAvailable = errors.New("doctor schedule is not available")
	ErrInvalidAppointment   = errors.New("invalid appointment details")
)

type IAppointmentRepository interface {
	BookAppointment(ctx context.Context, appointment models.Appointment) (models.Appointment, error)
	// GetAppointmentByID(ctx context.Context, id string) (models.Appointment, error)
	// GetAllAppointments(ctx context.Context, filter models.AppointmentFilter) ([]models.Appointment, error)
	GetAppointmentsByDoctorID(ctx context.Context, doctorID uint) ([]models.Appointment, error)
	// CancelAppointment(ctx context.Context, appointmentID string) error
	// UpdateAppointmentStatus(ctx context.Context, appointmentID string, status string) error
}

type AppointmentRepository struct {
	db *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) IAppointmentRepository {
	return &AppointmentRepository{
		db: db,
	}
}

// BookAppointment creates a new appointment after checking doctor's availability
func (r *AppointmentRepository) BookAppointment(ctx context.Context, appointment models.Appointment) (models.Appointment, error) {
	tx := r.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 1. Check if the schedule is available
	var schedule models.DoctorSchedule
	if err := tx.Where("id = ? AND is_available = ?", appointment.ScheduleID, true).First(&schedule).Error; err != nil {
		tx.Rollback()
		// Add more detailed logging
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.WithFields(logrus.Fields{
				"schedule_id": appointment.ScheduleID,
			}).Error("Doctor schedule is not available")
			return models.Appointment{}, ErrScheduleNotAvailable
		}
		return models.Appointment{}, err
	}

	// 2. Mark schedule as booked
	if err := tx.Model(&models.DoctorSchedule{}).
		Where("id = ?", appointment.ScheduleID).
		Update("is_available", false).Error; err != nil {
		tx.Rollback()
		return models.Appointment{}, err
	}

	// 3. Create the appointment
	appointment.Status = models.StatusConfirmed // or "booked"

	if err := tx.Create(&appointment).Error; err != nil {
		tx.Rollback()
		return models.Appointment{}, err
	}

	// 4. Commit transaction
	if err := tx.Commit().Error; err != nil {
		return models.Appointment{}, err
	}
	// 5. Reload appointment with related fields (after commit, use new context)
	if err := r.db.WithContext(ctx).
		Preload("Doctor").
		Preload("Patient").
		Preload("Schedule").
		First(&appointment, appointment.ID).Error; err != nil {
		logrus.WithField("appointment_id", appointment.ID).Error("Failed to preload appointment relations")
		return appointment, nil // Still return appointment without relationships
	}

	return appointment, nil
}

func (r *AppointmentRepository) GetAppointmentsByDoctorID(ctx context.Context, doctorID uint) ([]models.Appointment, error) {
	var appointments []models.Appointment

	err := r.db.WithContext(ctx).Where("doctor_id = ?", doctorID).Preload("Doctor").Preload("Patient").Preload("Schedule").Find(&appointments).Error
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"doctor_id": doctorID,
			"error":     err,
		}).Error("failed to get appointments by doctor id")
		return nil, err
	}

	return appointments, nil
}
