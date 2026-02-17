package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Handler for POST /v1/cars/read
func (a *applicationDependencies) readCarsHandler(w http.ResponseWriter, r *http.Request) {
	// create a struct to hold a car
	var incomingData struct {
		Make  string `json:"make"`
		Model string `json:"model"`
		Year  int    `json:"year"`
	}

	err := json.NewDecoder(r.Body).Decode(&incomingData)
	if err != nil {
		a.errorResponseJSON(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// for now display the result
	fmt.Fprintf(w, "%+v\n", incomingData)
}

// Handler for GET /v1/cars/write
func (a *applicationDependencies) writeCarsHandler(w http.ResponseWriter, r *http.Request) {
	cars := []map[string]interface{}{
		{"make": "Toyota", "model": "Camry", "year": 2020},
		{"make": "Honda", "model": "Civic", "year": 2019},
		{"make": "Ford", "model": "Mustang", "year": 2021},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}
