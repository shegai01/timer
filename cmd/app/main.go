package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/shegai01/timer/internal/handlers"
	"github.com/shegai01/timer/internal/storage"
)

type Config struct {
	DatabaseURI string `json:"database_uri"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Printf("godotenv.Load: %s\n", err)
		return
	}
	databaseuri := os.Getenv("app_database_uri")

	router := mux.NewRouter()

	db := storage.NewStorage()
	if err := db.Connect(databaseuri); err != nil {
		log.Fatalf("Connect failed: %s\n", err)
		return
	}
	defer func() {
		if err := db.Close(); err != nil {

			return
		}
	}()

	log.Println("connection database successfully")

	timer := handlers.NewTimer(db)

	router.HandleFunc("/create", timer.Create)
	router.HandleFunc("/show", timer.ShowAll)
	router.HandleFunc("/get", timer.GetbyID)
	router.HandleFunc("/delete", timer.Delete)
	router.HandleFunc("/stop", timer.Stop)

	appPOrt := os.Getenv("app_port")

	if err := http.ListenAndServe(appPOrt, router); err != nil {
		log.Fatalf("ListenAndServe %s\n", err)
	}
}
