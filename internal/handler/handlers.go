package internal

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/shegai01/timer/internal/model"
	"github.com/shegai01/timer/internal/storage"
)

type TrackerHandler struct {
	repo storage.TimeTracker
}

func NewTrackerHandler(repo storage.TimeTracker) *TrackerHandler {
	return &TrackerHandler{
		repo: repo,
	}
}
func initHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func (h *TrackerHandler) StartTimer(w http.ResponseWriter, r *http.Request) {
	nameTask := r.URL.Query().Get("name")
	log.Printf("getting url params by query %s\n", nameTask)
	timer, err := h.repo.CreateTimer()
	if err != nil {
		http.Error(w, "creating failed", http.StatusBadRequest)
		log.Println(err)

		return
	}

	initHeader(w)

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(timer); err != nil {
		http.Error(w, "encoding failed", http.StatusBadRequest)
		log.Println(err)

		return
	}
}
func (h *TrackerHandler) AllTimers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	initHeader(w)
	var arrTimer []model.Timer
	arrTimer, err := h.repo.ShowAllTimers()
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
}
func (h *TrackerHandler) StopTimer(w http.ResponseWriter, r *http.Request) {

}
func (h *TrackerHandler) GetTimerbyID(w http.ResponseWriter, r *http.Request) {
	idTask := r.URL.Query().Get("id")
	// w.Header().Set(contentType, applicationsJson)
	// id := r.URL.Query()
	initHeader(w)
	strID, err := strconv.Atoi(idTask)
	if err != nil {
		log.Println("incorrect id", err)
	}

	if err = DeletebyID(strID); err != nil {
		http.Error(w, "delete func failed", http.StatusBadRequest)
		log.Println(err)

		return
	}
	w.WriteHeader(http.StatusAccepted)
}
