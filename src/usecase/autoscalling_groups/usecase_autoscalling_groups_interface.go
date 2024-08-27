package usecase_autoscalling_groups

import "app/entity"

//go:generate mockgen -destination=../../mocks/mock_usecase_repository_autoscalling_groups.go -package=mocks app/usecase/autoscalling_groups IRepositoryAutoScalingGroup
type IRepositoryAutoScalingGroup interface {
	GetFromID(id int) (*entity.EntityAutoScalingGroup, error)
	GetAll(searchParams entity.SearchEntityAutoScalingGroupParams) (response []entity.EntityAutoScalingGroup, totalRegisters int64, err error)
	Create(*entity.EntityAutoScalingGroup) error
	Update(*entity.EntityAutoScalingGroup) error
	Delete(id int) error
	FromCalendar(calendarID int) (response []entity.EntityAutoScalingGroup, err error)
	GetByID(autoScallingGroupID string) (autoScallingGroup *entity.EntityAutoScalingGroup, err error)
	ActiveDeactive(id int64, status bool) (autoScallingGroup *entity.EntityAutoScalingGroup, err error)
}

//go:generate mockgen -destination=../../mocks/mock_usecase_autoscalling_groups.go -package=mocks app/usecase/autoscalling_groups IUsecaseAutoScalingGroup
type IUsecaseAutoScalingGroup interface {
	Get(id int) (*entity.EntityAutoScalingGroup, error)
	GetAll(searchParams entity.SearchEntityAutoScalingGroupParams) (response []entity.EntityAutoScalingGroup, totalRegisters int64, err error)
	Create(*entity.EntityAutoScalingGroup) error
	Update(*entity.EntityAutoScalingGroup) error
	CreateOrUpdate(autoScallingGroup *entity.EntityAutoScalingGroup, updateCalendars bool) error
	Delete(id int) error
	GetAllOFCalendar(calendarID int) (response []entity.EntityAutoScalingGroup, err error)
	ActiveDeactive(id int64, status bool) (autoScallingGroup *entity.EntityAutoScalingGroup, err error)
	GetByID(autoScallingGroupID string) (autoScallingGroup *entity.EntityAutoScalingGroup, err error)
}
