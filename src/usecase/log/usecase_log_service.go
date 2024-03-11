package usecase_log

import "app/entity"

type UsecaseLog struct {
	repo IRepositoryLog
}

func NewService(repository IRepositoryLog) *UsecaseLog {
	return &UsecaseLog{repo: repository}
}

func (u *UsecaseLog) Get(id int) (*entity.EntityLog, error) {
	return u.repo.GetFromID(id)
}

func (u *UsecaseLog) GetAll(searchParams entity.SearchEntityLogParams) (response []entity.EntityLog, totalRegisters int64, err error) {
	return u.repo.GetAll(searchParams)
}

func (u *UsecaseLog) Create(log *entity.EntityLog) error {
	return u.repo.Create(log)
}

func (u *UsecaseLog) Update(log *entity.EntityLog) error {
	return u.repo.Update(log)
}

func (u *UsecaseLog) Delete(id int) error {
	return u.repo.Delete(id)
}
