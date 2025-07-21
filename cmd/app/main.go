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
		log.Fatalf("godotenv.Load: %s\n", err)
	}

	databaseURI := os.Getenv("APP_DATABASE_URI")
	if databaseURI == "" {
		log.Fatalln("env is missing: APP_DATABASE_URI")
	}

	router := mux.NewRouter()

	db := storage.NewStorage()
	if err := db.Connect(databaseURI); err != nil {
		log.Fatalf("db.Connect: %s\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("db.Close: %s\n", err)
		}
	}()

	log.Println("connection to database established successfully")

	timer := handlers.NewTimer(db)

	router.HandleFunc("/create", timer.Create)
	router.HandleFunc("/show", timer.ShowAll)
	router.HandleFunc("/get", timer.GetbyID)
	router.HandleFunc("/delete", timer.Delete)
	router.HandleFunc("/stop", timer.Stop)

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatalln("env is missing: APP_PORT")
	}

	if err := http.ListenAndServe(":"+appPort, router); err != nil {
		log.Fatalf("ListenAndServe: %s\n", err)
	}
}
