package entity

import "time"

type EntityInstance struct {
	Id             int64  `gorm:"column:id;primary_key" json:"id"`
	CloudAccountID int64  `gorm:"column:cloud_account_id" json:"cloud_account_id"`
	InstanceID     string `gorm:"column:instance_id" json:"instance_id"`
	InstanceType   string `gorm:"column:instance_type" json:"instance_type"`
	InstanceName   string `gorm:"column:instance_name" json:"instance_name"`
	InstanceState  string `gorm:"column:instance_state" json:"instance_state"`
	InstanceRegion string `gorm:"column:instance_region" json:"instance_region"`
	Active         bool   `gorm:"column:active;default:true" json:"active"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
