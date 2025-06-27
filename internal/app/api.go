package app

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/shegai01/timer/internal/storage"
)

type API struct {
	config  *ConfigAPP
	storage *storage.Storage
	router  *mux.Router
}

func NewAPI(config *ConfigAPP) *API {
	return &API{
		config: config,
		router: mux.NewRouter(),
	}
}
func (api *API) Start() error {
	if err := api.configureStorage(); err != nil {
		return err
	}
	api.configureRouter()
	// slog.Info("app starting at port : %s", api.config.Port)
	return http.ListenAndServe(api.config.Port, api.router)
}
