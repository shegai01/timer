package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hashicorp/consul/api"
	"github.com/shegai01/timer/internal"
)

const (
	contentType      = "Content-Type"
	applicationsJson = "application/json"
	configKey        = "APP_CONFIG"
	confgiVal        = `{
		"app" : {"port" : "8080"},
		"db" : {"dns" : "host=localhost port=54321 user=alex01 password=pwd1234 dbname=timer"}
	
	}`
)

var cfg Config

func main() {
	fmt.Println("start timer")
	// a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Println(http.StatusBadRequest)
		log.Println("init client kv store")
		log.Println(err)
		return
	}
	kv := client.KV()
	// put configs kv pair
	// putConfig := &api.KVPair{Key: "DB_DNS", Value: []byte("postgresql://alex01:pwd1234@localhost:54321/timer")}
	putConfig := &api.KVPair{Key: "APP_CONFIG", Value: []byte(confgiVal)}
	_, err = kv.Put(putConfig, nil)
	if err != nil {
		log.Println(http.StatusBadRequest)
		log.Println("configs put failed", err)
		return
	}
	pair, _, err := kv.Get("DB_DNS", nil)
	if err != nil {
		log.Println(http.StatusBadRequest)
		log.Println("configs incorrect")
		return
	}
	//
	if err := json.Unmarshal([]byte(pair.Value), &cfg); err != nil {
		log.Println(http.StatusBadRequest)
		return
	}
	// a new gorilla router

	router := mux.NewRouter()
	tracker, err := internal.NewTimerDB("postgresql://alex01:pwd1234@localhost:54321")
	if err != nil {
		log.Println(http.StatusBadRequest)
		return
	}
	// handler get all timers")
	router.HandleFunc("/all", func(w http.ResponseWriter, r *http.Request) {
		// w.Write([]byte("hello"))
		w.WriteHeader(http.StatusOK)
		w.Header().Set(contentType, applicationsJson)

		var arrTimer []internal.Timer
		arrTimer, err = tracker.ShowAllTimers()
		if err != nil {
			log.Println(http.StatusBadRequest)
			log.Println("", err)
			return
		}
		timers, err := json.MarshalIndent(arrTimer, "", "\t")
		if err != nil {
			log.Println(http.StatusBadRequest)
			log.Println("encoding failed")
			return
		}
		_, err = w.Write([]byte(timers))
		if err != nil {
			return
		}
	})
	// create timers
	router.HandleFunc("/start", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(contentType, applicationsJson)
		nameTask := r.URL.Query().Get("name")
		log.Printf("getting url params by query %s\n", nameTask)
		w.WriteHeader(http.StatusOK)
		timer, err := tracker.CreateTimer(nameTask)
		if err != nil {
			http.Error(w, "creating failed", http.StatusBadRequest)
			log.Println(err)
			return
		}
		if err := json.NewEncoder(w).Encode(timer); err != nil {
			http.Error(w, "encoding failed", http.StatusBadRequest)
			log.Println(err)
			return
		}
	})
	// delete timers
	router.HandleFunc("/stop/{id}", func(w http.ResponseWriter, r *http.Request) {
		idTask := r.URL.Query().Get("id")
		// w.Header().Set(contentType, applicationsJson)
		// id := r.URL.Query()

		strID, err := strconv.Atoi(idTask)
		if err != nil {
			log.Println("incorrect id", err)
		}
		if err = tracker.DeletebyID(strID); err != nil {
			http.Error(w, "delete func failed", http.StatusBadRequest)
			log.Println(err)
			return
		}
		w.WriteHeader(http.StatusAccepted)
	})
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Println("can't listen port 8080")
		os.Exit(1)
		return
	}
}

// todo:
// git checkout -b new branch
// git add .
// git commit -m "start" && git commit -m "stop"
//
