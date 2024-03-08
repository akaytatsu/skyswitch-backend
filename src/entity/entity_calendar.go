package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

const TYPE_ACTION_ON = "on"
const TYPE_ACTION_OFF = "off"

type EntityCalendar struct {
	ID           int       `json:"id" gorm:"column:id;primary_key"`
	Name         string    `json:"name" gorm:"column:name;varchar(80);not null"`
	TypeAction   string    `json:"type_action" gorm:"column:type_action;varchar(3);not null"`
	Active       bool      `json:"active" gorm:"column:active;default:true"`
	ExecuteTime  string    `json:"execute_time" gorm:"column:execute_time;varchar(5);not null;default:'00:00'"`
	ValidHoliday bool      `json:"valid_holiday" gorm:"column:valid_holiday;default:false"`
	ValidWeekend bool      `json:"valid_weekend" gorm:"column:valid_weekend;default:false"`
	Sunday       bool      `json:"sunday" gorm:"column:sunday;default:false"`
	Monday       bool      `json:"monday" gorm:"column:monday;default:false"`
	Tuesday      bool      `json:"tuesday" gorm:"column:tuesday;default:false"`
	Wednesday    bool      `json:"wednesday" gorm:"column:wednesday;default:false"`
	Thursday     bool      `json:"thursday" gorm:"column:thursday;default:false"`
	Friday       bool      `json:"friday" gorm:"column:friday;default:false"`
	Saturday     bool      `json:"saturday" gorm:"column:saturday;default:false"`
	CreatedAt    time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewEntityCalendar(entityCalendarParam EntityCalendar) (*EntityCalendar, error) {
	u := &EntityCalendar{
		ID:        entityCalendarParam.ID,
		Name:      entityCalendarParam.Name,
		Active:    entityCalendarParam.Active,
		CreatedAt: entityCalendarParam.CreatedAt,
		UpdatedAt: entityCalendarParam.UpdatedAt,
	}

	return u, nil
}

func (u *EntityCalendar) Validate() error {
	return validator.New().Struct(u)
}

type SearchEntityCalendarParams struct {
	OrderBy   string `json:"order_by"`
	SortOrder string `json:"sort_order"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	Q         string `json:"q"`
	CreatedAt string `json:"created_at"`
}
