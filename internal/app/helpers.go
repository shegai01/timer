package app

import (
	"log/slog"
	"net/http"

	"github.com/shegai01/timer/internal/storage"
)

// starting router
func (a *API) configureRouter() {
	a.router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
}

// starting db
func (a *API) configureStorage() error {
	storage := storage.NewStorage(a.config.Storage)
	if err := storage.ConnectDB(); err != nil {
		slog.Error("connection database failed")
		return err
	}

	// slog.Info("configureated storage")
	return nil

}
