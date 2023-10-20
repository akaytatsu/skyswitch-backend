package usecase_instance

import "app/entity"

type UseCaseInstance struct {
	repo IRepositoryInstance
}

func NewService(repository IRepositoryInstance) *UseCaseInstance {
	return &UseCaseInstance{repo: repository}
}

func (u *UseCaseInstance) GetAll() (instances []*entity.EntityInstance, err error) {
	return u.repo.GetAll()
}

func (u *UseCaseInstance) GetByID(id int64) (instance *entity.EntityInstance, err error) {
	return u.repo.GetByID(id)
}

func (u *UseCaseInstance) CreateInstance(instance *entity.EntityInstance) error {
	return u.repo.CreateInstance(instance)
}

func (u *UseCaseInstance) UpdateInstance(instance *entity.EntityInstance) error {
	return u.repo.UpdateInstance(instance)
}

func (u *UseCaseInstance) DeleteInstance(instance *entity.EntityInstance) error {
	return u.repo.DeleteInstance(instance)
}

func (u *UseCaseInstance) ActiveDeactiveInstance(id int64, status bool) (instance *entity.EntityInstance, err error) {
	return u.repo.ActiveDeactiveInstance(id, status)
}
