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
		log.Println("databaseURI string not founded")
		return
	}
	databaseuri := os.Getenv("app_database_uri")

	router := mux.NewRouter()

	db := storage.NewStorage()
	if err := db.ConnectDB(databaseuri); err != nil {
		log.Fatal("connection failed")
		return
	}
	defer func() {
		if err := db.CloseDB(); err != nil {
			return
		}
	}()

	log.Println("connection database successfully")

	timer := handlers.NewTimer(db)

	router.HandleFunc("/create", timer.CreateTimer)
	router.HandleFunc("/show", timer.ShowAllTimersHandler)
	router.HandleFunc("/get", timer.GetTimerbyID)
	router.HandleFunc("/delete", timer.DeletebyID)
	router.HandleFunc("/stop", timer.StopTime)
	router.HandleFunc("/duration", timer.Duration)

	appPOrt := os.Getenv("app_port")
	// server := &http.Server{
	// 	Addr:    appPOrt,
	// 	Handler: router,
	// }
	if err := http.ListenAndServe(appPOrt, router); err != nil {
		log.Fatalln("server not starting")
		log.Println("timer starting")
	}
}
