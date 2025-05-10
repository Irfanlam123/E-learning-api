package initializers

import (
	"doctor-on-demand/config"
	"doctor-on-demand/handlers"
	repository "doctor-on-demand/repositories"
	"doctor-on-demand/routes"
	"doctor-on-demand/services"

	"github.com/labstack/echo"
	"gorm.io/gorm"
)

type App struct {
	DB       *gorm.DB
	Handlers struct {
		Doctor         *handlers.DoctorHandler
		Patient        *handlers.PatientHandler
		DoctorSchedule *handlers.DoctorScheduleHandler
		Appointment    *handlers.AppointmentHandler
		Article        *handlers.ArticleHandler
		Prescription   *handlers.PrescriptionHandler
		User           *handlers.UserHandler
		Medicine       *handlers.MedicineHandler
	}
}

func Initializers() *App {
	// Initialize database connection
	db := config.ConnectDB()
	// Initialize PDF service
	pdfService := services.NewPDFService("./prescriptions")
	// 2. Initialize Repositories
	doctorRepo := repository.NewDoctorRepository(db)
	patientRepo := repository.NewPatientRepository(db)
	scheduleRepo := repository.NewDoctorScheduleRepository(db)
	appointmentRepo := repository.NewAppointmentRepository(db)
	articleRepo := repository.NewArticlesRepository(db)
	prescriptionRepo := repository.NewPrescriptionRepository(db, pdfService)
	userrepo := repository.NewUserRepository(db)
	medicineRepo := repository.NewMedicineRepository(db)
	app := &App{
		DB: db,
		Handlers: struct {
			Doctor         *handlers.DoctorHandler
			Patient        *handlers.PatientHandler
			DoctorSchedule *handlers.DoctorScheduleHandler
			Appointment    *handlers.AppointmentHandler
			Article        *handlers.ArticleHandler
			Prescription   *handlers.PrescriptionHandler
			User           *handlers.UserHandler
			Medicine       *handlers.MedicineHandler
		}{
			Doctor:         handlers.NewDoctorHandler(doctorRepo),
			Patient:        handlers.NewPatientHandler(patientRepo),
			DoctorSchedule: handlers.NewDoctorScheduleHandler(scheduleRepo),
			Appointment:    handlers.NewAppointmentHandler(appointmentRepo),
			Article:        handlers.NewArticleHandler(articleRepo),
			Prescription:   handlers.NewPrescriptionHandler(prescriptionRepo),
			User:           handlers.NewUserHandler(userrepo),
			Medicine:       handlers.NewMedicineHandler(medicineRepo),
		},
	}
	return app
}
func (a *App) SetupRoutes(e *echo.Echo) {
	// Create prescriptions directory if not exists
	if err := config.CreateDirectoryIfNotExist("./prescriptions"); err != nil {
		panic(err)
	}
	routes.Routes(e, a.Handlers.Doctor)
	routes.PatientRoutes(e, a.Handlers.Patient)
	routes.DoctorSchedule(e, a.Handlers.DoctorSchedule)
	routes.AppointmentRoutes(e, a.Handlers.Appointment)
	routes.Article(e, a.Handlers.Article)
	routes.PrescriptionRoutes(e, a.Handlers.Prescription)
	routes.UserRoutes(e, a.Handlers.User)
	routes.MedicineRoutes(e, a.Handlers.Medicine)
}
