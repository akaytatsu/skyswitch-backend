package usecase_holiday

import (
	"app/entity"
	"time"
)

type IRepositoryHoliday interface {
	GetFromID(id int) (*entity.EntityHoliday, error)
	GetAll(searchParams entity.SearchEntityHolidayParams) (response []entity.EntityHoliday, totalRegisters int64, err error)
	Create(*entity.EntityHoliday) error
	Update(*entity.EntityHoliday) error
	Delete(id int) error
}

type IUsecaseHoliday interface {
	Get(id int) (*entity.EntityHoliday, error)
	GetAll(searchParams entity.SearchEntityHolidayParams) (response []entity.EntityHoliday, totalRegisters int64, err error)
	Create(*entity.EntityHoliday) error
	Update(*entity.EntityHoliday) error
	Delete(id int) error
	DateStringToTime(date string) (time.Time, error)
}
