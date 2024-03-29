package usecase_holiday

import (
	"app/entity"
	"time"
)

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_holiday.go -package=mocks app/usecase/holiday IRepositoryHoliday
type IRepositoryHoliday interface {
	GetFromID(id int) (*entity.EntityHoliday, error)
	CheckDateExists(date time.Time) (bool, error)
	GetAll(searchParams entity.SearchEntityHolidayParams) (response []entity.EntityHoliday, totalRegisters int64, err error)
	Create(*entity.EntityHoliday) error
	Update(*entity.EntityHoliday) error
	Delete(id int) error
}

//go:generate mockgen -destination=../../mocks/mock_usecase_holiday.go -package=mocks app/usecase/holiday IUsecaseHoliday
type IUsecaseHoliday interface {
	Get(id int) (*entity.EntityHoliday, error)
	GetAll(searchParams entity.SearchEntityHolidayParams) (response []entity.EntityHoliday, totalRegisters int64, err error)
	Create(*entity.EntityHoliday) error
	Update(*entity.EntityHoliday) error
	Delete(id int) error
	DateStringToTime(date string) (time.Time, error)
	IsHoliday(date time.Time) (bool, error)
}
