package cron

import (
	"app/entity"
	infrastructure_cloud_provider_aws "app/infrastructure/cloud_provider/aws"
	"app/infrastructure/postgres"
	"app/infrastructure/repository"
	usecase_autoscalling_groups "app/usecase/autoscalling_groups"
	usecase_calendar "app/usecase/calendar"
	usecase_cloud_account "app/usecase/cloud_account"
	usecase_dbinstance "app/usecase/dbinstance"
	usecase_holiday "app/usecase/holiday"
	usecase_instance "app/usecase/instance"
	usecase_log "app/usecase/log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
)

var Scheduler *gocron.Scheduler

func StartCronJobs() {
	var timezoneParam string = os.Getenv("TIME_ZONE")
	var timezone *time.Location

	if timezoneParam == "" {
		timezone = time.UTC
	} else {
		timezone, _ = time.LoadLocation(timezoneParam)
	}

	s := gocron.NewScheduler(timezone)

	Scheduler = s

	s.Every(30).Minutes().Do(updateInstances)

	s.StartAsync()
}

func StartJobsCalendars() {
	conn := postgres.Connect()

	var usecaseLog usecase_log.IUsecaseLog = usecase_log.NewService(
		repository.NewLogPostgres(conn),
	)

	var usecaseDbinstance usecase_dbinstance.IUsecaseDbinstance = usecase_dbinstance.NewService(
		repository.NewDbinstancePostgres(conn),
	)

	var usecaseAutoScallingGroup usecase_autoscalling_groups.IUsecaseAutoScalingGroup = usecase_autoscalling_groups.NewService(
		repository.NewAutoScalingGroupPostgres(conn),
	)

	var usecaseInstance usecase_instance.IUseCaseInstance = usecase_instance.NewService(
		repository.NewInstancePostgres(conn),
	)

	var usecaseCloudAccount usecase_cloud_account.IUsecaseCloudAccount = usecase_cloud_account.NewAWSService(
		repository.NewCloudAccountPostgres(conn),
		usecaseInstance,
		infrastructure_cloud_provider_aws.NewAWSCloudProvider(),
		usecaseDbinstance,
		usecaseAutoScallingGroup,
	)

	var usecaseHoliday usecase_holiday.IUsecaseHoliday = usecase_holiday.NewService(
		repository.NewHolidayPostgres(conn),
	)

	var usecaseCalendar usecase_calendar.IUsecaseCalendar = usecase_calendar.NewService(
		repository.NewCalendarPostgres(conn),
		Scheduler,
		usecaseInstance,
		infrastructure_cloud_provider_aws.NewAWSCloudProvider(),
		usecaseCloudAccount,
		usecaseHoliday,
		usecaseLog,
		usecaseDbinstance,
		usecaseAutoScallingGroup,
	)

	usecaseCalendar.CreateAllCalendarsJob()

}

func updateInstances() {
	db := postgres.Connect()

	var repoCloudProvider usecase_cloud_account.IRepositoryCloudAccount = repository.NewCloudAccountPostgres(db)
	var repoInstances usecase_instance.IRepositoryInstance = repository.NewInstancePostgres(db)

	var ucIntances usecase_instance.IUseCaseInstance = usecase_instance.NewService(repoInstances)
	var ucAutoScallingGroup usecase_autoscalling_groups.IUsecaseAutoScalingGroup = usecase_autoscalling_groups.NewService(repository.NewAutoScalingGroupPostgres(db))
	var ucDbinstance usecase_dbinstance.IUsecaseDbinstance = usecase_dbinstance.NewService(repository.NewDbinstancePostgres(db))
	var ucCloudProvider usecase_cloud_account.IUsecaseCloudAccount = usecase_cloud_account.NewAWSService(
		repoCloudProvider,
		ucIntances,
		infrastructure_cloud_provider_aws.NewAWSCloudProvider(),
		ucDbinstance,
		ucAutoScallingGroup,
	)

	cloudAccounts, _, _ := ucCloudProvider.GetAll(entity.SearchEntityCloudAccountParams{
		Page:     0,
		PageSize: 10000,
	})

	for _, cloudAccount := range cloudAccounts {
		if cloudAccount.Active {
			ucCloudProvider.UpdateAllInstancesOnCloudAccountProvider(&cloudAccount)
			ucCloudProvider.UpdateAllDBInstancesOnCloudAccountProvider(&cloudAccount)
			ucCloudProvider.UpdateAllAutoScalingGroupsOnCloudAccountProvider(&cloudAccount)

		}
	}
}
