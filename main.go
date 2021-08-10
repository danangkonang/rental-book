package main

import (
	"fmt"
	"net/http"

	"github.com/danangkonang/rental-book/config"
	"github.com/danangkonang/rental-book/helper"
	"github.com/danangkonang/rental-book/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter().StrictSlash(true)

	router.StudentRouter(r, config.Connection())
	router.BookRouter(r, config.Connection())
	router.TransactionRouter(r, config.Connection())

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		helper.MakeRespon(w, 200, "ok", nil)
	})

	fmt.Println("local server started at http://localhost:9000")
	header := []string{
		"X-Requested-With",
		"Access-Control-Allow-Origin",
		"Content-Type",
	}
	method := []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}
	origin := []string{"*"}
	http.ListenAndServe(":9000", handlers.CORS(
		handlers.AllowedHeaders(header),
		handlers.AllowedMethods(method),
		handlers.AllowedOrigins(origin),
	)(r))

}
