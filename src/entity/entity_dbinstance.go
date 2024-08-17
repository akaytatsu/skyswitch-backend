package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type EntityDbinstance struct {
	ID               int                `json:"id" gorm:"primaryKey"`
	CloudAccountID   int64              `gorm:"column:cloud_account_id" json:"cloud_account_id"`
	CloudAccount     EntityCloudAccount `gorm:"foreignKey:CloudAccountID;references:ID" json:"cloud_account"`
	Calendars        []EntityCalendar   `gorm:"many2many:entity_dbinstance_calendars" json:"calendars"`
	DBInstanceID     string             `gorm:"column:db_instance_id" json:"db_instance_id"`
	DBInstanceType   string             `gorm:"column:db_instance_type" json:"db_instance_type"`
	DBInstanceName   string             `gorm:"column:db_instance_name" json:"db_instance_name"`
	DBInstanceState  string             `gorm:"column:db_instance_state" json:"db_instance_state"`
	DBInstanceRegion string             `gorm:"column:db_instance_region" json:"db_instance_region"`
	DBInstanceClass  string             `gorm:"column:db_instance_class" json:"db_instance_class"`
	Endpoint         string             `gorm:"column:endpoint" json:"endpoint"`
	Port             int64              `gorm:"column:port" json:"port"`
	Engine           string             `gorm:"column:engine" json:"engine"`
	Active           bool               `json:"active" gorm:"default:true"`
	CreatedAt        time.Time          `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt        time.Time          `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewEntityDbinstance(entityDbinstanceParam EntityDbinstance) (*EntityDbinstance, error) {
	u := &EntityDbinstance{
		ID:               entityDbinstanceParam.ID,
		CloudAccountID:   entityDbinstanceParam.CloudAccountID,
		CloudAccount:     entityDbinstanceParam.CloudAccount,
		Calendars:        entityDbinstanceParam.Calendars,
		DBInstanceID:     entityDbinstanceParam.DBInstanceID,
		DBInstanceType:   entityDbinstanceParam.DBInstanceType,
		DBInstanceName:   entityDbinstanceParam.DBInstanceName,
		DBInstanceState:  entityDbinstanceParam.DBInstanceState,
		DBInstanceRegion: entityDbinstanceParam.DBInstanceRegion,
		Active:           entityDbinstanceParam.Active,
		CreatedAt:        entityDbinstanceParam.CreatedAt,
		UpdatedAt:        entityDbinstanceParam.UpdatedAt,
	}

	return u, nil
}

func (u *EntityDbinstance) Validate() error {
	return validator.New().Struct(u)
}

type SearchEntityDbinstanceParams struct {
	OrderBy   string `json:"order_by"`
	SortOrder string `json:"sort_order"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	Q         string `json:"q"`
	CreatedAt string `json:"created_at"`
}
