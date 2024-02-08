package repository

import (
	"app/entity"
	"time"

	"gorm.io/gorm"
)

type RepositoryHoliday struct {
	DB *gorm.DB
}

func NewHolidayPostgres(db *gorm.DB) *RepositoryHoliday {
	return &RepositoryHoliday{DB: db}
}

func (r *RepositoryHoliday) GetAll() ([]entity.EntityHoliday, error) {
	var holidays []entity.EntityHoliday
	err := r.DB.Find(&holidays).Error
	if err != nil {
		return nil, err
	}
	return holidays, nil
}

func (r *RepositoryHoliday) GetByDate(date time.Time) (entity.EntityHoliday, error) {
	var holiday entity.EntityHoliday
	err := r.DB.Where("date = ?", date).First(&holiday).Error
	if err != nil {
		return entity.EntityHoliday{}, err
	}
	return holiday, nil
}

func (r *RepositoryHoliday) CreateUpdate(name string, date time.Time) (entity.EntityHoliday, error) {
	holiday := entity.EntityHoliday{
		Name: name,
		Date: date,
	}

	err := r.DB.Save(&holiday).Error
	if err != nil {
		return entity.EntityHoliday{}, err
	}

	return holiday, nil
}
