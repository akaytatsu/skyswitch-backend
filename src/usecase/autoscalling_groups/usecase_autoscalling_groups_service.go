package usecase_autoscalling_groups

import "app/entity"

type UsecaseAutoScalingGroup struct {
	repo IRepositoryAutoScalingGroup
}

func NewService(repository IRepositoryAutoScalingGroup) *UsecaseAutoScalingGroup {
	return &UsecaseAutoScalingGroup{repo: repository}
}

func (u *UsecaseAutoScalingGroup) Get(id int) (*entity.EntityAutoScalingGroup, error) {
	return u.repo.GetFromID(id)
}

func (u *UsecaseAutoScalingGroup) GetAll(searchParams entity.SearchEntityAutoScalingGroupParams) (response []entity.EntityAutoScalingGroup, totalRegisters int64, err error) {
	return u.repo.GetAll(searchParams)
}

func (u *UsecaseAutoScalingGroup) Create(autoScallingGroup *entity.EntityAutoScalingGroup) error {
	return u.repo.Create(autoScallingGroup)
}

func (u *UsecaseAutoScalingGroup) Update(autoScallingGroup *entity.EntityAutoScalingGroup) error {
	return u.repo.Update(autoScallingGroup)
}

func (u *UsecaseAutoScalingGroup) CreateOrUpdate(autoScallingGroup *entity.EntityAutoScalingGroup, updateCalendars bool) error {
	if autoScallingGroup.AutoScalingGroupID != "" {
		autoScallingGroupLocal, err := u.repo.GetByID(autoScallingGroup.AutoScalingGroupID)

		if err != nil {
			return u.repo.Create(autoScallingGroup)
		}

		autoScallingGroup.ID = autoScallingGroupLocal.ID

		return u.repo.Update(autoScallingGroup)
	}

	return u.repo.Create(autoScallingGroup)
}

func (u *UsecaseAutoScalingGroup) Delete(id int) error {
	return u.repo.Delete(id)
}

func (u *UsecaseAutoScalingGroup) GetAllOFCalendar(calendarID int) (response []entity.EntityAutoScalingGroup, err error) {
	return u.repo.FromCalendar(calendarID)
}

func (u *UsecaseAutoScalingGroup) ActiveDeactive(id int64, status bool) (autoScallingGroup *entity.EntityAutoScalingGroup, err error) {
	return u.repo.ActiveDeactive(id, status)
}

func (u *UsecaseAutoScalingGroup) GetByID(autoScallingGroupID string) (autoScallingGroup *entity.EntityAutoScalingGroup, err error) {
	return u.repo.GetByID(autoScallingGroupID)
}
