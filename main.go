package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hashicorp/consul/api"
	"github.com/shegai01/timer/internal/app"
	"github.com/shegai01/timer/internal/storage"
)

const (
	configKey = "APP_CONFIG"
	confgiVal = `{
		"app" : {"port" : "8080"},
		"db" : {"dsn" : "host=localhost port=54321 user=alex01 dbname=timerdb password=pwd1234 sslmode=disable"}
	
	}`
)

type Config struct {
	app *app.ConfigAPP
	db  *storage.ConfigDB
}

var cfg Config

func init() {
	fmt.Println("start timer")
	// a new client
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Println(http.StatusBadRequest)
		log.Println("init client kv store")
		log.Fatal(err)
		// os.Exit(1)
	}

	// store kv_consul
	kv := client.KV()
	putConfig := &api.KVPair{Key: "APP_CONFIG", Value: []byte(confgiVal)}
	_, err = kv.Put(putConfig, nil)
	if err != nil {
		log.Println(http.StatusBadRequest)
		log.Println("configs put failed", err)

		return
	}

	pair, _, err := kv.Get("APP_CONFIG", nil)
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
}

func main() {

	config := storage.NewConfigStorage()
	server := app.NewAPI(app.NewConfig())
	if err := server.Start(); err != nil {
		return
	}
	// // tracker, err := internal.NewTimerDB(cfg.Dsn.Dsn)
	// if err != nil {
	// 	log.Println(http.StatusBadRequest)

	// 	return
	// }
	// if err := http.ListenAndServe(":"+"%s", cfg.App.Port); err != nil {
	// 	log.Println("can't listen port 8080")
	// 	os.Exit(1)

	// 	return
	// }
	log.Println("connection OK!")
}

// todo:
// git checkout -b new branch
// git add .
// git commit -m "start" && git commit -m "stop"
//
