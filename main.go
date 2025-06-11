package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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
	// var cfg Config
	// if err := envconfig.Process("APP", &cfg); err != nil {
	// 	log.Println("can't read config file")
	// 	return
	// }

	router := mux.NewRouter()
	// urlExample :=
	tracker, err := internal.NewTimerDB("postgres://alex01:pwd1234@localhost:54321/")
	if err != nil {
		return
	}
	// router.HandleFunc("/pause")
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("hello"))
		arrTimer, err := tracker.ShowAllTimers()
		if err != nil {
			log.Println("", err)
			return
		}
		timers, err := json.MarshalIndent(arrTimer, "", "\t")
		if err != nil {
			log.Println("encoding failed")
			return
		}
		w.Write([]byte(timers))

		w.WriteHeader(http.StatusOK)
	})
	router.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		tracker.CreateTimer()
	})
	router.HandleFunc("create", func(w http.ResponseWriter, r *http.Request) {
		r.URL.Query() //todo разобрать
	})
	if err := (http.ListenAndServe(":8080", router)); err != nil {
		log.Println("error in server")
		return
	}
	log.Printf("server listen at : %s", port)

}
