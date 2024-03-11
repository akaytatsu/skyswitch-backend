package usecase_calendar

import "app/entity"

type IRepositoryCalendar interface {
	GetFromID(id int) (*entity.EntityCalendar, error)
	GetAll(searchParams entity.SearchEntityCalendarParams) (response []entity.EntityCalendar, totalRegisters int64, err error)
	Create(*entity.EntityCalendar) error
	Update(*entity.EntityCalendar) error
	Delete(id int) error
}

type IUsecaseCalendar interface {
	Get(id int) (*entity.EntityCalendar, error)
	GetAll(searchParams entity.SearchEntityCalendarParams) (response []entity.EntityCalendar, totalRegisters int64, err error)
	Create(*entity.EntityCalendar) error
	Update(*entity.EntityCalendar) error
	Delete(id int) error
	CreateAllCalendarsJob() error
	ProccessCalendar(*entity.EntityCalendar) error
}
