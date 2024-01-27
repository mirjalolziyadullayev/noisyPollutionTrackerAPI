package Handlers

import (
	"encoding/json"
	"net/http"
	"noisyPollutionTrackerAPI/models"
	"os"
)

func CityHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getCities(w, r)
	}
}

func getCities(w http.ResponseWriter, r *http.Request) {
	var CitiesData []models.CityModel
	ByteCollection, _ := os.ReadFile("db/all.json")
	json.Unmarshal(ByteCollection, &CitiesData)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(CitiesData)
}
