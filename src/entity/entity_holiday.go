package entity

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type EntityHoliday struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Date      time.Time `json:"date" gorm:"column:date;type:date;unique;primary_key;not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func NewEntityHoliday(entityHolidayParam EntityHoliday) (*EntityHoliday, error) {
	u := &EntityHoliday{
		ID:        entityHolidayParam.ID,
		Date:      entityHolidayParam.Date,
		UpdatedAt: entityHolidayParam.UpdatedAt,
	}

	return u, nil
}

func (u *EntityHoliday) Validate() error {
	return validator.New().Struct(u)
}

type SearchEntityHolidayParams struct {
	OrderBy   string `json:"order_by"`
	SortOrder string `json:"sort_order"`
	Page      int    `json:"page"`
	PageSize  int    `json:"page_size"`
	Q         string `json:"q"`
	CreatedAt string `json:"created_at"`
}
