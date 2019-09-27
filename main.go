package main

import (
	"log"
	"net/http"

	"chitChat/controllers"
	"chitChat/models"

	"github.com/gorilla/mux"
)

func main() {
	models.InitDB("user=ryanback dbname=chit_chat sslmode=disable")

	r := mux.NewRouter()
	r.HandleFunc("/users", controllers.UsersIndex).Methods("GET")
	r.HandleFunc("/users/{id}", controllers.UsersShow).Methods("GET")
	log.Fatal(http.ListenAndServe(":3001", r))
}
