package entity

import "time"

type EntityCloudAccount struct {
	ID              int64  `gorm:"column:id;primary_key" json:"id"`
	CloudProvider   string `gorm:"column:cloud_provider" json:"cloud_provider"`
	Nickname        string `gorm:"column:nickname" json:"nickname"`
	AccessKeyID     string `gorm:"column:access_key_id" json:"access_key_id"`
	SecretAccessKey string `gorm:"column:secret_access_key" json:"secret_access_key"`
	Active          bool   `gorm:"column:active;default:true" json:"active"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
