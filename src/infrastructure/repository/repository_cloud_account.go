package repository

import (
	"app/entity"

	"gorm.io/gorm"
)

type RepositoryCloudAccount struct {
	DB *gorm.DB
}

func NewCloudAccountPostgres(db *gorm.DB) *RepositoryCloudAccount {
	return &RepositoryCloudAccount{DB: db}
}

func (u *RepositoryCloudAccount) GetAll() (cloudAccounts []*entity.EntityCloudAccount, err error) {
	err = u.DB.Find(&cloudAccounts).Error

	return cloudAccounts, err
}

func (u *RepositoryCloudAccount) GetByID(id int64) (cloudAccount *entity.EntityCloudAccount, err error) {
	err = u.DB.Where("id = ?", id).First(&cloudAccount).Error

	return cloudAccount, err
}

func (u *RepositoryCloudAccount) CreateCloudAccount(cloudAccount *entity.EntityCloudAccount) error {
	return u.DB.Create(&cloudAccount).Error
}

func (u *RepositoryCloudAccount) UpdateCloudAccount(cloudAccount *entity.EntityCloudAccount) error {

	_, err := u.GetByID(cloudAccount.ID)

	if err != nil {
		return err
	}

	return u.DB.Save(&cloudAccount).Error
}

func (u *RepositoryCloudAccount) DeleteCloudAccount(cloudAccount *entity.EntityCloudAccount) error {

	_, err := u.GetByID(cloudAccount.ID)

	if err != nil {
		return err
	}

	return u.DB.Delete(&cloudAccount).Error
}

func (u *RepositoryCloudAccount) ActiveDeactiveCloudAccount(id int64, status bool) (cloudAccount *entity.EntityCloudAccount, err error) {

	cloudAccount, err = u.GetByID(id)

	if err != nil {
		return nil, err
	}

	cloudAccount.Active = status

	return cloudAccount, u.DB.Save(&cloudAccount).Error
}
