package initializers

import (
	"doctor-on-demand/config"
	"doctor-on-demand/handlers"
	repository "doctor-on-demand/repositories"
	"doctor-on-demand/routes"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type App struct {
	DB       *gorm.DB
	Handlers struct {
		Teacher *handlers.TeacherHandler
		Article *handlers.ArticleHandler
		User    *handlers.UserHandler
		Course  *handlers.CourseHandler
	}
}

func Initializers() *App {
	// Initialize database connection
	db := config.ConnectDB()

	// 2. Initialize Repositories
	teacherRepo := repository.NewTeacherRepository(db)

	articleRepo := repository.NewArticlesRepository(db)

	userrepo := repository.NewUserRepository(db)
	courseRepo := repository.NewCourseRepository(db)

	app := &App{
		DB: db,
		Handlers: struct {
			Teacher *handlers.TeacherHandler
			Article *handlers.ArticleHandler
			User    *handlers.UserHandler
			Course  *handlers.CourseHandler
		}{
			Teacher: handlers.NewTeacherHandler(teacherRepo),
			Article: handlers.NewArticleHandler(articleRepo),
			User:    handlers.NewUserHandler(userrepo),
			Course:  handlers.NewCourseHandler(courseRepo),
		},
	}
	return app
}
func (a *App) SetupRoutes(e *echo.Echo) {

	routes.TeacherRoutes(e, a.Handlers.Teacher)
	routes.Article(e, a.Handlers.Article)
	routes.UserRoutes(e, a.Handlers.User)
	routes.CourseRoute(e, a.Handlers.Course)

}
