package entity

import "time"

type EntityHoliday struct {
	Name      string    `json:"name" gorm:"column:name;type:varchar(70);default:''"`
	Date      time.Time `json:"date" gorm:"column:date;type:date;unique;primary_key;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
