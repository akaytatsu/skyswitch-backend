package entity

import (
	"time"
)

type EntityCloudAccount struct {
	ID              int64  `gorm:"column:id;primary_key" json:"id"`
	CloudProvider   string `gorm:"column:cloud_provider" json:"cloud_provider"`
	Nickname        string `gorm:"column:nickname" json:"nickname"`
	AccessKeyID     string `gorm:"column:access_key_id" json:"access_key_id"`
	SecretAccessKey string `gorm:"column:secret_access_key" json:"secret_access_key"`
	Region          string `gorm:"varchar(20);column:region" json:"region"`
	Active          bool   `gorm:"column:active;default:true" json:"active"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// func (e *EntityCloudAccount) MarshalJSON() ([]byte, error) {
// 	type Temp struct {
// 		ID            int64  `gorm:"column:id;primary_key" json:"id"`
// 		CloudProvider string `gorm:"column:cloud_provider" json:"cloud_provider"`
// 		Nickname      string `gorm:"column:nickname" json:"nickname"`
// 		AccessKeyID   string `gorm:"column:access_key_id" json:"access_key_id"`
// 		Active        bool   `gorm:"column:active;default:true" json:"active"`
// 		CreatedAt     time.Time
// 		UpdatedAt     time.Time
// 	}

// 	t := Temp{
// 		ID:            e.ID,
// 		CloudProvider: e.CloudProvider,
// 		Nickname:      e.Nickname,
// 		AccessKeyID:   e.AccessKeyID,
// 		Active:        e.Active,
// 		CreatedAt:     e.CreatedAt,
// 		UpdatedAt:     e.UpdatedAt,
// 	}

// 	return json.Marshal(t)
// }

type SearchEntityCloudAccountParams struct {
	OrderBy   string `json:"order_by"`
	SortOrder string `json:"sort_order"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	Q         string `json:"q"`
	CreatedAt string `json:"created_at"`
}
