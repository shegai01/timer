package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

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

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		log.Fatalln("env is missing: APP_PORT")
	}

	server := &http.Server{
		Addr:    ":" + appPort,
		Handler: handlers.NewTimerHandler(db),
	}

	var wg sync.WaitGroup
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := server.ListenAndServe(); err != nil {
			return
		}
	}()

	<-signalChan

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("server.Shutdown: %v\n", err)
	}

	wg.Wait()
}
