package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/shegai01/timer/internal/model"
	"github.com/shegai01/timer/internal/storage"
)

func initContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

type Timerhandler struct {
	storage *storage.Storage
}

func NewTimer(db *storage.Storage) *Timerhandler {
	return &Timerhandler{
		storage: db,
	}
}
func (h *Timerhandler) CreateTimer(w http.ResponseWriter, r *http.Request) {
	var timer *model.Timer
	title := r.URL.Query().Get("title")
	timer, err := h.storage.CreateTimer(title)
	if err != nil {
		return
	}
	log.Println("timer created")
	fmt.Println("create timer successfully")
	initContentType(w)
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(timer); err != nil {
		http.Error(w, "encoding", http.StatusBadRequest)
		log.Println("json encoder in 'CreateTimer' failed")
		return
	}

}

// func (h *Timerhandler) ShowAllTimersHandler(w http.ResponseWriter, r *http.Request) ([]model.Timer, error) {
// 	var allTimers []*model.Timer
// 	w.WriteHeader(http.StatusOK)
// 	initContentType(w)
// 	allTimers, err := h.storage.ShowAllTimers()
// 	if err != nil {
// 		http.Error(w, "cant get all timers", http.StatusBadRequest)
// 		log.Println(err)
// 		return nil, err
// 	}
// 	return allTimers, nil
// }
