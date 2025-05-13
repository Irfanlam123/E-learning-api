package routes

import (
	"doctor-on-demand/handlers"

	"github.com/labstack/echo"
)

func TeacherRoutes(e *echo.Echo, teacherHandler *handlers.TeacherHandler) {
	e.GET("/teachers", teacherHandler.GetAll())
	e.GET("/teacher/:id", teacherHandler.GetById())
	e.POST("/teacher", teacherHandler.Create())
	e.PUT("/teacher/:id", teacherHandler.Update())
	e.DELETE("/teacher/:id", teacherHandler.Delete())
	// e.GET("/count",te)
}

func Article(e *echo.Echo, ArticleHandler *handlers.ArticleHandler) {
	e.POST("/article", ArticleHandler.Create())
	e.GET("/articles", ArticleHandler.GetAll())

}
func CourseRoute(e *echo.Echo, courseHandler *handlers.CourseHandler) {
	e.POST("/course", courseHandler.Create())
	e.GET("/courses", courseHandler.GetAll())
	e.GET("/courses/:id", courseHandler.GetByID())

}

func UserRoutes(e *echo.Echo, userHandler *handlers.UserHandler) {

	e.POST("/signup", userHandler.Create())
	e.POST("/login", userHandler.GetUserByEmail())
}
