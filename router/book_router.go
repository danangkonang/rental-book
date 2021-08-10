package router

import (
	"github.com/danangkonang/rental-book/config"
	"github.com/danangkonang/rental-book/controller"
	"github.com/danangkonang/rental-book/repository"
	"github.com/danangkonang/rental-book/service"
	"github.com/gorilla/mux"
)

func BookRouter(router *mux.Router, db *config.DB) {
	rest := controller.NewControllerBook(service.NewServiceBook(repository.NewRepositoryBook(db)))
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/book", rest.CreateBook).Methods("POST")
	v1.HandleFunc("/books", rest.FindBooks).Methods("GET")
	v1.HandleFunc("/book", rest.UpdateBook).Methods("PUT")
	v1.HandleFunc("/book", rest.DeleteBook).Methods("DELETE")
}
