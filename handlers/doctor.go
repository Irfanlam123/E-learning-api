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

type IDoctorHandler interface {
	Create() echo.HandlerFunc
	GetById() echo.HandlerFunc
	Delete() echo.HandlerFunc
	Update() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Count() echo.HandlerFunc
}

type DoctorHandler struct {
	repo repository.IDoctorRepository
}

func NewDoctorHandler(repo repository.IDoctorRepository) *DoctorHandler {
	return &DoctorHandler{repo: repo}
}

func (d *DoctorHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := models.DoctorList{}
		id, _ := config.GenerateId()
		req.ID = id
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		err := d.repo.CreateDoctor(c.Request().Context(), &req)
		if err != nil {
			logrus.Error("Failed to create doctor: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create doctor"})
		}

		return c.JSON(http.StatusCreated, req) // Return created doctor with ID
	}
}

func (d *DoctorHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		if idStr == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "id must be provided"})
		}

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id format"})
		}

		doctor, err := d.repo.GetByID(c.Request().Context(), uint(id))
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"id":    id,
				"error": err,
			}).Error("Failed to get doctor")
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get doctor"})
		}

		return c.JSON(http.StatusOK, doctor)
	}
}

func (d *DoctorHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		if idStr == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "id must be provided"})
		}

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id format"})
		}

		req := models.DoctorList{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		err = d.repo.UpdateDoctor(c.Request().Context(), uint(id), &req)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"id":    id,
				"error": err,
			}).Error("Failed to update doctor")
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update doctor"})
		}

		return c.JSON(http.StatusOK, map[string]string{"message": "Doctor updated successfully"})
	}
}

func (d *DoctorHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		idStr := c.Param("id")
		if idStr == "" {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "id must be provided"})
		}

		id, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id format"})
		}

		err = d.repo.DeleteDoctor(c.Request().Context(), uint(id))
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"id":    id,
				"error": err,
			}).Error("Failed to delete doctor")
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to delete doctor",
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "Doctor deleted successfully",
			"id":      id,
		})
	}
}

func (d *DoctorHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := d.repo.GetAll(c.Request().Context())
		if err != nil {
			logrus.Error("Failed to get doctors: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get doctors"})
		}
		return c.JSON(http.StatusOK, result)
	}
}

func (d *DoctorHandler) Count() echo.HandlerFunc {
	return func(c echo.Context) error {
		count, err := d.repo.Count(c.Request().Context())
		if err != nil {
			logrus.Error("Error getting the count of doctors", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to count doctors"})
		}
		return c.JSON(http.StatusOK, count)
	}
}
