package usecase_calendar

import (
	"app/entity"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
)

type UsecaseCalendar struct {
	repo      IRepositoryCalendar
	scheduler *gocron.Scheduler
}

func NewService(repository IRepositoryCalendar, scheduler *gocron.Scheduler) *UsecaseCalendar {
	return &UsecaseCalendar{repo: repository, scheduler: scheduler}
}

func (u *UsecaseCalendar) Get(id int) (*entity.EntityCalendar, error) {
	return u.repo.GetFromID(id)
}

func (u *UsecaseCalendar) GetAll(searchParams entity.SearchEntityCalendarParams) (response []entity.EntityCalendar, totalRegisters int64, err error) {
	return u.repo.GetAll(searchParams)
}

func (u *UsecaseCalendar) Create(calendar *entity.EntityCalendar) error {
	err := u.repo.Create(calendar)

	if err != nil {
		return err
	}

	u.configureSchedules(calendar)

	return nil
}

func (u *UsecaseCalendar) Update(calendar *entity.EntityCalendar) error {
	err := u.repo.Update(calendar)

	if err != nil {
		return err
	}

	u.configureSchedules(calendar)

	return nil
}

func (u *UsecaseCalendar) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u *UsecaseCalendar) configureSchedules(calendar *entity.EntityCalendar) {

	u.cleanTags(calendar)

	if calendar.Active {
		var days []int = u.toDaysInt(calendar)

		for _, day := range days {

			var tag string = "calendar_" + strconv.Itoa(calendar.ID) + "_" + strconv.Itoa(day)

			weekday := time.Weekday(day)
			_, err := u.scheduler.Every(1).Weekday(weekday).At(calendar.ExecuteTime).Tag(tag).Do(func() {
				println("Calendar " + calendar.Name + " executed at " + time.Now().String())
			})

			if err != nil {
				println(err.Error())
			}
		}
	}

	// lista todos os jobs agendados
	// jobs := u.scheduler.Jobs()
	// for _, job := range jobs {
	// 	println(job.Tags, job.NextRun().GoString(), job.ScheduledTime().GoString())
	// }
}

func (u *UsecaseCalendar) cleanTags(calendar *entity.EntityCalendar) {

	var days []int = u.toDaysInt(calendar)

	for _, day := range days {
		var tag string = "calendar_" + strconv.Itoa(calendar.ID) + "_" + strconv.Itoa(day)

		u.scheduler.RemoveByTag(tag)
	}

}

func (u *UsecaseCalendar) toDaysInt(calendar *entity.EntityCalendar) (days []int) {

	if calendar.Sunday {
		days = append(days, 0)
	}

	if calendar.Monday {
		days = append(days, 1)
	}

	if calendar.Tuesday {
		days = append(days, 2)
	}

	if calendar.Wednesday {
		days = append(days, 3)
	}

	if calendar.Thursday {
		days = append(days, 4)
	}

	if calendar.Friday {
		days = append(days, 5)
	}

	if calendar.Saturday {
		days = append(days, 6)
	}

	return days
}
