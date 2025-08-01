package storage_test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/shegai01/timer/internal/model"
	"github.com/shegai01/timer/internal/storage"
	"github.com/shegai01/timer/internal/utils"
	"github.com/stretchr/testify/assert"
)

var (
	TestDB *storage.Storage
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

	TestDB = storage.NewStorage()
	if err := TestDB.Connect(databaseURI); err != nil {
		log.Fatalf("db.Connect(databaseURI) %v\n", err)
	}

	if err := TestDB.Drop(); err != nil {
		log.Fatalf("TestDB.Drop: %v", err)
	}
	os.Exit(m.Run())
}

func helperNewTimer(count int) *model.Timer {
	return &model.Timer{
		Title: utils.RandomTitle(count),
	}
}

func createRandomTimer(t *testing.T) *model.Timer {
	ctx := context.Background()
	newTimer := helperNewTimer(4)
	timer, err := TestDB.CreateTimer(ctx, newTimer.Title)
	assert.NoError(t, err)
	assert.Equal(t, newTimer.Title, timer.Title)
	assert.NotZero(t, timer.StartTime)
	return timer
}

func TestCreatet(t *testing.T) {
	ctx := context.Background()
	localtimer := createRandomTimer(t)
	gotTimer, err := TestDB.GetTimerByID(ctx, localtimer.ID)
	assert.NoError(t, err)
	assert.Equal(t, localtimer.ID, gotTimer.ID)

}

func TestShow(t *testing.T) {
	var arrayTimers []*model.Timer
	ctx := context.Background()
	for i := 0; i < 4; i++ {
		timer := createRandomTimer(t)
		arrayTimers = append(arrayTimers, timer)
	}

	timers, err := TestDB.ShowAllTimers(ctx)
	assert.NoError(t, err)
	assert.Equal(t, len(arrayTimers), len(timers))

	assert.NotNil(t, timers)
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	localtimer := createRandomTimer(t)
	// createTimer, err := TestDB.CreateTimer(ctx, localtimer.Title)
	// assert.NoError(t, err)
	gotTimer, err := TestDB.GetTimerByID(ctx, localtimer.ID)
	assert.NoError(t, err)
	assert.Equal(t, localtimer.ID, gotTimer.ID)

}

func TestDelete(t *testing.T) {
	ctx := context.Background()
	localtimer := createRandomTimer(t)
	err := TestDB.Delete(ctx, localtimer.ID)
	assert.NoError(t, err)
	_, err = TestDB.GetTimerByID(ctx, localtimer.ID)
	assert.Error(t, err)

}
func TestStop(t *testing.T) {
	ctx := context.Background()
	localtimer := createRandomTimer(t)
	gotTimer, err := TestDB.StopTimer(ctx, localtimer.ID)
	assert.NoError(t, err)
	assert.Equal(t, localtimer.StopTime, gotTimer.StopTime)

}
