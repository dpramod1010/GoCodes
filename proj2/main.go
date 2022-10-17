package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/cars", GetCars).Methods("GET")
	r.HandleFunc("/cars/{id}", GetCar).Methods("GET")
	r.HandleFunc("/cars", CreateCar).Methods("POST")
	r.HandleFunc("/cars/{id}", UpdateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", DeleteCar).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", r))
}

func main() {
	InitialMigration()
	initializeRouter()
}
