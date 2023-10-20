package usecase_cloud_account

import "app/entity"

type IRepositoryCloudAccount interface {
	GetAll() (cloudAccounts []*entity.EntityCloudAccount, err error)
	GetByID(id int64) (cloudAccount *entity.EntityCloudAccount, err error)
	CreateCloudAccount(cloudAccount *entity.EntityCloudAccount) error
	UpdateCloudAccount(cloudAccount *entity.EntityCloudAccount) error
	DeleteCloudAccount(cloudAccount *entity.EntityCloudAccount) error
	ActiveDeactiveCloudAccount(id int64, status bool) (cloudAccount *entity.EntityCloudAccount, err error)
}

type IUsecaseCloudAccount interface {
	GetAll() (cloudAccounts []*entity.EntityCloudAccount, err error)
	GetByID(id int64) (cloudAccount *entity.EntityCloudAccount, err error)
	CreateCloudAccount(cloudAccount *entity.EntityCloudAccount) error
	UpdateCloudAccount(cloudAccount *entity.EntityCloudAccount) error
	DeleteCloudAccount(cloudAccount *entity.EntityCloudAccount) error
	ActiveDeactiveCloudAccount(id int64, status bool) (cloudAccount *entity.EntityCloudAccount, err error)
	UpdateAllInstancesOnCloudAccountProvider(cloudAccount *entity.EntityCloudAccount) (instances []*entity.EntityInstance, err error)
}
