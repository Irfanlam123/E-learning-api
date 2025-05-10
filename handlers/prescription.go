package handlers

import (
	"doctor-on-demand/config"
	"doctor-on-demand/models"
	repository "doctor-on-demand/repositories"
	"net/http"

	"github.com/labstack/echo"
)

type IPrescriptionHandler interface {
	Create() echo.HandlerFunc
}
type PrescriptionHandler struct {
	repo repository.IPrescriptionRepositoy
}

func NewPrescriptionHandler(repo repository.IPrescriptionRepositoy) *PrescriptionHandler {
	return &PrescriptionHandler{repo: repo}
}

func (h *PrescriptionHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		var prescription models.Prescription
		id, _ := config.GenerateId()
		prescription.ID = id
		if err := c.Bind(&prescription); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{
				"error": "Invalid request payload",
			})
		}

		createdPrescription, err := h.repo.Generate(c.Request().Context(), &prescription)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Failed to create prescription",
			})
		}

		return c.JSON(http.StatusCreated, createdPrescription)
	}
}
