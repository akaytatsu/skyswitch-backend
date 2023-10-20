package repository

import (
	"app/entity"

	"gorm.io/gorm"
)

type RepositoryCloudAccount struct {
	db *gorm.DB
}

func NewCloudAccountPostgres(db *gorm.DB) *RepositoryCloudAccount {
	return &RepositoryCloudAccount{db: db}
}

func (u *RepositoryCloudAccount) GetAll() (cloudAccounts []*entity.EntityCloudAccount, err error) {
	err = u.db.Find(&cloudAccounts).Error

	return cloudAccounts, err
}

func (u *RepositoryCloudAccount) GetByID(id int64) (cloudAccount *entity.EntityCloudAccount, err error) {
	err = u.db.Where("id = ?", id).First(&cloudAccount).Error

	return cloudAccount, err
}

func (u *RepositoryCloudAccount) CreateCloudAccount(cloudAccount *entity.EntityCloudAccount) error {
	return u.db.Create(&cloudAccount).Error
}

func (u *RepositoryCloudAccount) UpdateCloudAccount(cloudAccount *entity.EntityCloudAccount) error {

	_, err := u.GetByID(cloudAccount.ID)

	if err != nil {
		return err
	}

	return u.db.Save(&cloudAccount).Error
}

func (u *RepositoryCloudAccount) DeleteCloudAccount(cloudAccount *entity.EntityCloudAccount) error {

	_, err := u.GetByID(cloudAccount.ID)

	if err != nil {
		return err
	}

	return u.db.Delete(&cloudAccount).Error
}

func (u *RepositoryCloudAccount) ActiveDeactiveCloudAccount(id int64, status bool) (cloudAccount *entity.EntityCloudAccount, err error) {

	cloudAccount, err = u.GetByID(id)

	if err != nil {
		return nil, err
	}

	cloudAccount.Active = status

	return cloudAccount, u.db.Save(&cloudAccount).Error
}
