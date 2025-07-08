package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

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
	// fmt.Println("create timer successfully")
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
		log.Println("not saved")
		return
	}
}

func (h *Timerhandler) DeletebyID(w http.ResponseWriter, r *http.Request) {
	// idTask := r.URL.Query().Get("id")
	tittleTask := r.URL.Query().Get("tittle")

	initContentType(w)
	// w.WriteHeader(http.StatusOK)
	// strID, err := strconv.Atoi(idTask)
	// if err != nil {
	// 	http.Error(w, "incorrect input", http.StatusOK)
	// 	log.Println("incorrect id", err)
	// 	return
	// }
	if err := h.storage.DeletebyID(tittleTask); err != nil {
		http.Error(w, "", http.StatusBadRequest)
		log.Println("can't delete timer")

		return
	}
	timerDeleted, err := json.MarshalIndent(tittleTask, "", " ")
	if err != nil {
		log.Println("not getting for marshalling")
		return
	}
	err = os.WriteFile("deleted.json", timerDeleted, 0666)
	if err != nil {
		log.Println("save in file failed")
		return
	}
	log.Println("deleted timers will be save in file timer.json")

}
func (h *Timerhandler) GetTimerbyID(w http.ResponseWriter, r *http.Request) {
	tittleTask := r.URL.Query().Get("tittle")
	// w.Header().Set(contentType, applicationsJson)
	// id := r.URL.Query()
	// w.WriteHeader(http.StatusOK)
	initContentType(w)

	timer, err := h.storage.GetTimerbyTittle(tittleTask)
	if err != nil {
		http.Error(w, "get function failed", http.StatusBadRequest)
		log.Println(err)

		return
	}
	timerbyTittle, err := json.MarshalIndent(timer, "", " ")
	if err != nil {
		http.Error(w, "", http.StatusBadRequest)
		log.Println("marshaling failed", err)
		return
	}

	err = os.WriteFile("gettimer.json", timerbyTittle, 0666)
	if err != nil {
		log.Println("not saved in file gettimer.json")
		return
	}
	w.Write([]byte(timerbyTittle))
}
func (h *Timerhandler) UpdateTimer(w http.ResponseWriter, r *http.Request) {
	tittle := r.URL.Query().Get("tittle")

	timer, err := h.storage.UpdateTimer(tittle, time.Now())
	if err != nil {
		http.Error(w, "update at handler failed", http.StatusBadRequest)
		log.Println(err)
		return
	}
	timerJson, err := json.MarshalIndent(timer, "", " ")
	if err != nil {
		log.Println(err)
		return
	}
	err = os.WriteFile("update.json", timerJson, 0666)
	if err != nil {
		log.Println("save in file update.json failed")
		return
	}

}
