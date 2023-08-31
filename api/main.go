package main

import (
	"log"
	"net/http"
	"os"

	"api/handlers"
	"api/storage"
)

func main() {
	log.Println("Creating DB connection")
	if err := storage.InitDB(); err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Setting up handlers")
	mux := http.NewServeMux()
	mux.HandleFunc("/foods", handlers.Foods)
	mux.HandleFunc("/measure-names", handlers.MeasureNames)

	log.Println("Listening on :" + port)
	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
