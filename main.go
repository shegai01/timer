package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/jackc/pgx/v4"
)

const (
	port   = "8080"
	portdb = "54321"
)

type Timer struct {
	ID int
	StartTime time.Time
	FinishTime time.Duration

}

func (t *Timer) Create() (*Timer,error){
return &Timer{
	StartTime: ,

}

}
func main() {
	fmt.Println("start timer")
	connStr := "postgres://:alex:pwd1234@localhost:54321/timer_db"
	db, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Println("db starting at 54321")
		return
	}
	defer db.Close(context.Background())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("hello"))
	})
	http.HandleFunc("/starttimer", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)

	})
	if err := (http.ListenAndServe(":"+port, nil)); err != nil {
		log.Println("error in server")
		return
	}
	log.Printf("server listen at : %s", port)

}
