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

type ITeacher interface {
	GetById() echo.HandlerFunc
	Create() echo.HandlerFunc
	Update() echo.HandlerFunc
	Delete() echo.HandlerFunc
	GetAll() echo.HandlerFunc
}

type TeacherHandler struct {
	teacher models.Teacher
	repo    repository.ITeacherRepository
}

func NewTeacherHandler(repo repository.ITeacherRepository) *TeacherHandler {
	return &TeacherHandler{
		repo: repo,
	}
}

func (p *TeacherHandler) GetById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			logrus.Error("Invalid ID format: ", err)
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}

		patient, err := p.repo.GetByID(c.Request().Context(), uint(id))
		if err != nil {
			logrus.Error("Error getting patient: ", err)
			return c.JSON(http.StatusInternalServerError, "Failed to get patient")
		}
		return c.JSON(http.StatusOK, patient)
	}
}

func (p *TeacherHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		patient := models.Teacher{}
		id, _ := config.GenerateId()
		patient.ID = id
		if err := c.Bind(&patient); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		err := p.repo.CreateTeacher(c.Request().Context(), &patient)
		if err != nil {
			logrus.Error("Error creating patient: ", err)
			return c.JSON(http.StatusInternalServerError, "Failed to create patient")
		}
		return c.JSON(http.StatusCreated, patient)
	}
}

func (p *TeacherHandler) Update() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			logrus.Error("Invalid ID format: ", err)
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}

		patient := models.Teacher{}
		if err := c.Bind(&patient); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		err = p.repo.UpdateTeacher(c.Request().Context(), uint(id), &patient)
		if err != nil {
			logrus.Error("Error updating patient: ", err)
			return c.JSON(http.StatusInternalServerError, "Failed to update patient")
		}
		return c.JSON(http.StatusOK, patient)
	}
}

func (p *TeacherHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			logrus.Error("Invalid ID format: ", err)
			return c.JSON(http.StatusBadRequest, "Invalid ID")
		}

		err = p.repo.DeleteTeacher(c.Request().Context(), uint(id))
		if err != nil {
			logrus.Error("Error deleting patient: ", err)
			return c.JSON(http.StatusInternalServerError, "Failed to delete patient")
		}
		return c.NoContent(http.StatusOK)
	}
}

func (p *TeacherHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		result, err := p.repo.GetAll(c.Request().Context())
		if err != nil {
			logrus.Error("Error getting patients: ", err)
		}

		return c.JSON(http.StatusOK, result)
	}
}
