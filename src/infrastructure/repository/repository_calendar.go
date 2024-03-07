package repository

import (
	"app/entity"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepositoryCalendar struct {
	DB *gorm.DB
}

func NewCalendarPostgres(DB *gorm.DB) *RepositoryCalendar {
	return &RepositoryCalendar{DB: DB}
}

func (r *RepositoryCalendar) GetFromID(id int) (calendar *entity.EntityCalendar, err error) {
	r.DB.First(&calendar, id)

	return
}

func (r *RepositoryCalendar) GetAll(searchParams entity.SearchEntityCalendarParams) (response []entity.EntityCalendar, totalRegisters int64, err error) {
	offset := (searchParams.Page) * searchParams.PageSize

	qry := r.DB.Model(entity.EntityCalendar{})

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

	// qry = qry.Order(searchParams.OrderBy + " " + searchParams.SortOrder).
	qry = qry.Order("id desc").
		Offset(offset).
		Limit(searchParams.PageSize)

	err = qry.Find(&response).Error
	if err != nil {
		return nil, 0, err
	}

	return response, totalRegisters, nil
}

func (r *RepositoryCalendar) Create(calendar *entity.EntityCalendar) (err error) {
	err = r.DB.Create(&calendar).Error

	return
}

func (r *RepositoryCalendar) Update(calendar *entity.EntityCalendar) (err error) {
	_, err = r.GetFromID(int(calendar.ID))

	if err != nil {
		return
	}

	err = r.DB.Save(&calendar).Error

	return
}

func (r *RepositoryCalendar) Delete(id int) (err error) {
	err = r.DB.Delete(&entity.EntityCalendar{}, id).Error

	return
}
