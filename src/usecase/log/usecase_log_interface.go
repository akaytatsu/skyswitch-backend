package usecase_log

import "app/entity"

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_log.go -package=mocks app/usecase/log IRepositoryLog
type IRepositoryLog interface {
	GetFromID(id int) (*entity.EntityLog, error)
	GetAll(searchParams entity.SearchEntityLogParams) (response []entity.EntityLog, totalRegisters int64, err error)
	Create(*entity.EntityLog) error
	Update(*entity.EntityLog) error
	Delete(id int) error
}

//go:generate mockgen -destination=../../mocks/mock_usecase_log.go -package=mocks app/usecase/log IUsecaseLog
type IUsecaseLog interface {
	Get(id int) (*entity.EntityLog, error)
	GetAll(searchParams entity.SearchEntityLogParams) (response []entity.EntityLog, totalRegisters int64, err error)
	Create(*entity.EntityLog) error
	Update(*entity.EntityLog) error
	Delete(id int) error
}
