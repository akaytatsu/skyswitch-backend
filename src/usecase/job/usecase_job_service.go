package usecase_job

import (
	"app/cron"
	"app/entity"
	usecase_calendar "app/usecase/calendar"
	"strconv"
	"strings"
)

type UsecaseJob struct {
	usecaseCalendar usecase_calendar.IUsecaseCalendar
}

func NewService(usecaseCalendar usecase_calendar.IUsecaseCalendar) *UsecaseJob {
	return &UsecaseJob{
		usecaseCalendar: usecaseCalendar,
	}
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

		// pega o id do calendario, se o job tiver tag, a tag é calendar_x_y, sendo o X o id
		if tagID != "" {
			if strings.Split(tagID, "_")[0] == "calendar" && len(strings.Split(tagID, "_")) > 1 {
				id, _ := strconv.Atoi(strings.Split(tagID, "_")[1])
				calendar, _ := u.usecaseCalendar.Get(id)
				job.Calendar = *calendar
			}
		}

		response = append(response, job)
	}

	return response, nil
}
