package model

import "time"

type Timer struct {
	ID        int           `json:"id"`
	Tittle    string        `json:"tittle"`
	StartTime time.Time     `json:"start_time"`
	StopTime  time.Time     `json:"stop_time"`
	Duration  time.Duration `json:"duration"`
}
