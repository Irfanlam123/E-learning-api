package handlers

import (
	"doctor-on-demand/config"
	"doctor-on-demand/models"
	repository "doctor-on-demand/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type IAppointmentHandler interface {
	Create() echo.HandlerFunc
	GetAppointmentsByDoctorID() echo.HandlerFunc
}

type AppointmentHandler struct {
	repo repository.IAppointmentRepository
}

func NewAppointmentHandler(repo repository.IAppointmentRepository) *AppointmentHandler {
	return &AppointmentHandler{
		repo: repo,
	}
}
func (h *AppointmentHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req models.Appointment
		id, _ := config.GenerateId()
		req.ID = id
		// Bind JSON payload to struct
		if err := c.Bind(&req); err != nil {
			logrus.WithField("error", err).Error("Failed to bind appointment")
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Invalid request payload",
			})
		}

		// Validate required fields
		if req.DoctorID == 0 || req.PatientID == 0 {
			logrus.Error("Doctor ID and Patient ID cannot be zero")
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "Doctor ID and Patient ID are required",
			})
		}

		if req.AppointmentDate.IsZero() {
			logrus.Error("Appointment date is missing or invalid")
			return c.JSON(http.StatusBadRequest, echo.Map{
				"error": "appointment_date must be in YYYY-MM-DD format",
			})
		}

		// Save appointment
		result, err := h.repo.BookAppointment(c.Request().Context(), req)
		if err != nil {
			logrus.WithField("error", err).Error("Failed to book appointment")

			if err == repository.ErrScheduleNotAvailable {
				return c.JSON(http.StatusConflict, echo.Map{
					"error": "Selected time slot is not available",
				})
			}

			return c.JSON(http.StatusInternalServerError, echo.Map{
				"error": "Could not book appointment",
			})
		}

		return c.JSON(http.StatusOK, result)
	}
}

func (h *AppointmentHandler) GetAppointmentsByDoctorID() echo.HandlerFunc {
	return func(c echo.Context) error {
		doctorID, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			logrus.WithField("error", err).Error("Failed to parse doctor ID")
		}
		if doctorID == 0 || doctorID == -1 {
			logrus.Error("Doctor  cannot be zero")
		}
		result, err := h.repo.GetAppointmentsByDoctorID(c.Request().Context(), uint(doctorID))
		if err != nil {
			logrus.WithField("error", err).Error(http.StatusBadRequest)
		}

		return c.JSON(http.StatusOK, result)
	}

}
