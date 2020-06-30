package master

import (
	"database/sql"
	"goMuxAPI/main/master/controllers"
	"goMuxAPI/main/master/repositories"
	"goMuxAPI/main/master/usecases"

	"github.com/gorilla/mux"
)

//Init for initialize route
func Init(r *mux.Router, db *sql.DB) {
	studentRepo := repositories.InitStudentRepoImpl(db)
	studentUsecase := usecases.InitStudentUsecaseImpl(studentRepo)
	controllers.StudentController(r, studentUsecase)
	// controllers.StudentController(r)
}
