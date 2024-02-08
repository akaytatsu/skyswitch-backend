package handlers

import (
	"app/entity"
	"app/infrastructure/repository"
	usecase_holiday "app/usecase/holiday"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HolidayHandler struct {
	usecaseHoliday usecase_holiday.IUsecaseHoliday
}

func NewHolidayHandlers(usecaseHoliday usecase_holiday.IUsecaseHoliday) *HolidayHandler {
	return &HolidayHandler{usecaseHoliday: usecaseHoliday}
}

// @Summary Get All Holidays
// @Description Get All Holidays
// @Tags Holiday
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} entity.EntityHoliday "success"
// @Router /api/holiday/ [get]
func (h *HolidayHandler) GetAllHolidaysHandler(c *gin.Context) {
	holidays, err := h.usecaseHoliday.GetAll()
	if exception := handleError(c, err); exception {
		return
	}
	jsonResponse(c, http.StatusOK, holidays)
}

// @Summary Create Or Update Holiday
// @Description Create Or Update Holiday
// @Tags Holiday
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param entity.EntityHoliday body entity.EntityHoliday true "Holiday"
// @Param Authorization header string true "Bearer Token"
// @Success 200 {object} entity.EntityHoliday "success"
// @Router /api/holiday/ [post]
func (h *HolidayHandler) CreateUpdateHolidayHandler(c *gin.Context) {
	var holiday entity.EntityHoliday
	err := c.ShouldBindJSON(&holiday)
	if exception := handleError(c, err); exception {
		return
	}

	holiday, err = h.usecaseHoliday.CreateUpdate(holiday.Name, holiday.Date)
	if exception := handleError(c, err); exception {
		return
	}

	jsonResponse(c, http.StatusOK, holiday)
}

func MountHolidayRoutes(r *gin.Engine, conn *gorm.DB) {

	repoHoliday := repository.NewHolidayPostgres(conn)

	holidayHandler := NewHolidayHandlers(
		usecase_holiday.NewHolidayService(repoHoliday),
	)

	group := r.Group("/api/holiday")

	SetAuthMiddleware(conn, group)

	group.GET("/", holidayHandler.GetAllHolidaysHandler)

}
