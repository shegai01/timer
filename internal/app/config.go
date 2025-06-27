package api

import (
	"github.com/shegai01/timer/internal/storage"
)

type ConfigAPP struct {
	Port    string `json:"port"`
	Storage *storage.ConfigDB
}

func NewConfig() *ConfigAPP {
	return &ConfigAPP{
		Port:    "8080",
		Storage: storage.NewConfigStorage(),
	}
}
