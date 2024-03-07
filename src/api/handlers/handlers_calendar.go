package handlers

import (
	"app/entity"
	usecase_calendar "app/usecase/calendar"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CalendarHandler struct {
	UsecaseCalendar usecase_calendar.IUsecaseCalendar
}

func NewCalendarHandler(usecase usecase_calendar.IUsecaseCalendar) *CalendarHandler {
	return &CalendarHandler{UsecaseCalendar: usecase}
}

// @Summary Get Calendar
// @Description Get Calendar
// @Tags Calendar
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} entity.EntityCalendar "success"
// @Router /api/calendar/{id} [get]
func (h *CalendarHandler) GetCalendar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}

	calendar, err := h.UsecaseCalendar.Get(id)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, calendar)
}

// @Summary Get All Calendar
// @Description Get All Calendar
// @Tags Calendar
// @Accept  json
// @Produce  json
// @Success 200 {object} []entity.EntityCalendar "success"
// @Router /api/calendar [get]
func (h *CalendarHandler) GetAllCalendar(c *gin.Context) {
	orderBy, sortOrder := getOrderByParams(c, "updated_at")
	pagina, tamanhoPagina := getPaginationParams(c)

	params := entity.SearchEntityCalendarParams{
		OrderBy:   orderBy,
		SortOrder: sortOrder,
		Page:      pagina,
		PageSize:  tamanhoPagina,
		Q:         c.Query("q"),
		CreatedAt: c.Query("created_at"),
	}

	calendar, totalRegs, err := h.UsecaseCalendar.GetAll(params)

	if exception := handleError(c, err); exception {
		return
	}

	paginationResponse := PaginationResponse{
		TotalPages:     getTotalPaginas(totalRegs, tamanhoPagina),
		Page:           pagina,
		TotalRegisters: int(totalRegs),
		Registers:      calendar,
	}

	c.JSON(http.StatusOK, paginationResponse)
}

// @Summary Create Calendar
// @Description Create Calendar
// @Tags Calendar
// @Accept  json
// @Produce  json
// @Param calendar body entity.EntityCalendar true "Calendar"
// @Success 201 {object} entity.EntityCalendar "success"
// @Router /api/calendar [post]
func (h *CalendarHandler) CreateCalendar(c *gin.Context) {
	var calendar entity.EntityCalendar

	if err := c.ShouldBindJSON(&calendar); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseCalendar.Create(&calendar)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusCreated, calendar)
}

// @Summary Update Calendar
// @Description Update Calendar
// @Tags Calendar
// @Accept  json
// @Produce  json
// @Param calendar body entity.EntityCalendar true "Calendar"
// @Success 200 {object} entity.EntityCalendar "success"
// @Router /api/calendar [put]
func (h *CalendarHandler) UpdateCalendar(c *gin.Context) {
	var calendar entity.EntityCalendar

	if err := c.ShouldBindJSON(&calendar); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseCalendar.Update(&calendar)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, calendar)
}

// @Summary Delete Calendar
// @Description Delete Calendar
// @Tags Calendar
// @Accept  json
// @Produce  json
// @Param id path int true "Calendar ID"
// @Success 200 {object} entity.EntityCalendar "success"
// @Router /api/calendar/{id} [delete]
func (h *CalendarHandler) DeleteCalendar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}

	err = h.UsecaseCalendar.Delete(id)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Calendar deleted successfully"})
}

func MountCalendarRoutes(gin *gin.Engine, conn *gorm.DB, usecase usecase_calendar.IUsecaseCalendar) {
	handler := NewCalendarHandler(usecase)

	group := gin.Group("/api/calendar")
	SetAdminMiddleware(conn, group)

	group.GET("/:id", handler.GetCalendar)
	group.GET("/", handler.GetAllCalendar)

	group.POST("/", handler.CreateCalendar)
	group.PUT("/:id", handler.UpdateCalendar)
	group.DELETE("/:id", handler.DeleteCalendar)
}
