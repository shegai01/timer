package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kelseyhightower/envconfig"
	"github.com/shegai01/timer/internal"
)

const (
	port   = "8080"
	portdb = "54321"
)

type Config struct {
	Port string
	DB   struct {
		DNS string `default:""`
	}
}

func main() {
	fmt.Println("start timer")
	var cfg Config
	if err := envconfig.Process("APP", &cfg); err != nil {
		log.Println("can't read config file")
		return
	}

	router := mux.NewRouter()

	tracker, err := internal.NewTimerDB(&cfg)
	if err != nil {
		return
	}
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})
	router.HandleFunc("/starttimer", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)

	})
	if err := (http.ListenAndServe(":"+cfg.Port, router)); err != nil {
		log.Println("error in server")
		return
	}
	log.Printf("server listen at : %s", cfg.Port)

}
