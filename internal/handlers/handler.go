package handlers

import (
	"encoding/json"
	"log"
	"net/http"
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

func (h *Timerhandler) Create(w http.ResponseWriter, r *http.Request) {
	var timer *model.Timer
	w.WriteHeader(http.StatusOK)
	tittle := r.URL.Query().Get("title")
	timer, err := h.storage.CreateTimer(tittle)
	if err != nil {
		log.Println("create timer failed", err)
		return
	}
	log.Println("timer created")

	initContentType(w)
	if err := json.NewEncoder(w).Encode(timer); err != nil {
		http.Error(w, "encoding", http.StatusBadRequest)
		log.Println("json encoder in 'CreateTimer' failed", err)
		return
	}

}

func (h *Timerhandler) ShowAll(w http.ResponseWriter, r *http.Request) {
	var allTimers []*model.Timer
	initContentType(w)
	allTimers, err := h.storage.ShowAllTimers()
	if err != nil {
		http.Error(w, "cant get all timers", http.StatusInternalServerError)
		log.Println(err)

		return
	}
	timers, err := json.MarshalIndent(allTimers, "", " ")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		log.Println("encoding failed", err)

		return
	}
	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(timers))
	if err != nil {
		log.Println("encoding failed", err)

		return
	}

}

func (h *Timerhandler) Delete(w http.ResponseWriter, r *http.Request) {
	initContentType(w)
	idTask := r.URL.Query().Get("id")
	strID, err := strconv.Atoi(idTask)
	if err != nil {
		http.Error(w, "incorrect input", http.StatusInternalServerError)
		log.Printf("strconv.atoi %s\n", err)
		return
	}
	if err := h.storage.Delete(strID); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Printf("h.storage.Delete %s\n", err)

		return
	}
	w.WriteHeader(http.StatusOK)
	log.Println("deleted")
}
func (h *Timerhandler) GetbyID(w http.ResponseWriter, r *http.Request) {

	idTask := r.URL.Query().Get("id")
	strID, err := strconv.Atoi(idTask)
	if err != nil {
		return
	}
	initContentType(w)

	timer, err := h.storage.GetTimerbyID(strID)
	if err != nil {
		http.Error(w, "get function failed", http.StatusBadRequest)
		log.Println("get in function storage failed", err)

		return
	}
	timerbyTittle, err := json.MarshalIndent(timer, "", " ")
	if err != nil {
		log.Println("marshaling failed", err)

		return
	}

	_, err = w.Write([]byte(timerbyTittle))
	if err != nil {
		return
	}

}
func (h *Timerhandler) Stop(w http.ResponseWriter, r *http.Request) {
	initContentType(w)
	idTask := r.URL.Query().Get("id")
	strID, err := strconv.Atoi(idTask)
	if err != nil {
		http.Error(w, "update at handler failed", http.StatusBadRequest)
		log.Println("stoptimer failed in storage", err)
		return
	}
	updatetimer, err := h.storage.GetTimerbyID(strID)
	if err != nil {
		return
	}

	timerJson, err := json.MarshalIndent(updatetimer, "", " ")
	if err != nil {
		log.Println(err)
		return
	}

	_, err = w.Write([]byte(timerJson))
	if err != nil {
		return
	}

}
