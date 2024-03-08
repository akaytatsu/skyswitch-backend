package entity

import "time"

type EntityInstance struct {
	ID             int64              `gorm:"column:id;primary_key" json:"id"`
	CloudAccountID int64              `gorm:"column:cloud_account_id" json:"cloud_account_id"`
	CloudAccount   EntityCloudAccount `gorm:"foreignKey:CloudAccountID;references:ID" json:"cloud_account"`
	Calendars      []EntityCalendar   `gorm:"many2many:entity_instance_calendars" json:"calendars"`
	InstanceID     string             `gorm:"column:instance_id" json:"instance_id"`
	InstanceType   string             `gorm:"column:instance_type" json:"instance_type"`
	InstanceName   string             `gorm:"column:instance_name" json:"instance_name"`
	InstanceState  string             `gorm:"column:instance_state" json:"instance_state"`
	InstanceRegion string             `gorm:"column:instance_region" json:"instance_region"`
	Active         bool               `gorm:"column:active;default:true" json:"active"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type SearchEntityInstanceParams struct {
	OrderBy   string `json:"order_by"`
	SortOrder string `json:"sort_order"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	Q         string `json:"q"`
	CreatedAt string `json:"created_at"`
}
