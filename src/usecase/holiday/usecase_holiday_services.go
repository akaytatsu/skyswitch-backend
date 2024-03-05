package usecase_holiday

// import (
// 	"app/entity"
// 	"time"
// )

// type UsecaseHoliday struct {
// 	repo IRepositoryHoliday
// }

// func NewHolidayService(repository IRepositoryHoliday) *UsecaseHoliday {
// 	return &UsecaseHoliday{repo: repository}
// }

// func (u *UsecaseHoliday) GetAll() (holidays []entity.EntityHoliday, err error) {
// 	return u.repo.GetAll()
// }

// func (u *UsecaseHoliday) GetByDate(date time.Time) (holiday entity.EntityHoliday, err error) {
// 	return u.repo.GetByDate(date)
// }

// func (u *UsecaseHoliday) CreateUpdate(name string, date time.Time) (holiday entity.EntityHoliday, err error) {
// 	return u.repo.CreateUpdate(name, date)
// }

// func (u *UsecaseHoliday) DateStringToTime(date string) (time.Time, error) {
// 	return time.Parse("2006-01-02", date)
// }
