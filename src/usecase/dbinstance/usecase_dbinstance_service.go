package usecase_dbinstance

import "app/entity"

type UsecaseDbinstance struct {
	repo IRepositoryDbinstance
}

func NewService(repository IRepositoryDbinstance) *UsecaseDbinstance {
	return &UsecaseDbinstance{repo: repository}
}

func (u *UsecaseDbinstance) Get(id int) (*entity.EntityDbinstance, error) {
	return u.repo.GetFromID(id)
}

func (u *UsecaseDbinstance) GetAll(searchParams entity.SearchEntityDbinstanceParams) (response []entity.EntityDbinstance, totalRegisters int64, err error) {
	return u.repo.GetAll(searchParams)
}

func (u *UsecaseDbinstance) Create(dbinstance *entity.EntityDbinstance) error {
	return u.repo.Create(dbinstance)
}

func (u *UsecaseDbinstance) Update(dbinstance *entity.EntityDbinstance) error {
	return u.repo.Update(dbinstance)
}

func (u *UsecaseDbinstance) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u *UsecaseDbinstance) CreateOrUpdateDbInstance(instance *entity.EntityDbinstance, updateCalendars bool) error {
	if instance.DBInstanceID != "" {
		instanceLocal, err := u.repo.GetByInstanceID(instance.DBInstanceID)

		if err != nil {
			return u.repo.Create(instance)
		}

		instance.ID = instanceLocal.ID

		return u.repo.Update(instance)
	}

	return u.repo.Create(instance)
}

func (u *UsecaseDbinstance) GetAllOFCalendar(calendarID int) (response []entity.EntityDbinstance, err error) {
	return u.repo.FromCalendar(calendarID)
}

func (u *UsecaseDbinstance) ActiveDeactiveInstance(id int64, status bool) (instance *entity.EntityDbinstance, err error) {
	return u.repo.ActiveDeactiveInstance(id, status)
}

func (u *UsecaseDbinstance) GetByInstanceID(instanceID string) (instance *entity.EntityDbinstance, err error) {
	return u.repo.GetByInstanceID(instanceID)
}
