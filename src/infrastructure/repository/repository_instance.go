package repository

import (
	"app/entity"

	"gorm.io/gorm"
)

type RepositoryInstance struct {
	DB *gorm.DB
}

func NewInstancePostgres(DB *gorm.DB) *RepositoryInstance {
	return &RepositoryInstance{DB: DB}
}

func (u *RepositoryInstance) GetAll() (instances []*entity.EntityInstance, err error) {
	err = u.DB.Find(&instances).Error

	return instances, err
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
