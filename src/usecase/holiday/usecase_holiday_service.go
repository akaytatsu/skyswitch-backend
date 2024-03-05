package usecase_holiday

import (
	"app/entity"
	"time"
)

type UsecaseHoliday struct {
	repo IRepositoryHoliday
}

func NewService(repository IRepositoryHoliday) *UsecaseHoliday {
	return &UsecaseHoliday{repo: repository}
}

func (u *UsecaseHoliday) Get(id int) (*entity.EntityHoliday, error) {
	return u.repo.GetFromID(id)
}

func (u *UsecaseHoliday) GetAll(searchParams entity.SearchEntityHolidayParams) (response []entity.EntityHoliday, totalRegisters int64, err error) {
	return u.repo.GetAll(searchParams)
}

func (u *UsecaseHoliday) Create(holiday *entity.EntityHoliday) error {
	return u.repo.Create(holiday)
}

func (u *UsecaseHoliday) Update(holiday *entity.EntityHoliday) error {
	return u.repo.Update(holiday)
}

func (u *UsecaseHoliday) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u *UsecaseHoliday) DateStringToTime(date string) (time.Time, error) {
	return time.Parse("2006-01-02", date)
}
