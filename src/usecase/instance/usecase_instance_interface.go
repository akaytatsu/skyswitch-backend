package usecase_instance

import "app/entity"

type IRepositoryInstance interface {
	GetAll(searchParams entity.SearchEntityInstanceParams) (response []entity.EntityInstance, totalRegisters int64, err error)
	GetByID(id int64) (instance *entity.EntityInstance, err error)
	GetByInstanceID(instanceID string) (instance *entity.EntityInstance, err error)
	CreateInstance(instance *entity.EntityInstance) error
	UpdateInstance(instance *entity.EntityInstance) error
	DeleteInstance(instance *entity.EntityInstance) error
	ActiveDeactiveInstance(id int64, status bool) (instance *entity.EntityInstance, err error)
}

type IUseCaseInstance interface {
	GetAll(searchParams entity.SearchEntityInstanceParams) (response []entity.EntityInstance, totalRegisters int64, err error)
	GetByID(id int64) (instance *entity.EntityInstance, err error)
	CreateInstance(instance *entity.EntityInstance) error
	UpdateInstance(instance *entity.EntityInstance) error
	CreateOrUpdateInstance(instance *entity.EntityInstance) error
	DeleteInstance(instance *entity.EntityInstance) error
	ActiveDeactiveInstance(id int64, status bool) (instance *entity.EntityInstance, err error)
}
