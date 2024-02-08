package cron

import (
	"app/infrastructure/postgres"
	"app/infrastructure/repository"
	usecase_cloud_account "app/usecase/cloud_account"
	usecase_instance "app/usecase/instance"
	"time"

	"github.com/go-co-op/gocron"
)

func StartCronJobs() {
	s := gocron.NewScheduler(time.UTC)

	updateInstances()

	s.StartAsync()
}

func updateInstances() {
	db := postgres.Connect()

	var repoCloudProvider usecase_cloud_account.IRepositoryCloudAccount = repository.NewCloudAccountPostgres(db)
	var repoInstances usecase_instance.IRepositoryInstance = repository.NewInstancePostgres(db)

	var ucIntances usecase_instance.IUseCaseInstance = usecase_instance.NewService(repoInstances)
	var ucCloudProvider usecase_cloud_account.IUsecaseCloudAccount = usecase_cloud_account.NewAWSService(repoCloudProvider, ucIntances)

	cloudAccounts, _ := ucCloudProvider.GetAll()

	for _, cloudAccount := range cloudAccounts {
		if cloudAccount.Active {
			ucCloudProvider.UpdateAllInstancesOnCloudAccountProvider(cloudAccount)

		}
	}
}
