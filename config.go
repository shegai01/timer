package main

type Config struct {
	App struct {
		Port string `json:"port"`
	} `json:"app"`
	Dsn struct {
		Dsn string `json:"dsn"`
	} `json:"db"`
}
