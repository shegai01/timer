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

func NewTimerHandler(db *storage.Storage) *TimerHandler {
	h := &TimerHandler{
		storage: db,
		Router:  mux.NewRouter(),
	}

	h.HandleFunc("/create", h.Create)
	h.HandleFunc("/show", h.ShowAll)
	h.HandleFunc("/get", h.GetbyID)
	h.HandleFunc("/delete", h.Delete)
	h.HandleFunc("/stop", h.Stop)

	return h
}

func Error(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func (h *TimerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var timer *model.Timer
	title := r.URL.Query().Get("title")
	timer, err := h.storage.CreateTimer(r.Context(), title)
	if err != nil {
		log.Printf("h.storage.CreateTimer: %v\n", err)
		Error(w, http.StatusInternalServerError)
		return
	}

	initContentType(w)
	if err := json.NewEncoder(w).Encode(timer); err != nil {
		log.Printf("json.NewEncoder.Encode: %v\n", err)
		Error(w, http.StatusInternalServerError)
		return
	}
}

func (h *TimerHandler) ShowAll(w http.ResponseWriter, r *http.Request) {
	var allTimers []*model.Timer
	allTimers, err := h.storage.ShowAllTimers(r.Context())
	if err != nil {
		log.Printf("h.storage.ShowAllTimers: %v\n", err)
		Error(w, http.StatusInternalServerError)
		return
	}

	initContentType(w)
	if err := json.NewEncoder(w).Encode(allTimers); err != nil {
		log.Printf("json.NewEncoder.Encode: %v\n", err)
		Error(w, http.StatusInternalServerError)
		return
	}

}

func (h *TimerHandler) Delete(w http.ResponseWriter, r *http.Request) {
	idTask := r.URL.Query().Get("id")
	if idTask == "" {
		log.Println("id query parameter is not set")
		Error(w, http.StatusBadRequest)
		return
	}

	strID, err := strconv.Atoi(idTask)
	if err != nil {
		log.Printf("strconv.Atoi: %v\n", err)
		Error(w, http.StatusBadRequest)
		return
	}

	if err := h.storage.Delete(r.Context(), strID); err != nil {
		log.Printf("h.storage.Delete %v\n", err)
		Error(w, http.StatusInternalServerError)
		return
	}
}

func (h *TimerHandler) GetbyID(w http.ResponseWriter, r *http.Request) {
	idTask := r.URL.Query().Get("id")
	if idTask == "" {
		log.Println("id query parameter is not set")
		Error(w, http.StatusBadRequest)
		return
	}

	strID, err := strconv.Atoi(idTask)
	if err != nil {
		log.Printf("strconv.Atoi: %v\n", err)
		Error(w, http.StatusBadRequest)
		return
	}

	timer, err := h.storage.GetTimerByID(r.Context(), strID)
	if err != nil {
		log.Printf("h.storage.GetTimerByID: %v\n", err)
		Error(w, http.StatusInternalServerError)
		return
	}

	initContentType(w)
	if err := json.NewEncoder(w).Encode(timer); err != nil {
		log.Printf("json.NewEncoder.Encode: %v\n", err)
		Error(w, http.StatusInternalServerError)
		return
	}
}

func (h *TimerHandler) Stop(w http.ResponseWriter, r *http.Request) { //TODO: fix
	initContentType(w)
	idTask := r.URL.Query().Get("id")
	if idTask == "" {
		log.Println("id query parameter is not set")
		Error(w, http.StatusBadRequest)
		return
	}

	strID, err := strconv.Atoi(idTask)
	if err != nil {
		log.Printf("strconv.Atoi: %v\n", err)
		Error(w, http.StatusBadRequest)
		return
	}

	timer, err := h.storage.StopTimer(r.Context(), strID)
	if err != nil {
		log.Printf("h.storage.StopTimer: %v\n", err)
		Error(w, http.StatusInternalServerError)
		return
	}

	initContentType(w)
	if err := json.NewEncoder(w).Encode(timer); err != nil {
		log.Printf("json.NewEncoder.Encode: %v\n", err)
		Error(w, http.StatusInternalServerError)
		return
	}
}
