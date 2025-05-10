package handlers

import (
	"doctor-on-demand/config"
	"doctor-on-demand/models"
	"doctor-on-demand/password"
	repository "doctor-on-demand/repositories"
	"net/http"

	"github.com/labstack/echo"
)

type IUserHandler interface {
	Create() echo.HandlerFunc
	GetUserByEmail() echo.HandlerFunc
}

type UserHandler struct {
	user models.User
	repo repository.IUserRepository
}

func NewUserHandler(repo repository.IUserRepository) *UserHandler {
	return &UserHandler{
		repo: repo,
	}
}

func (p *UserHandler) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{}
		id, _ := config.GenerateId()
		user.ID = id

		if err := c.Bind(&user); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		hashedPassword, err := password.HashPassword(user.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Error hashing password")
		}
		user.Password = hashedPassword
		err = p.repo.Create(c.Request().Context(), &user)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Failed to create user")
		}
		return c.JSON(http.StatusCreated, user)
	}
}

func (p *UserHandler) GetUserByEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		type req struct {
			Email    string `json:"email"`    // Correct JSON tags for binding
			Password string `json:"password"` // Correct JSON tags for binding
		}
		r := new(req)
		// Bind the request body to the struct
		if err := c.Bind(r); err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		// Get the user by email
		user, err := p.repo.GetUserByEmail(c.Request().Context(), r.Email)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid credentials ")
		}

		// Verify the password with the stored hash
		if err := password.VerifyPassword(r.Password, user.Password); err != nil {
			// If passwords don't match, return unauthorized
			return c.JSON(http.StatusUnauthorized, "Invalid credentials")
		}

		// If login is successful
		return c.JSON(http.StatusOK, user)
	}
}
