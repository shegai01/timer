package handlers

import (
	"encoding/json"
	"errors"
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
	initContentType(w)
	tittle := r.URL.Query().Get("title")
	timer, err := h.storage.CreateTimer(r.Context(), tittle)
	if err != nil {
		log.Printf("create: %v", err)

		return
	}
	log.Println("timer created")

	if err := json.NewEncoder(w).Encode(timer); err != nil {
		http.Error(w, "encoding", http.StatusBadRequest)
		log.Println("json encoder in 'CreateTimer' failed", err)

		return
	}

}

func (h *Timerhandler) ShowAll(w http.ResponseWriter, r *http.Request) {
	var allTimers []*model.Timer
	initContentType(w)
	allTimers, err := h.storage.ShowAllTimers(r.Context())
	if err != nil {
		http.Error(w, "ShowAll failed", http.StatusInternalServerError)
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
	if err := h.storage.Delete(r.Context(), strID); err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Printf("h.storage.Delete %s\n", err)

		return
	}
	log.Println("deleted")
	w.WriteHeader(http.StatusOK)
}
func (h *Timerhandler) GetbyID(w http.ResponseWriter, r *http.Request) {

	idTask := r.URL.Query().Get("id")
	strID, err := strconv.Atoi(idTask)
	if err != nil {
		return
	}
	initContentType(w)

	timer, err := h.storage.GetTimerbyID(r.Context(), strID)
	if err != nil {
		http.Error(w, "get function failed", http.StatusInternalServerError)
		log.Println("get in function storage failed", err)

		return
	}
	timerbyTittle, err := json.MarshalIndent(timer, "", " ")
	if err != nil {
		log.Println("marshaling failed", err)
		errors.Is(err, &json.MarshalerError{})
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
		http.Error(w, "stop failed", http.StatusInternalServerError)
		log.Println("stoptimer failed in storage", err)
		return
	}
	updatetimer, err := h.storage.GetTimerbyID(r.Context(), strID)
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
