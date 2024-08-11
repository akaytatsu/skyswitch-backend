package usecase_dbinstance

import "app/entity"

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_dbinstance.go -package=mocks app/usecase/dbinstance IRepositoryDbinstance
type IRepositoryDbinstance interface {
	GetFromID(id int) (*entity.EntityDbinstance, error)
	GetAll(searchParams entity.SearchEntityDbinstanceParams) (response []entity.EntityDbinstance, totalRegisters int64, err error)
	Create(*entity.EntityDbinstance) error
	Update(*entity.EntityDbinstance) error
	Delete(id int) error
	FromCalendar(calendarID int) (response []entity.EntityDbinstance, err error)
	GetByInstanceID(instanceID string) (instance *entity.EntityDbinstance, err error)
	ActiveDeactiveInstance(id int64, status bool) (instance *entity.EntityDbinstance, err error)
}

//go:generate mockgen -destination=../../mocks/mock_usecase_dbinstance.go -package=mocks app/usecase/dbinstance IUsecaseDbinstance
type IUsecaseDbinstance interface {
	Get(id int) (*entity.EntityDbinstance, error)
	GetAll(searchParams entity.SearchEntityDbinstanceParams) (response []entity.EntityDbinstance, totalRegisters int64, err error)
	Create(*entity.EntityDbinstance) error
	Update(*entity.EntityDbinstance) error
	CreateOrUpdateDbInstance(instance *entity.EntityDbinstance, updateCalendars bool) error
	Delete(id int) error
	GetAllOFCalendar(calendarID int) (response []entity.EntityDbinstance, err error)
	ActiveDeactiveInstance(id int64, status bool) (instance *entity.EntityDbinstance, err error)
	GetByInstanceID(instanceID string) (instance *entity.EntityDbinstance, err error)
}
