package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type EntityJob struct {
	ID            string    `json:"id"`
	IsRunning     bool      `json:"is_running"`
	LastRun       time.Time `json:"last_run"`
	NextRun       time.Time `json:"next_run"`
	Count         int       `json:"count"`
	Error         string    `json:"error"`
	ScheduledTime time.Time `json:"scheduled_time"`
}

func NewEntityJob(entityJobParam EntityJob) (*EntityJob, error) {
	u := &EntityJob{
		ID:            entityJobParam.ID,
		IsRunning:     entityJobParam.IsRunning,
		ScheduledTime: entityJobParam.ScheduledTime,
		Count:         entityJobParam.Count,
		LastRun:       entityJobParam.LastRun,
		NextRun:       entityJobParam.NextRun,
		Error:         entityJobParam.Error,
	}

	return u, nil
}

func (u *EntityJob) Validate() error {
	return validator.New().Struct(u)
}
