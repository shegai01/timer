package model

import "time"

type Timer struct {
	ID         int           `json:"id"`
	Tittle     string        `json:"tittle"`
	StartTime  time.Time     `json:"start_time"`
	FinishTime time.Time     `json:"finish_time"`
	Duration   time.Duration `json:"duration"`
	// Finished   bool          `json:"finish"`
}
