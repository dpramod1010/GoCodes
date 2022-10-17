package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Adhar struct {
	AdharNo   string `json:"ano"`
	Name      string `json:"name"`
	State     string `json:"state"`
	district  string `json:"district"`
	BirthDate string `json:"date"`
}
type Pan struct {
	PanNo     string `json:"pno"`
	Name      string `json:"name"`
	State     string `json:"state"`
	district  string `json:"district"`
	BirthYear string `json:"birthyear"`
}

var adharc []Adhar
var panc []Pan

//GerAdhar

func getAdharc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(adharc)
}

// Get Pan

func getPanc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(panc)
}

// delet Adhar

func deleteAdhar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range adharc {
		if item.AdharNo == params["ano"] {
			adharc = append(adharc[:index], adharc[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(adharc)
}

//delet pan

func deletePan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range panc {
		if item.PanNo == params["pno"] {
			panc = append(panc[:index], panc[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(panc)
}

//grt adhar
func getAdhar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range adharc {
		if item.AdharNo == params["ano"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

//get pan
func getPan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range panc {
		if item.PanNo == params["pno"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// create adhar

func createAdhar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var adhar Adhar
	_ = json.NewDecoder(r.Body).Decode(&adhar)
	adhar.AdharNo = strconv.Itoa(rand.Intn(101))
	adharc = append(adharc, adhar)
}

//create pan

func createPan(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var pan Pan
	_ = json.NewDecoder(r.Body).Decode(&pan)
	pan.PanNo = strconv.Itoa(rand.Intn(101))
	panc = append(panc, pan)
}

// update adhar
func updateAdharc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range adharc {
		if item.AdharNo == params["ano"] {
			adharc = append(adharc[:index], adharc[index+1:]...)
			var adhar Adhar
			_ = json.NewDecoder(r.Body).Decode(&adhar)
			adhar.AdharNo = params["ano"]
			adharc = append(adharc, adhar)
			json.NewEncoder(w).Encode(adhar)

		}
	}
}

//update pan

func updatePanc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range panc {
		if item.PanNo == params["pno"] {
			panc = append(panc[:index], panc[index+1:]...)
			var pan Pan
			_ = json.NewDecoder(r.Body).Decode(&pan)
			pan.PanNo = params["pno"]
			panc = append(panc, pan)
			json.NewEncoder(w).Encode(pan)

		}
	}
}

func main() {
	r := mux.NewRouter()

	adharc = append(adharc, Adhar{AdharNo: "8912 1234 3567 5543", Name: "Nikita Joshi", State: "Maharastra", district: "Pune", BirthDate: "10 Jun 1999"})
	panc = append(panc, Pan{PanNo: "8912 1234 3567 5543", Name: "Nikita Joshi", State: "Maharastra", district: "Pune", BirthYear: "10 Jun 1999"})
	//pan = append(adharc, Adhar{Number: "8888", Scode: "DL", Colour: "BLACK", Fuel: "diesel", Rto: &Rto{Name: "AUDI", Pyear: "2019"}})

	r.HandleFunc("/adharc", getAdharc).Methods("GET")
	r.HandleFunc("/adhar/{no}", getAdhar).Methods("GET")
	r.HandleFunc("/adharc", createAdhar).Methods("POST")
	r.HandleFunc("/adharc/{no}", updateAdharc).Methods("PUT")
	r.HandleFunc("/adharc/{no}", deleteAdhar).Methods("DELETE")

	r.HandleFunc("/panc", getPanc).Methods("GET")
	r.HandleFunc("/pan/{no}", getPan).Methods("GET")
	r.HandleFunc("/panc", createPan).Methods("POST")
	r.HandleFunc("/panc/{no}", updatePanc).Methods("PUT")
	r.HandleFunc("/panc/{no}", deletePan).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))

}

//go get -u github.com/gorilla/mux
//go mod vendor
//go mod tidy
