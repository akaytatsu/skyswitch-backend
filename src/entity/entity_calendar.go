package entity

import "time"

const TYPE_ACTION_ON = "on"
const TYPE_ACTION_OFF = "off"

type EntityCalendar struct {
	ID            int
	Name          string
	TypeAction    string
	Active        bool
	ValidHoliday  bool
	ValidWeekend  bool
	Sunday        bool
	SundayDate    time.Time
	Monday        bool
	MondayDate    time.Time
	Tuesday       bool
	TuesdayDate   time.Time
	Wednesday     bool
	WednesdayDate time.Time
	Thursday      bool
	ThursdayDate  time.Time
	Friday        bool
	FridayDate    time.Time
	Saturday      bool
	SaturdayDate  time.Time
}
