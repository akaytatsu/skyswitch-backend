package usecase_job

import "app/entity"

type IUsecaseJob interface {
	GetAll() (response []entity.EntityJob, err error)
}
