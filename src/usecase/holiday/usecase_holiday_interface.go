package usecase_holiday

import (
	"app/entity"
	"time"
)

type IRepositoryHoliday interface {
	GetAll() ([]entity.EntityHoliday, error)
	GetByDate(date time.Time) (entity.EntityHoliday, error)
	CreateUpdate(name string, date time.Time) (entity.EntityHoliday, error)
}

type IUsecaseHoliday interface {
	GetAll() ([]entity.EntityHoliday, error)
	GetByDate(date time.Time) (entity.EntityHoliday, error)
	CreateUpdate(name string, date time.Time) (entity.EntityHoliday, error)
	DateStringToTime(date string) (time.Time, error)
}
