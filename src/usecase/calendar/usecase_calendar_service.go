package usecase_calendar

import "app/entity"

type UsecaseCalendar struct {
	repo IRepositoryCalendar
}

func NewService(repository IRepositoryCalendar) *UsecaseCalendar {
	return &UsecaseCalendar{repo: repository}
}

func (u *UsecaseCalendar) Get(id int) (*entity.EntityCalendar, error) {
	return u.repo.GetFromID(id)
}

func (u *UsecaseCalendar) GetAll(searchParams entity.SearchEntityCalendarParams) (response []entity.EntityCalendar, totalRegisters int64, err error) {
	return u.repo.GetAll(searchParams)
}

func (u *UsecaseCalendar) Create(calendar *entity.EntityCalendar) error {
	return u.repo.Create(calendar)
}

func (u *UsecaseCalendar) Update(calendar *entity.EntityCalendar) error {
	return u.repo.Update(calendar)
}

func (u *UsecaseCalendar) Delete(id int) error {
	return u.repo.Delete(id)
}
