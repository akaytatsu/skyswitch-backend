package usecase_cloud_account

import (
	"app/entity"
	infrastructure_cloud_provider "app/infrastructure/cloud_provider"
	usecase_dbinstance "app/usecase/dbinstance"
	usecase_instance "app/usecase/instance"
	"log"
)

type UseCaseAWSCloudAccount struct {
	repo               IRepositoryCloudAccount
	useCaseInstances   usecase_instance.IUseCaseInstance
	UsecaseDbinstance  usecase_dbinstance.IUsecaseDbinstance
	infraCloudProvider infrastructure_cloud_provider.ICloudProvider
}

func NewAWSService(repository IRepositoryCloudAccount, usecaseInstances usecase_instance.IUseCaseInstance,
	infraCloudProvider infrastructure_cloud_provider.ICloudProvider, usecaseDbinstances usecase_dbinstance.IUsecaseDbinstance) *UseCaseAWSCloudAccount {
	return &UseCaseAWSCloudAccount{repo: repository, useCaseInstances: usecaseInstances, infraCloudProvider: infraCloudProvider, UsecaseDbinstance: usecaseDbinstances}
}

func (u *UseCaseAWSCloudAccount) GetAll(searchParams entity.SearchEntityCloudAccountParams) (response []entity.EntityCloudAccount, totalRegisters int64, err error) {
	return u.repo.GetAll(searchParams)
}

func (u *UseCaseAWSCloudAccount) GetByID(id int64) (cloudAccount *entity.EntityCloudAccount, err error) {
	return u.repo.GetByID(id)
}

func (u *UseCaseAWSCloudAccount) CreateCloudAccount(cloudAccount *entity.EntityCloudAccount) error {
	err := u.repo.CreateCloudAccount(cloudAccount)

	if err != nil {
		return err
	}

	go u.UpdateAllInstancesOnCloudAccountProvider(cloudAccount)

	return nil
}

func (u *UseCaseAWSCloudAccount) UpdateCloudAccount(cloudAccount *entity.EntityCloudAccount) error {
	err := u.repo.UpdateCloudAccount(cloudAccount)

	if err != nil {
		return err
	}

	go u.UpdateAllInstancesOnCloudAccountProvider(cloudAccount)

	return nil
}

func (u *UseCaseAWSCloudAccount) DeleteCloudAccount(cloudAccount *entity.EntityCloudAccount) error {
	return u.repo.DeleteCloudAccount(cloudAccount)
}

func (u *UseCaseAWSCloudAccount) ActiveDeactiveCloudAccount(id int64, status bool) (cloudAccount *entity.EntityCloudAccount, err error) {
	return u.repo.ActiveDeactiveCloudAccount(id, status)
}

func (u *UseCaseAWSCloudAccount) UpdateAllInstancesOnAllCloudAccountProvider() (instances []*entity.EntityInstance, err error) {

	params := entity.SearchEntityCloudAccountParams{
		Page:     0,
		PageSize: 10000,
	}

	cloudAccounts, _, err := u.repo.GetAll(params)
	if err != nil {
		return nil, err
	}

	for _, cloudAccount := range cloudAccounts {
		instances, err = u.UpdateAllInstancesOnCloudAccountProvider(&cloudAccount)
		if err != nil {
			log.Println("Error updating all instances on cloud account provider: ", err)
		}
	}

	return instances, nil
}

func (u *UseCaseAWSCloudAccount) UpdateAllInstancesOnCloudAccountProvider(cloudAccount *entity.EntityCloudAccount) (instances []*entity.EntityInstance, err error) {

	cloudProvier, err := u.infraCloudProvider.Connect(*cloudAccount)

	if err != nil {
		return nil, err
	}

	instances, err = cloudProvier.GetInstances()

	if err != nil {
		return instances, err
	}

	for _, instance := range instances {

		aux, _ := u.useCaseInstances.GetByInstanceID(instance.InstanceID)

		instance.Active = aux.Active

		err = u.useCaseInstances.CreateOrUpdateInstance(instance, false)

		if err != nil {
			log.Println("Error creating or updating instance: ", err)
		}
	}

	return instances, nil
}

func (u *UseCaseAWSCloudAccount) UpdateAllInstancesOnCloudAccountProviderFromID(id int) (instances []*entity.EntityInstance, err error) {
	cloudAccount, err := u.GetByID(int64(id))
	if err != nil {
		return nil, err
	}

	return u.UpdateAllInstancesOnCloudAccountProvider(cloudAccount)
}

func (u *UseCaseAWSCloudAccount) UpdateAllDBInstancesOnCloudAccountProvider(cloudAccount *entity.EntityCloudAccount) (dbInstances []*entity.EntityDbinstance, err error) {
	cloudProvier, err := u.infraCloudProvider.Connect(*cloudAccount)

	if err != nil {
		return nil, err
	}

	dbinstances, err := cloudProvier.GetDBInstances()

	if err != nil {
		return dbinstances, err
	}

	for _, dbinstance := range dbinstances {
		err = u.UsecaseDbinstance.CreateOrUpdateDbInstance(dbinstance, false)

		if err != nil {
			log.Println("Error creating or updating db instance: ", err)
		}
	}

	return dbinstances, nil
}

func (u *UseCaseAWSCloudAccount) UpdateAllDBInstancesOnCloudAccountProviderFromID(id int) (dbInstances []*entity.EntityDbinstance, err error) {
	cloudAccount, err := u.GetByID(int64(id))
	if err != nil {
		return nil, err
	}

	return u.UpdateAllDBInstancesOnCloudAccountProvider(cloudAccount)
}

func (u *UseCaseAWSCloudAccount) UpdateAllDBInstancesOnAllCloudAccountProvider() (dbInstances []*entity.EntityDbinstance, err error) {

	params := entity.SearchEntityCloudAccountParams{
		Page:     0,
		PageSize: 10000,
	}

	cloudAccounts, _, err := u.repo.GetAll(params)
	if err != nil {
		return nil, err
	}

	for _, cloudAccount := range cloudAccounts {
		dbInstances, err = u.UpdateAllDBInstancesOnCloudAccountProvider(&cloudAccount)
		if err != nil {
			log.Println("Error updating all db instances on cloud account provider: ", err)
		}
	}

	return dbInstances, nil
}
