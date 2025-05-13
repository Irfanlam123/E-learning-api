package handlers

import (
	"doctor-on-demand/config"
	"doctor-on-demand/models"
	repository "doctor-on-demand/repositories"
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type ICourseHandler interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	// Count() echo.HandlerFunc
	GetByID() echo.HandlerFunc
}

type CourseHandler struct {
	repo repository.ICourseRepository
}

func NewCourseHandler(repo repository.ICourseRepository) *CourseHandler {
	return &CourseHandler{repo: repo}
}

func (h *CourseHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := models.Course{}
		id, err := config.GenerateId()
		if err != nil {
			return err
		}
		req.ID = id
		req.CreatedAt = time.Now().UTC()
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		err = h.repo.Create(c.Request().Context(), &req)
		if err != nil {
			logrus.Error("Failed to create course: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create article"})
		}
		return c.JSON(http.StatusOK, req)
	}
}

func (h *CourseHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		courses, err := h.repo.GetAll(c.Request().Context())
		if err != nil {
			logrus.Error("Failed to get all courses: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get all courses"})
		}
		return c.JSON(http.StatusOK, courses)
	}
}
func (h *CourseHandler) GetByID()echo.HandlerFunc{
	return func(c echo.Context) error{
	idParam := c.Param("id")
		id, _ := strconv.Atoi(idParam)
		result,err:=h.repo.GetByID(c.Request().Context(),uint(id))
		if err!=nil{
			return err
		}
		return c.JSON(http.StatusOK,result)
	}
}
