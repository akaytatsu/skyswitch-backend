package usecase_log

import "app/entity"

type IRepositoryLog interface {
	GetFromID(id int) (*entity.EntityLog, error)
	GetAll(searchParams entity.SearchEntityLogParams) (response []entity.EntityLog, totalRegisters int64, err error)
	Create(*entity.EntityLog) error
	Update(*entity.EntityLog) error
	Delete(id int) error
}

type IUsecaseLog interface {
	Get(id int) (*entity.EntityLog, error)
	GetAll(searchParams entity.SearchEntityLogParams) (response []entity.EntityLog, totalRegisters int64, err error)
	Create(*entity.EntityLog) error
	Update(*entity.EntityLog) error
	Delete(id int) error
}
