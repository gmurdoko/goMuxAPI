package controllers

import (
	"encoding/json"
	"goMuxAPI/main/master/usecases"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// StudentHandler app
type StudentHandler struct {
	studentUsecase usecases.StudentUsecase
}

//StudentController app
func StudentController(r *mux.Router, s usecases.StudentUsecase) {
	studentHandler := StudentHandler{s}
	r.HandleFunc("/students", studentHandler.ListStudents).Methods(http.MethodGet)
	// r.HandleFunc("/hello", HelloWorld).Methods(http.MethodGet)

}

// //HelloWorld app
// func HelloWorld(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Hello with gorilla mux"))
// 	log.Println("Starting at port: ")
// }

//ListStudents app
func (s *StudentHandler) ListStudents(w http.ResponseWriter, r *http.Request) {
	students, err := s.studentUsecase.GetStudent()
	if err != nil {
		w.Write([]byte("Data Not Found"))
	}

	w.Header().Set("content-type", "application/json")
	byteOfStudents, err := json.Marshal(students)
	if err != nil {
		w.Write([]byte("Opps, Something Wrong"))
	}
	w.Write([]byte(byteOfStudents))
	log.Println("Hit point: get all data student")
}
