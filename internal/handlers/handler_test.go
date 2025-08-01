package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/joho/godotenv"
	"github.com/shegai01/timer/internal/model"
	"github.com/shegai01/timer/internal/storage"
	"github.com/stretchr/testify/assert"
)

var (
	TestTimerHandler *TimerHandler
	exampleTimer     *model.Timer
)

func TestMain(m *testing.M) {

	err := godotenv.Load("/home/alex/studyGolang/mentor/timer/.env")
	if err != nil {
		log.Fatalf("godotenv.Load: %v\n", err)
	}

	databaseURI := os.Getenv("APP_DATABASE_URI")
	if databaseURI == "" {
		log.Fatalf("env is missing: APP_DATABASE_URI %v\n", err)
	}
	db := storage.NewStorage()
	db.Connect(databaseURI)
	handlTest := NewTimerHandler(db)
	TestTimerHandler = handlTest
	os.Exit(m.Run())
}

func TestCreate(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/create?title=test", nil)
	w := httptest.NewRecorder()
	TestTimerHandler.Create(w, req)
	res := w.Result()
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	err = json.Unmarshal(data, &exampleTimer)
	assert.NoError(t, err)
	assert.NotNil(t, exampleTimer)
}

func TestDelete(t *testing.T) {
	reqCreate := httptest.NewRequest(http.MethodDelete, "/create?title=test", nil)
	wCreate := httptest.NewRecorder()
	TestTimerHandler.Create(wCreate, reqCreate)
	resCreate := wCreate.Result()
	defer resCreate.Body.Close()

	dataTimer, err := io.ReadAll(resCreate.Body)
	assert.NoError(t, err)
	err = json.Unmarshal(dataTimer, &exampleTimer)
	assert.NoError(t, err)

	reqDelete := httptest.NewRequest(http.MethodDelete, "/delete?id="+strconv.Itoa(exampleTimer.ID), nil)
	wdelete := httptest.NewRecorder()
	TestTimerHandler.Delete(wdelete, reqDelete)
	resDelete := wdelete.Result()
	data, err := io.ReadAll(resDelete.Body)
	assert.NoError(t, err)
	assert.Len(t, data, 0)
}

func TestGET(t *testing.T) {
	reqCreate := httptest.NewRequest(http.MethodGet, "/create?title=test", nil)
	wCreate := httptest.NewRecorder()
	TestTimerHandler.Create(wCreate, reqCreate)
	resCreate := wCreate.Result()
	defer resCreate.Body.Close()
	dataTimer, err := io.ReadAll(resCreate.Body)
	assert.NoError(t, err)
	err = json.Unmarshal(dataTimer, &exampleTimer)
	assert.NoError(t, err)

	reqGet := httptest.NewRequest(http.MethodPost, "/delete?id="+strconv.Itoa(exampleTimer.ID), nil)
	wGet := httptest.NewRecorder()
	TestTimerHandler.GetbyID(wGet, reqGet)
	resGet := wGet.Result()
	data, err := io.ReadAll(resGet.Body)
	assert.NoError(t, err)

	var gotTimer model.Timer
	err = json.Unmarshal(data, &gotTimer)
	assert.NoError(t, err)

	assert.Equal(t, exampleTimer.ID, gotTimer.ID)
}

func TestShow(t *testing.T) {
	reqCreate := httptest.NewRequest(http.MethodGet, "/create?title=test", nil)
	wCreate := httptest.NewRecorder()
	TestTimerHandler.Create(wCreate, reqCreate)
	resCreate := wCreate.Result()
	defer resCreate.Body.Close()

	reqShow := httptest.NewRequest(http.MethodPost, "/show", nil)
	wShow := httptest.NewRecorder()
	TestTimerHandler.ShowAll(wShow, reqShow)
	resShow := wShow.Result()

	var gotTimers []*model.Timer
	data, err := io.ReadAll(resShow.Body)
	assert.NoError(t, err)

	err = json.Unmarshal(data, &gotTimers)
	assert.NoError(t, err)
	assert.NotNil(t, gotTimers)

}
