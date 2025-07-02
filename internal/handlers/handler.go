package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

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
	tittle := r.URL.Query().Get("tittle")
	timer, err := h.storage.CreateTimer(tittle)
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

func (h *Timerhandler) ShowAllTimersHandler(w http.ResponseWriter, r *http.Request) {
	var allTimers []*model.Timer
	initContentType(w)
	allTimers, err := h.storage.ShowAllTimers()
	if err != nil {
		http.Error(w, "cant get all timers", http.StatusBadRequest)
		log.Println(err)
	}
	timers, err := json.MarshalIndent(allTimers, "", " ")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		log.Println("encoding failed")

		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(timers))
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		log.Println("encoding failed")

		return
	}
	err = os.WriteFile("timer.json", timers, 0666)
	if err != nil {
		log.Println("save in file failed")
		return
	}
	log.Println("saved in file timer.json")

}

func (h *Timerhandler) DeletebyID(w http.ResponseWriter, r *http.Request) {
	idTask := r.URL.Query().Get("id")
	initContentType(w)
	w.WriteHeader(http.StatusOK)
	strID, err := strconv.Atoi(idTask)
	if err != nil {
		http.Error(w, "incorrect input", http.StatusOK)
		log.Println("incorrect id", err)
		return
	}
	if err := h.storage.DeletebyID(strID); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		log.Println("can't delete timer")

		return
	}
}
func (h *Timerhandler) GetTimerbyID(w http.ResponseWriter, r *http.Request) {
	idTask := r.URL.Query().Get("tittle")
	// w.Header().Set(contentType, applicationsJson)
	// id := r.URL.Query()
	initContentType(w)

	timer, err := h.storage.GetTimerbyTittle(idTask)
	if err != nil {
		http.Error(w, "delete func failed", http.StatusBadRequest)
		log.Println(err)

		return
	}
	timerbyid, err := json.MarshalIndent(timer, "", " ")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		log.Println("marshaling failed", err)
		return
	}
	w.Write([]byte(timerbyid))
	w.WriteHeader(http.StatusAccepted)
}
