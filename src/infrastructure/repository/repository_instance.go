package repository

import (
	"app/entity"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepositoryInstance struct {
	DB *gorm.DB
}

func NewInstancePostgres(DB *gorm.DB) *RepositoryInstance {
	return &RepositoryInstance{DB: DB}
}

func (u *RepositoryInstance) GetAll(searchParams entity.SearchEntityInstanceParams) (response []entity.EntityInstance, totalRegisters int64, err error) {
	offset := (searchParams.Page) * searchParams.PageSize

	qry := u.DB.Model(entity.EntityInstance{})

	if gin.IsDebugging() {
		qry = qry.Debug()
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

func (u *RepositoryInstance) GetByID(id int64) (instance *entity.EntityInstance, err error) {
	err = u.DB.Where("id = ?", id).First(&instance).Error

	return instance, err
}

func (u *RepositoryInstance) GetByInstanceID(instanceID string) (instance *entity.EntityInstance, err error) {
	err = u.DB.Where("instance_id = ?", instanceID).First(&instance).Error

	return instance, err
}

func (u *RepositoryInstance) CreateInstance(instance *entity.EntityInstance) error {
	return u.DB.Create(&instance).Error
}

func (u *RepositoryInstance) UpdateInstance(instance *entity.EntityInstance) error {

	_, err := u.GetByID(instance.ID)

	if err != nil {
		return err
	}

	return u.DB.Save(&instance).Error
}

func (u *RepositoryInstance) DeleteInstance(instance *entity.EntityInstance) error {

	_, err := u.GetByID(instance.ID)

	if err != nil {
		return err
	}

	return u.DB.Delete(&instance).Error
}

func (u *RepositoryInstance) ActiveDeactiveInstance(id int64, status bool) (instance *entity.EntityInstance, err error) {

	instance, err = u.GetByID(id)

	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (u *RepositoryInstance) CreateOrUpdateInstance(instance *entity.EntityInstance) error {
	return u.DB.Save(&instance).Error
}
