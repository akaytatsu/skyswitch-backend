package repository

import (
	"app/entity"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepositoryLog struct {
	DB *gorm.DB
}

func NewLogPostgres(DB *gorm.DB) *RepositoryLog {
	return &RepositoryLog{DB: DB}
}

func (r *RepositoryLog) GetFromID(id int) (log *entity.EntityLog, err error) {
	r.DB.First(&log, id)

	return
}

func (r *RepositoryLog) GetAll(searchParams entity.SearchEntityLogParams) (response []entity.EntityLog, totalRegisters int64, err error) {
	offset := (searchParams.Page) * searchParams.PageSize

	qry := r.DB.Model(entity.EntityLog{})

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

	if searchParams.OrderBy == "" {
		searchParams.OrderBy = "id"
	}

	if searchParams.SortOrder == "" {
		searchParams.SortOrder = "asc"
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

func (r *RepositoryLog) Create(log *entity.EntityLog) (err error) {
	err = r.DB.Create(&log).Error

	return
}

func (r *RepositoryLog) Update(log *entity.EntityLog) (err error) {
	_, err = r.GetFromID(int(log.ID))

	if err != nil {
		return
	}

	err = r.DB.Save(&log).Error

	return
}

func (r *RepositoryLog) Delete(id int) (err error) {
	err = r.DB.Delete(&entity.EntityLog{}, id).Error

	return
}
