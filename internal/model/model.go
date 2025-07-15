package model

import "time"

type Timer struct {
	ID        int       `json:"id"`
	Tittle    string    `json:"tittle"`
	StartTime time.Time `json:"start_time"`
	StopTime  time.Time `json:"omiempty"`
}

func (t *Timer) DurationTimer() time.Duration {

	return t.StopTime.Sub(t.StartTime)
}
