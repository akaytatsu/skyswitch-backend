package usecase_instance

import (
	"app/entity"
	"log"
)

type UseCaseInstance struct {
	repo IRepositoryInstance
}

func NewService(repository IRepositoryInstance) *UseCaseInstance {
	return &UseCaseInstance{repo: repository}
}

func (u *UseCaseInstance) GetAll(searchParams entity.SearchEntityInstanceParams) (response []entity.EntityInstance, totalRegisters int64, err error) {
	return u.repo.GetAll(searchParams)
}

func (u *UseCaseInstance) GetByID(id int64) (instance *entity.EntityInstance, err error) {
	return u.repo.GetByID(id)
}

func (u *UseCaseInstance) CreateInstance(instance *entity.EntityInstance) error {
	return u.repo.CreateInstance(instance)
}

func (u *UseCaseInstance) UpdateInstance(instance *entity.EntityInstance, updateCalendars bool) error {
	return u.repo.UpdateInstance(instance, updateCalendars)
}

func (u *UseCaseInstance) DeleteInstance(instance *entity.EntityInstance) error {
	return u.repo.DeleteInstance(instance)
}

func (u *UseCaseInstance) ActiveDeactiveInstance(id int64, status bool) (instance *entity.EntityInstance, err error) {
	return u.repo.ActiveDeactiveInstance(id, status)
}

func (u *UseCaseInstance) CreateOrUpdateInstance(instance *entity.EntityInstance, updateCalendars bool) error {

	if instance.InstanceID != "" {
		instanceLocal, err := u.repo.GetByInstanceID(instance.InstanceID)

		if err != nil {
			log.Println("Error on get instance by instance id: ", err)
			return u.repo.CreateInstance(instance)
		}

		instance.ID = instanceLocal.ID

		return u.repo.UpdateInstance(instance, updateCalendars)
	}

	return u.repo.CreateInstance(instance)
}

func (u *UseCaseInstance) GetAllOFCalendar(calendarID int) (response []entity.EntityInstance, err error) {
	return u.repo.FromCalendar(calendarID)
}

func (u *UseCaseInstance) GetByInstanceID(instanceID string) (instance *entity.EntityInstance, err error) {
	return u.repo.GetByInstanceID(instanceID)
}
