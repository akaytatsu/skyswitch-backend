package usecase_cloud_account

import "app/entity"

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_cloud_account.go -package=mocks app/usecase/cloud_account IRepositoryCloudAccount
type IRepositoryCloudAccount interface {
	GetAll(searchParams entity.SearchEntityCloudAccountParams) (response []entity.EntityCloudAccount, totalRegisters int64, err error)
	GetByID(id int64) (cloudAccount *entity.EntityCloudAccount, err error)
	CreateCloudAccount(cloudAccount *entity.EntityCloudAccount) error
	UpdateCloudAccount(cloudAccount *entity.EntityCloudAccount) error
	DeleteCloudAccount(cloudAccount *entity.EntityCloudAccount) error
	ActiveDeactiveCloudAccount(id int64, status bool) (cloudAccount *entity.EntityCloudAccount, err error)
}

//go:generate mockgen -destination=../../mocks/mock_usecase_cloud_account.go -package=mocks app/usecase/cloud_account IUsecaseCloudAccount
type IUsecaseCloudAccount interface {
	GetAll(searchParams entity.SearchEntityCloudAccountParams) (response []entity.EntityCloudAccount, totalRegisters int64, err error)
	GetByID(id int64) (cloudAccount *entity.EntityCloudAccount, err error)
	CreateCloudAccount(cloudAccount *entity.EntityCloudAccount) error
	UpdateCloudAccount(cloudAccount *entity.EntityCloudAccount) error
	DeleteCloudAccount(cloudAccount *entity.EntityCloudAccount) error
	ActiveDeactiveCloudAccount(id int64, status bool) (cloudAccount *entity.EntityCloudAccount, err error)
	UpdateAllInstancesOnCloudAccountProvider(cloudAccount *entity.EntityCloudAccount) (instances []*entity.EntityInstance, err error)
	UpdateAllInstancesOnCloudAccountProviderFromID(id int) (instances []*entity.EntityInstance, err error)
	UpdateAllInstancesOnAllCloudAccountProvider() (instances []*entity.EntityInstance, err error)
	UpdateAllDBInstancesOnCloudAccountProvider(cloudAccount *entity.EntityCloudAccount) (dbInstances []*entity.EntityDbinstance, err error)
	UpdateAllDBInstancesOnCloudAccountProviderFromID(id int) (dbInstances []*entity.EntityDbinstance, err error)
	UpdateAllDBInstancesOnAllCloudAccountProvider() (dbInstances []*entity.EntityDbinstance, err error)
}
