package router

import (
	"github.com/danangkonang/rental-book/config"
	"github.com/danangkonang/rental-book/controller"
	"github.com/danangkonang/rental-book/repository"
	"github.com/danangkonang/rental-book/service"
	"github.com/gorilla/mux"
)

func TransactionRouter(router *mux.Router, db *config.DB) {
	rest := controller.NewControllerTransaction(
		service.NewServiceTransaction(
			repository.NewRepositoryTransaction(db),
			repository.NewRepositoryBook(db),
		),
	)
	v1 := router.PathPrefix("/v1").Subrouter()
	v1.HandleFunc("/transactions", rest.FindTransactions).Methods("GET")
	v1.HandleFunc("/transaction", rest.CreateTansaction).Methods("POST")
	v1.HandleFunc("/transaction/return-book", rest.ReturnBook).Methods("PUT")
}
