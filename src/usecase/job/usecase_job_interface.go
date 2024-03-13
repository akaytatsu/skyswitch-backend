package usecase_job

import "app/entity"

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_job.go -package=mocks app/usecase/job IUsecaseJob
type IUsecaseJob interface {
	GetAll() (response []entity.EntityJob, err error)
}
