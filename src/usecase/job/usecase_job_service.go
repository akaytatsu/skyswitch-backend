package usecase_job

import (
	"app/cron"
	"app/entity"
)

type UsecaseJob struct {
}

func NewService() *UsecaseJob {
	return &UsecaseJob{}
}

func (u *UsecaseJob) GetAll() (response []entity.EntityJob, err error) {
	schedule := cron.Scheduler

	// busca todos os job do cron
	entries := schedule.Jobs()
	for _, entry := range entries {

		var tagID string
		var errorData string

		if entry.Tags() != nil && len(entry.Tags()) > 0 {
			tagID = entry.Tags()[0]
		}

		if entry.Error() != nil {
			errorData = entry.Error().Error()
		}

		job := entity.EntityJob{
			ID:            tagID,
			IsRunning:     entry.IsRunning(),
			LastRun:       entry.LastRun().Local(),
			NextRun:       entry.NextRun().Local(),
			Count:         entry.RunCount(),
			Error:         errorData,
			ScheduledTime: entry.ScheduledTime().Local(),
		}
		response = append(response, job)
	}

	return response, nil
}
