package usecase_calendar

import (
	"app/entity"
	infrastructure_cloud_provider "app/infrastructure/cloud_provider"
	usecase_cloud_account "app/usecase/cloud_account"
	usecase_holiday "app/usecase/holiday"
	usecase_instance "app/usecase/instance"
	usecase_log "app/usecase/log"
	"errors"
	"fmt"
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
	usecaseHoliday       usecase_holiday.IUsecaseHoliday
	usecaseLog           usecase_log.IUsecaseLog
	Now                  func() time.Time
}

func NewService(repository IRepositoryCalendar, scheduler *gocron.Scheduler,
	usecaseInstance usecase_instance.IUseCaseInstance,
	infraCloudProvider infrastructure_cloud_provider.ICloudProvider,
	usecaseCloudAccoount usecase_cloud_account.IUsecaseCloudAccount,
	usecaseHoliday usecase_holiday.IUsecaseHoliday,
	usecaseLog usecase_log.IUsecaseLog) *UsecaseCalendar {
	return &UsecaseCalendar{
		repo:                 repository,
		scheduler:            scheduler,
		usecaseInstance:      usecaseInstance,
		infraCloudProvider:   infraCloudProvider,
		usecaseCloudAccoount: usecaseCloudAccoount,
		usecaseHoliday:       usecaseHoliday,
		usecaseLog:           usecaseLog,
		Now:                  time.Now,
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

func (u *UsecaseCalendar) CreateAllCalendarsJob() error {

	calendars, _, err := u.GetAll(entity.SearchEntityCalendarParams{
		PageSize: 1000000,
		Page:     0,
	})

	if err != nil {
		return err
	}

	for _, calendar := range calendars {
		if calendar.Active {
			u.configureSchedules(&calendar)
		}
	}

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

func (u *UsecaseCalendar) ProcessInstance(instance *entity.EntityInstance, calendar *entity.EntityCalendar) error {
	err := u.infraCloudProvider.Connect(instance.CloudAccount)
	if err != nil {
		log.Println("Error on connect to cloud provider: ", err)
		return err
	}

	if instance.Active && calendar.Active {
		if check, _ := u.usecaseHoliday.IsHoliday(u.Now()); check && !calendar.ValidHoliday {
			return errors.New("today is holiday")
		}

		logInstance := entity.EntityLog{
			Code:     "job execute",
			Instance: fmt.Sprintf("instance id: %s, instance name: %s", instance.InstanceID, instance.InstanceName),
			Content: fmt.Sprintf(
				"instance id: %s, calendar id: %d, calendar name: %s, instance name: %s",
				instance.InstanceID,
				calendar.ID,
				calendar.Name,
				instance.InstanceName),
			CreatedAt: u.Now(),
		}

		if calendar.TypeAction == "on" {
			err = u.infraCloudProvider.StartInstance(instance.InstanceID)
			logInstance.Type = "start"
			if err != nil {
				logInstance.Error = err.Error()
				u.usecaseLog.Create(&logInstance)
				return err
			}
			u.ScheduleUpdateInstance(instance.CloudAccount, *instance, "running")
			u.usecaseLog.Create(&logInstance)

			return nil
		} else if calendar.TypeAction == "off" {
			logInstance.Type = "stop"
			err = u.infraCloudProvider.StopInstance(instance.InstanceID)
			if err != nil {
				logInstance.Error = err.Error()
				u.usecaseLog.Create(&logInstance)
				return err
			}
			u.ScheduleUpdateInstance(instance.CloudAccount, *instance, "stopped")
			u.usecaseLog.Create(&logInstance)
			return nil
		}

	}

	return errors.New("instance or calendar is not active")
}

func (u *UsecaseCalendar) ProccessCalendar(calendar *entity.EntityCalendar) error {
	instances, err := u.usecaseInstance.GetAllOFCalendar(calendar.ID)
	if err != nil {
		return err
	}

	for _, instance := range instances {
		if !instance.Active {
			continue
		}

		go u.ProcessInstance(&instance, calendar)
	}

	return nil
}

func (u *UsecaseCalendar) ScheduleUpdateInstance(cloudAccount entity.EntityCloudAccount,
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

	println("update instance: ", instance.InstanceID, " - ", "counter: ", counter)
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
}

func (u *UsecaseCalendar) cleanTags(calendar *entity.EntityCalendar) {

	var days []int = u.toDaysInt(calendar)

	for _, day := range days {
		var tag string = "calendar_" + strconv.Itoa(calendar.ID) + "_" + strconv.Itoa(day)

		u.scheduler.RemoveByTag(tag)
	}

	for _, job := range u.scheduler.Jobs() {
		if len(job.Tags()) > 0 {
			u.scheduler.RemoveByReference(job)
		}
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
