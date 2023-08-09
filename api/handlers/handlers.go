package handlers

import (
	"api/storage"
	"encoding/json"
	"log"
	"net/http"
)

func Foods(w http.ResponseWriter, r *http.Request) {

	foodNames, err := storage.AllFoodNames(r.URL.Query())
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(foodNames)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func MeasureNames(w http.ResponseWriter, r *http.Request) {
	measureNames, err := storage.AllMeasureNames()
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(measureNames)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
