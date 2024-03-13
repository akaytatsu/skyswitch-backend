package usecase_calendar

import "app/entity"

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_calendar.go -package=mocks app/usecase/calendar IRepositoryCalendar
type IRepositoryCalendar interface {
	GetFromID(id int) (*entity.EntityCalendar, error)
	GetAll(searchParams entity.SearchEntityCalendarParams) (response []entity.EntityCalendar, totalRegisters int64, err error)
	Create(*entity.EntityCalendar) error
	Update(*entity.EntityCalendar) error
	Delete(id int) error
}

//go:generate mockgen -destination=../../mocks/mock_usecase_calendar.go -package=mocks app/usecase/calendar IUsecaseCalendar
type IUsecaseCalendar interface {
	Get(id int) (*entity.EntityCalendar, error)
	GetAll(searchParams entity.SearchEntityCalendarParams) (response []entity.EntityCalendar, totalRegisters int64, err error)
	Create(*entity.EntityCalendar) error
	Update(*entity.EntityCalendar) error
	Delete(id int) error
	CreateAllCalendarsJob() error
	ProccessCalendar(*entity.EntityCalendar) error
}
