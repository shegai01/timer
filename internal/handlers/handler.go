package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shegai01/timer/internal/model"
	"github.com/shegai01/timer/internal/storage"
)

func initContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

type TimerHandler struct {
	*mux.Router
	storage *storage.Storage
}

func NewTimerHandler(db *storage.Storage, router *mux.Router) *TimerHandler {
	return &TimerHandler{
		storage: db,
		Router:  router,
	}
}

func (h *TimerHandler) Create(w http.ResponseWriter, r *http.Request) {
	initContentType(w)
	var timer *model.Timer
	tittle := r.URL.Query().Get("title")
	timer, err := h.storage.CreateTimer(r.Context(), tittle)
	if err != nil {
		log.Printf("h.storage.CreateTimer: %v", err)

		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(timer); err != nil {
		http.Error(w, "encoding", http.StatusInternalServerError)
		log.Println("json.NewEncoder().Encode()", err)

		return
	}

}

func (h *TimerHandler) ShowAll(w http.ResponseWriter, r *http.Request) {
	var allTimers []*model.Timer
	initContentType(w)
	allTimers, err := h.storage.ShowAllTimers(r.Context())
	if err != nil {
		http.Error(w, "h.storage.ShowAllTimers", http.StatusInternalServerError)
		return
	}

	timers, err := json.MarshalIndent(allTimers, "", " ")
	if err != nil {
		http.Error(w, "MarshalIndent", http.StatusInternalServerError)
		log.Printf("json.MarshalIndent:%v\n", err)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(timers))
	if err != nil {
		log.Printf("w.Write:%v\n", err)

		return
	}
}

func (h *TimerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	initContentType(w)
	idTask := r.URL.Query().Get("id")
	strID, err := strconv.Atoi(idTask)
	if err != nil {
		http.Error(w, "incorrect input", http.StatusInternalServerError)
		log.Printf("strconv.Atoi: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := h.storage.Delete(r.Context(), strID); err != nil {
		http.Error(w, "h.storage.Delete", http.StatusInternalServerError)
		log.Printf("h.storage.Delete %v\n", err)

		return
	}

}

func (h *TimerHandler) GetbyID(w http.ResponseWriter, r *http.Request) {

	idTask := r.URL.Query().Get("id")
	strID, err := strconv.Atoi(idTask)
	if err != nil {
		log.Printf("strconv.Atoi(idTask):%v\n", err)

		return
	}

	initContentType(w)
	timer, err := h.storage.GetTimerByID(r.Context(), strID)
	if err != nil {
		http.Error(w, "h.storage.GetTimerByID", http.StatusInternalServerError)
		log.Printf("%v\n", err)

		return
	}
	timerbyTittle, err := json.MarshalIndent(timer, "", " ")
	if err != nil {
		log.Printf("json.MarshalIndent: %v\n", err)

		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(timerbyTittle))
	if err != nil {
		log.Printf("w.Write: %v\n", err)
		return
	}

}
func (h *TimerHandler) Stop(w http.ResponseWriter, r *http.Request) { //TODO: fix
	initContentType(w)
	idTask := r.URL.Query().Get("id")
	strID, err := strconv.Atoi(idTask)
	if err != nil {
		http.Error(w, "strconv.Atoi", http.StatusInternalServerError)
		log.Printf("strconv.Atoi: %v\n", err)
		return
	}

	timer, err := h.storage.StopTimer(r.Context(), strID)
	if err != nil {
		log.Printf("h.storage.StopTimer: %v\n", err)

		return
	}
	timerJson, err := json.MarshalIndent(timer, "", " ")
	if err != nil {
		log.Printf("json.MarshalIndent: %v\n", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write([]byte(timerJson))
	if err != nil {
		log.Printf("w.Write: %v\n", err)

		return
	}

}
