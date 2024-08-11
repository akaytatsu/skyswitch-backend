package repository

import (
	"app/entity"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepositoryDbinstance struct {
	DB *gorm.DB
}

func NewDbinstancePostgres(DB *gorm.DB) *RepositoryDbinstance {
	return &RepositoryDbinstance{DB: DB}
}

func (r *RepositoryDbinstance) GetFromID(id int) (dbinstance *entity.EntityDbinstance, err error) {
	r.DB.First(&dbinstance, id)

	return
}

func (r *RepositoryDbinstance) GetAll(searchParams entity.SearchEntityDbinstanceParams) (response []entity.EntityDbinstance, totalRegisters int64, err error) {
	offset := (searchParams.Page) * searchParams.PageSize

	qry := r.DB.Model(entity.EntityDbinstance{})

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

func (r *RepositoryDbinstance) Create(dbinstance *entity.EntityDbinstance) (err error) {
	err = r.DB.Create(&dbinstance).Error

	return err
}

func (r *RepositoryDbinstance) Update(dbinstance *entity.EntityDbinstance) (err error) {
	_, err = r.GetFromID(int(dbinstance.ID))

	if err != nil {
		return err
	}

	err = r.DB.Save(&dbinstance).Error

	return err
}

func (r *RepositoryDbinstance) Delete(id int) (err error) {
	err = r.DB.Delete(&entity.EntityDbinstance{}, id).Error

	return err
}

func (u *RepositoryDbinstance) FromCalendar(calendarID int) (response []entity.EntityDbinstance, err error) {

	err = u.DB.Debug().Model(&entity.EntityDbinstance{}).Preload("CloudAccount").
		Joins("JOIN entity_instance_calendars ON entity_instance_calendars.entity_dbinstance_id = entity_dbinstances.id").
		Where("entity_instance_calendars.entity_calendar_id = ?", calendarID).
		Find(&response).Error

	return response, err
}

func (u *RepositoryDbinstance) GetByInstanceID(instanceID string) (instance *entity.EntityDbinstance, err error) {
	err = u.DB.Where("db_instance_id = ?", instanceID).First(&instance).Error

	return instance, err
}

func (u *RepositoryDbinstance) ActiveDeactiveInstance(id int64, status bool) (instance *entity.EntityDbinstance, err error) {
	instance, err = u.GetFromID(int(id))

	if err != nil {
		return nil, err
	}

	return instance, nil
}
