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
		log.Println("")
		return
	}
	databaseuri := os.Getenv("app_database_uri")

	router := mux.NewRouter()

	db := storage.NewStorage()
	if err := db.ConnectDB(databaseuri); err != nil {
		log.Fatal("connection failed")
		return
	}
	timer := handlers.NewTimer(db)
	router.HandleFunc("/create", timer.CreateTimer)
	router.HandleFunc("/showalltimers", timer.ShowAllTimersHandler)
	appPOrt := os.Getenv("app_port")
	if err := http.ListenAndServe(appPOrt, router); err != nil {
		log.Fatalln("server not starting")
	}
}
