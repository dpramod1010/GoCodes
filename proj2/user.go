package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const DNS = "root:Expleo@1010@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"

type Car struct {
	gorm.Model
	CarName   string `json:"carname"`
	CarNo     string `json:"carno"`
	StateCode string `json:"statecode"`
	Price     string `json:"price"`
}

func InitialMigration() {
	DB, err = gorm.Open(mysql.Open(DNS), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Cannot connect to DB")
	}
	DB.AutoMigrate(&Car{})
}

func GetCars(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cars []Car
	DB.Find(&cars)
	json.NewEncoder(w).Encode(cars)
}

func GetCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var car Car
	DB.First(&car, params["id"])
	json.NewEncoder(w).Encode(car)
}

func CreateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var car Car
	json.NewDecoder(r.Body).Decode(&car)
	DB.Create(&car)
	json.NewEncoder(w).Encode(car)
}

func UpdateCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var car Car
	DB.First(&car, params["id"])
	json.NewDecoder(r.Body).Decode(&car)
	DB.Save(&car)
	json.NewEncoder(w).Encode(car)
}

func DeleteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var car Car
	DB.Delete(&car, params["id"])
	json.NewEncoder(w).Encode("Car Details Deleted Successfully!")
}
