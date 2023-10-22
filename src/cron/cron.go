package cron

import (
	"time"

	"github.com/go-co-op/gocron"
)

func StartCronJobs() {
	s := gocron.NewScheduler(time.UTC)

	s.StartAsync()
}

func updateInstances() {
	// db := postgres.Connect()

	// repo := repository.NewCloudAccountPostgres(db)

	// var usecase usecase_cloud_account.IUsecaseCloudAccount = usecase_cloud_account.NewAWSService(repo)

	// cloudAccounts, _ := usecase.GetAll()

	// for _, cloudAccount := range cloudAccounts {
	// 	if cloudAccount.Active {
	// 		usecase.UpdateAllInstancesOnCloudAccountProvider(cloudAccount)

	// 	}
	// }
}
