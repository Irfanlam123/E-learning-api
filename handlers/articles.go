package handlers

import (
	"doctor-on-demand/config"
	"doctor-on-demand/models"
	repository "doctor-on-demand/repositories"
	"net/http"

	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type IArticleHandler interface {
	Create() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	// Count() echo.HandlerFunc
}

type ArticleHandler struct {
	repo repository.IArticlesRepository
}

func NewArticleHandler(repo repository.IArticlesRepository) *ArticleHandler {
	return &ArticleHandler{repo: repo}
}

func (d *ArticleHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := models.Articles{}
		id, _ := config.GenerateId()
		req.ID = id
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request payload"})
		}

		err := d.repo.Create(c.Request().Context(), &req)
		if err != nil {
			logrus.Error("Failed to create article: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create article"})
		}

		return c.JSON(http.StatusCreated, req) // Return created doctor with ID
	}
}

func (d *ArticleHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		result, err := d.repo.GetAll(c.Request().Context())
		if err != nil {
			logrus.Error("Failed to get article: ", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to get article"})
		}
		return c.JSON(http.StatusOK, result)
	}
}

// func (d *ArticleHandler) Count() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		count, err := d.repo.Count(c.Request().Context())
// 		if err != nil {
// 			logrus.Error("Error getting the count of doctors", err)
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to count doctors"})
// 		}
// 		return c.JSON(http.StatusOK, count)
// 	}
// }
