package repository

import (
	"app/entity"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepositoryHoliday struct {
	DB *gorm.DB
}

func NewHolidayPostgres(DB *gorm.DB) *RepositoryHoliday {
	return &RepositoryHoliday{DB: DB}
}

func (r *RepositoryHoliday) GetFromID(id int) (holiday *entity.EntityHoliday, err error) {
	r.DB.First(&holiday, id)

	return
}

func (r *RepositoryHoliday) GetAll(searchParams entity.SearchEntityHolidayParams) (response []entity.EntityHoliday, totalRegisters int64, err error) {
	offset := (searchParams.Page) * searchParams.PageSize

	qry := r.DB.Model(entity.EntityHoliday{})

	if gin.IsDebugging() {
		qry = qry.Debug()
	}

	if searchParams.Q != "" {
		qry = qry.Where("name LIKE ?", "%"+searchParams.Q+"%")
	}

	if searchParams.CreatedAt != "" {
		dates := strings.Split(searchParams.CreatedAt, ",")
		if len(dates) == 2 {
			_, err1 := time.Parse("2006-01-02", dates[0])
			_, err2 := time.Parse("2006-01-02", dates[1])
			if err1 == nil && err2 == nil {
				qry = qry.Where("created_at BETWEEN ? AND ?", dates[0], dates[1])
			}
		}
	}

	err = qry.Count(&totalRegisters).Error

	if err != nil {
		return nil, 0, err
	}

	qry = qry.Order(searchParams.OrderBy + " " + searchParams.SortOrder).
		Offset(offset).
		Limit(searchParams.PageSize)

	err = qry.Find(&response).Error
	if err != nil {
		return nil, 0, err
	}

	return response, totalRegisters, nil
}

func (r *RepositoryHoliday) Create(holiday *entity.EntityHoliday) (err error) {
	err = r.DB.Create(&holiday).Error

	return
}

func (r *RepositoryHoliday) Update(holiday *entity.EntityHoliday) (err error) {
	_, err = r.GetFromID(int(holiday.ID))

	if err != nil {
		return
	}

	err = r.DB.Save(&holiday).Error

	return
}

func (r *RepositoryHoliday) Delete(id int) (err error) {
	err = r.DB.Delete(&entity.EntityHoliday{}, id).Error

	return
}
