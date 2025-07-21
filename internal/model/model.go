package model

import "time"

type Timer struct {
	ID        int        `json:"id"`
	Title     string     `json:"title"`
	StartTime *time.Time `json:"start_time"`
	StopTime  *time.Time `json:"stop_time,omitempty"`
}

func (t *Timer) Duration() time.Duration {
	return t.StopTime.Sub(*t.StartTime)
}
