package main

import (
	"log"
	"net/http"

	"api/handlers"
	"api/storage"
)

func main() {
	log.Println("Creating DB connection")
	if err := storage.InitDB(); err != nil {
		log.Fatal(err)
	}

	log.Println("Setting up handlers")
	mux := http.NewServeMux()
	mux.HandleFunc("/foods", handlers.Foods)
	mux.HandleFunc("/measure-names", handlers.MeasureNames)

	log.Println("Listening on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
