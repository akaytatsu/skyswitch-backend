package usecase_calendar

import (
	"app/entity"
	infrastructure_cloud_provider "app/infrastructure/cloud_provider"
	usecase_cloud_account "app/usecase/cloud_account"
	usecase_instance "app/usecase/instance"
	"log"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
)

type UsecaseCalendar struct {
	repo                 IRepositoryCalendar
	scheduler            *gocron.Scheduler
	usecaseInstance      usecase_instance.IUseCaseInstance
	infraCloudProvider   infrastructure_cloud_provider.ICloudProvider
	usecaseCloudAccoount usecase_cloud_account.IUsecaseCloudAccount
}

func NewService(repository IRepositoryCalendar, scheduler *gocron.Scheduler,
	usecaseInstance usecase_instance.IUseCaseInstance,
	infraCloudProvider infrastructure_cloud_provider.ICloudProvider,
	usecaseCloudAccoount usecase_cloud_account.IUsecaseCloudAccount) *UsecaseCalendar {
	return &UsecaseCalendar{
		repo:                 repository,
		scheduler:            scheduler,
		usecaseInstance:      usecaseInstance,
		infraCloudProvider:   infraCloudProvider,
		usecaseCloudAccoount: usecaseCloudAccoount,
	}
}

func (u *UsecaseCalendar) Get(id int) (*entity.EntityCalendar, error) {
	return u.repo.GetFromID(id)
}

func (u *UsecaseCalendar) GetAll(searchParams entity.SearchEntityCalendarParams) (response []entity.EntityCalendar, totalRegisters int64, err error) {
	return u.repo.GetAll(searchParams)
}

func (u *UsecaseCalendar) Create(calendar *entity.EntityCalendar) error {
	err := u.repo.Create(calendar)

	if err != nil {
		return err
	}

	u.configureSchedules(calendar)

	return nil
}

func (u *UsecaseCalendar) Update(calendar *entity.EntityCalendar) error {
	err := u.repo.Update(calendar)

	if err != nil {
		return err
	}

	u.configureSchedules(calendar)

	return nil
}

func (u *UsecaseCalendar) Delete(id int) error {

	calendar, err := u.Get(id)

	if err != nil {
		return err
	}

	u.cleanTags(calendar)

	return u.repo.Delete(id)
}

func (u *UsecaseCalendar) ProccessCalendar(calendar *entity.EntityCalendar) error {

	instances, err := u.usecaseInstance.GetAllOFCalendar(calendar.ID)

	if err != nil {
		return err
	}

	for _, instance := range instances {
		err := u.infraCloudProvider.Connect(instance.CloudAccount)

		if err != nil {
			log.Println("Error on connect to cloud provider: ", err)
			continue
		}
		if instance.Active {

			// cloudInstance, err := u.infraCloudProvider.GetInstanceByID(instance.InstanceID)

			// if err != nil {
			// 	log.Println("Error on get instance by id: ", err)
			// 	continue
			// }

			if !calendar.Active {
				continue
			}

			if calendar.TypeAction == "on" {
				println("start: ", instance.InstanceID, " - ", calendar.TypeAction, " - ", calendar.ExecuteTime, " - ", calendar.ID, " - ", calendar.Name, " - ", calendar.Active)
				u.infraCloudProvider.StartInstance(instance.InstanceID)

				u.scheduleUpdateInstance(instance.CloudAccount, instance, "running")
			} else if calendar.TypeAction == "off" {
				println("stop: ", instance.InstanceID, " - ", calendar.TypeAction, " - ", calendar.ExecuteTime, " - ", calendar.ID, " - ", calendar.Name, " - ", calendar.Active)
				u.infraCloudProvider.StopInstance(instance.InstanceID)

				u.scheduleUpdateInstance(instance.CloudAccount, instance, "stopped")
			}
		}
	}

	return nil
}

func (u *UsecaseCalendar) scheduleUpdateInstance(cloudAccount entity.EntityCloudAccount,
	instance entity.EntityInstance, finishStatus string) {

	var counter int = 0

	for {
		instances, _ := u.usecaseCloudAccoount.UpdateAllInstancesOnCloudAccountProvider(&cloudAccount)

		var cloudInstance *entity.EntityInstance

		for _, i := range instances {
			if i.InstanceID == instance.InstanceID {
				cloudInstance = i
				break
			}
		}

		if cloudInstance == nil {
			break
		}

		if cloudInstance.InstanceState == "terminated" {
			break
		}

		if cloudInstance.InstanceState == finishStatus {
			break
		}
		counter++

		if counter > 15 {
			break
		}
		time.Sleep(30 * time.Second)
	}
}

func (u *UsecaseCalendar) configureSchedules(calendar *entity.EntityCalendar) {

	u.cleanTags(calendar)

	if calendar.Active {
		var days []int = u.toDaysInt(calendar)

		for _, day := range days {

			var tag string = "calendar_" + strconv.Itoa(calendar.ID) + "_" + strconv.Itoa(day)

			weekday := time.Weekday(day)
			_, err := u.scheduler.Every(1).Weekday(weekday).At(calendar.ExecuteTime).Tag(tag).Do(func() {
				u.ProccessCalendar(calendar)
			})

			if err != nil {
				println(err.Error())
			}
		}
	}

	// lista todos os jobs agendados
	// jobs := u.scheduler.Jobs()
	// for _, job := range jobs {
	// 	println(job.Tags, job.NextRun().GoString(), job.ScheduledTime().GoString())
	// }
}

func (u *UsecaseCalendar) cleanTags(calendar *entity.EntityCalendar) {

	var days []int = u.toDaysInt(calendar)

	for _, day := range days {
		var tag string = "calendar_" + strconv.Itoa(calendar.ID) + "_" + strconv.Itoa(day)

		u.scheduler.RemoveByTag(tag)
	}

}

func (u *UsecaseCalendar) toDaysInt(calendar *entity.EntityCalendar) (days []int) {

	if calendar.Sunday {
		days = append(days, 0)
	}

	if calendar.Monday {
		days = append(days, 1)
	}

	if calendar.Tuesday {
		days = append(days, 2)
	}

	if calendar.Wednesday {
		days = append(days, 3)
	}

	if calendar.Thursday {
		days = append(days, 4)
	}

	if calendar.Friday {
		days = append(days, 5)
	}

	if calendar.Saturday {
		days = append(days, 6)
	}

	return days
}
