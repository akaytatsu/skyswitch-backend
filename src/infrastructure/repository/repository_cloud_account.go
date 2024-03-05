package repository

import (
	"app/entity"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RepositoryCloudAccount struct {
	DB *gorm.DB
}

func NewCloudAccountPostgres(db *gorm.DB) *RepositoryCloudAccount {
	return &RepositoryCloudAccount{DB: db}
}

func (u *RepositoryCloudAccount) GetAll(searchParams entity.SearchEntityCloudAccountParams) (response []entity.EntityCloudAccount, totalRegisters int64, err error) {
	offset := (searchParams.Page) * searchParams.PageSize

	qry := u.DB.Model(entity.EntityCloudAccount{})

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
