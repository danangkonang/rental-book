package router

import (
	"github.com/danangkonang/rental-book/config"
	"github.com/danangkonang/rental-book/controller"
	"github.com/danangkonang/rental-book/repository"
	"github.com/danangkonang/rental-book/service"
	"github.com/gorilla/mux"
)

func StudentRouter(router *mux.Router, db *config.DB) {
	rest := controller.NewControllerStudent(
		service.NewServiceStudent(
			repository.NewRepositoryStudent(db),
		),
	)
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/student", rest.CreateStudent).Methods("POST")
	v1.HandleFunc("/students", rest.FindStudents).Methods("GET")
	v1.HandleFunc("/student", rest.UpdateStudent).Methods("PUT")
	v1.HandleFunc("/student", rest.DeleteStudent).Methods("DELETE")
}
