package routes

import (
	"doctor-on-demand/handlers"

	"github.com/labstack/echo"
)

func Routes(e *echo.Echo, DoctorHandler *handlers.DoctorHandler) {
	e.POST("/doctor", DoctorHandler.Create())
	e.GET("/doctor/:id", DoctorHandler.GetById())
	e.PUT("/doctor/:id", DoctorHandler.Update())
	e.DELETE("/doctor/:id", DoctorHandler.Delete())
	e.GET("/doctors", DoctorHandler.GetAll())
	e.GET("/doctorCount", DoctorHandler.Count())

}

func PatientRoutes(e *echo.Echo, patientHandler *handlers.PatientHandler) {
	e.GET("/patients", patientHandler.GetAll())
	e.GET("/patient/:id", patientHandler.GetById())
	e.POST("/patient", patientHandler.Create())
	e.PUT("/patient/:id", patientHandler.Update())
	e.DELETE("/patient/:id", patientHandler.Delete())
}

func AppointmentRoutes(e *echo.Echo, appointmentHandler *handlers.AppointmentHandler) {
	e.GET("/appointment/:id", appointmentHandler.GetAppointmentsByDoctorID())
	e.POST("/appointment", appointmentHandler.Create())
	// e.PUT("/appointment/:id", appointmentHandler.Update())
	// e.DELETE("/appointment/:id", appointmentHandler.Delete())
}

func DoctorSchedule(e *echo.Echo, ScheduleHandler *handlers.DoctorScheduleHandler) {
	e.POST("/schedule", ScheduleHandler.Create())
	e.GET("/schedule/:id", ScheduleHandler.GetByID())
	e.PUT("/schedule/:id", ScheduleHandler.Update())

}
func Article(e *echo.Echo, ArticleHandler *handlers.ArticleHandler) {
	e.POST("/article", ArticleHandler.Create())
	e.GET("/articles", ArticleHandler.GetAll())

}

func PrescriptionRoutes(e *echo.Echo, prescriptionHandler *handlers.PrescriptionHandler) {

	e.POST("/prescription", prescriptionHandler.Create())
	// g.GET("/:id/download", h.DownloadPrescription)
	// g.GET("/patient/:patientId", h.GetPatientPrescriptions)
}
func UserRoutes(e *echo.Echo, userHandler *handlers.UserHandler) {

	e.POST("/signup", userHandler.Create())
	e.POST("/login", userHandler.GetUserByEmail())
}
func MedicineRoutes(e *echo.Echo, medicineHandler *handlers.MedicineHandler) {

	e.POST("/medicine", medicineHandler.Create())
	e.GET("/medicine/:id", medicineHandler.GetById())
	e.GET("/medicines", medicineHandler.GetAll())
	e.PUT("/medicine/:id", medicineHandler.Update())
}
