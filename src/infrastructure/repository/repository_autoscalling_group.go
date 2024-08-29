package repository

import (
	"app/entity"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepositoryAutoScalingGroup struct {
	DB *gorm.DB
}

func NewAutoScalingGroupPostgres(DB *gorm.DB) *RepositoryAutoScalingGroup {
	return &RepositoryAutoScalingGroup{DB: DB}
}

func (r *RepositoryAutoScalingGroup) GetFromID(id int) (autoScalingGroup *entity.EntityAutoScalingGroup, err error) {
	r.DB.Model(entity.EntityAutoScalingGroup{}).Preload("CloudAccount").First(&autoScalingGroup, id)

	return
}

func (r *RepositoryAutoScalingGroup) GetAll(searchParams entity.SearchEntityAutoScalingGroupParams) (response []entity.EntityAutoScalingGroup, totalRegisters int64, err error) {
	offset := (searchParams.Page) * searchParams.PageSize

	qry := r.DB.Model(entity.EntityAutoScalingGroup{}).Preload("Calendars").Preload("CloudAccount")

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

	qry = qry.Order(searchParams.OrderBy + " " + searchParams.SortOrder)

	err = qry.Offset(offset).Limit(searchParams.PageSize).Find(&response).Error

	return
}

func (r *RepositoryAutoScalingGroup) Create(autoScalingGroup *entity.EntityAutoScalingGroup) error {
	return r.DB.Create(&autoScalingGroup).Error
}

func (r *RepositoryAutoScalingGroup) Update(autoScalingGroup *entity.EntityAutoScalingGroup, updateCalendars bool) error {
	_, err := r.GetFromID(autoScalingGroup.ID)

	if err != nil {
		return err
	}

	if updateCalendars {
		calendarsAux := autoScalingGroup.Calendars
		r.DB.Model(autoScalingGroup).Association("Calendars").Clear()
		autoScalingGroup.Calendars = calendarsAux
	}

	return r.DB.Save(&autoScalingGroup).Error
}

func (r *RepositoryAutoScalingGroup) Delete(id int) error {
	return r.DB.Delete(&entity.EntityAutoScalingGroup{}, id).Error
}

func (r *RepositoryAutoScalingGroup) FromCalendar(calendarID int) (response []entity.EntityAutoScalingGroup, err error) {
	err = r.DB.Model(&entity.EntityAutoScalingGroup{}).
		Joins("JOIN entity_autoscalling_groups_calendars ON entity_autoscalling_groups_calendars.entity_auto_scaling_group_id = entity_auto_scaling_groups.id").
		Where("entity_autoscalling_groups_calendars.entity_calendar_id = ?", calendarID).
		Preload("CloudAccount").
		Find(&response).Error

	return response, err
}

func (r *RepositoryAutoScalingGroup) GetByID(autoScallingGroupID string) (autoScallingGroup *entity.EntityAutoScalingGroup, err error) {
	err = r.DB.Where("auto_scaling_group_id = ?", autoScallingGroupID).Preload("CloudAccount").First(&autoScallingGroup).Error

	return
}

func (r *RepositoryAutoScalingGroup) ActiveDeactive(id int64, status bool) (autoScallingGroup *entity.EntityAutoScalingGroup, err error) {
	autoScallingGroup, err = r.GetFromID(int(id))

	if err != nil {
		return
	}

	autoScallingGroup.Active = status

	err = r.DB.Save(autoScallingGroup).Error

	return
}
