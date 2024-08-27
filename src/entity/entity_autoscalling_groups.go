package entity

import (
	"github.com/go-playground/validator/v10"
)

type EntityAutoScalingGroup struct {
	ID                   int                `json:"id" gorm:"primaryKey"`
	CloudAccountID       int64              `gorm:"column:cloud_account_id" json:"cloud_account_id"`
	CloudAccount         EntityCloudAccount `gorm:"foreignKey:CloudAccountID;references:ID" json:"cloud_account"`
	Calendars            []EntityCalendar   `gorm:"many2many:entity_autoscalling_groups_calendars" json:"calendars"`
	AutoScalingGroupID   string             `gorm:"column:auto_scaling_group_id" json:"auto_scaling_group_id"`
	AutoScalingGroupName string             `gorm:"column:auto_scaling_group_name" json:"auto_scaling_group_name"`
	MinSize              int                `gorm:"column:min_size" json:"min_size"`
	MaxSize              int                `gorm:"column:max_size" json:"max_size"`
	DesiredCapacity      int                `gorm:"column:desired_capacity" json:"desired_capacity"`
	TotalInstances       int                `gorm:"column:total_instances" json:"total_instances"`
}

func NewEntityAutoScalingGroup(entityAutoScalingGroupParam EntityAutoScalingGroup) (*EntityAutoScalingGroup, error) {
	u := &EntityAutoScalingGroup{
		ID:                   entityAutoScalingGroupParam.ID,
		CloudAccountID:       entityAutoScalingGroupParam.CloudAccountID,
		CloudAccount:         entityAutoScalingGroupParam.CloudAccount,
		Calendars:            entityAutoScalingGroupParam.Calendars,
		AutoScalingGroupID:   entityAutoScalingGroupParam.AutoScalingGroupID,
		AutoScalingGroupName: entityAutoScalingGroupParam.AutoScalingGroupName,
		MinSize:              entityAutoScalingGroupParam.MinSize,
		MaxSize:              entityAutoScalingGroupParam.MaxSize,
		DesiredCapacity:      entityAutoScalingGroupParam.DesiredCapacity,
		TotalInstances:       entityAutoScalingGroupParam.TotalInstances,
	}

	return u, nil
}

func (u *EntityAutoScalingGroup) Validate() error {
	return validator.New().Struct(u)
}

type SearchEntityAutoScalingGroupParams struct {
	OrderBy   string `json:"order_by"`
	SortOrder string `json:"sort_order"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	Q         string `json:"q"`
	CreatedAt string `json:"created_at"`
}
