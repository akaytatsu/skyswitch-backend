package handlers

import (
	"app/entity"
	usecase_holiday "app/usecase/holiday"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HolidayHandler struct {
	UsecaseHoliday usecase_holiday.IUsecaseHoliday
}

func NewHolidayHandler(usecase usecase_holiday.IUsecaseHoliday) *HolidayHandler {
	return &HolidayHandler{UsecaseHoliday: usecase}
}

// @Summary Get Holiday
// @Description Get Holiday
// @Tags Holiday
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} entity.EntityHoliday "success"
// @Router /api/holiday/{id} [get]
func (h *HolidayHandler) GetHoliday(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}

	holiday, err := h.UsecaseHoliday.Get(id)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, holiday)
}

// @Summary Get All Holiday
// @Description Get All Holiday
// @Tags Holiday
// @Accept  json
// @Produce  json
// @Success 200 {object} []entity.EntityHoliday "success"
// @Router /api/holiday [get]
func (h *HolidayHandler) GetAllHoliday(c *gin.Context) {
	orderBy, sortOrder := getOrderByParams(c, "updated_at")
	pagina, tamanhoPagina := getPaginationParams(c)

	params := entity.SearchEntityHolidayParams{
		OrderBy:   orderBy,
		SortOrder: sortOrder,
		Page:      pagina,
		PageSize:  tamanhoPagina,
		Q:         c.Query("q"),
		CreatedAt: c.Query("created_at"),
	}

	holiday, totalRegs, err := h.UsecaseHoliday.GetAll(params)

	if exception := handleError(c, err); exception {
		return
	}

	paginationResponse := PaginationResponse{
		TotalPages:     getTotalPaginas(totalRegs, tamanhoPagina),
		Page:           pagina,
		TotalRegisters: int(totalRegs),
		Registers:      holiday,
	}

	c.JSON(http.StatusOK, paginationResponse)
}

// @Summary Create Holiday
// @Description Create Holiday
// @Tags Holiday
// @Accept  json
// @Produce  json
// @Param holiday body entity.EntityHoliday true "Holiday"
// @Success 201 {object} entity.EntityHoliday "success"
// @Router /api/holiday [post]
func (h *HolidayHandler) CreateHoliday(c *gin.Context) {
	var holiday entity.EntityHoliday

	if err := c.ShouldBindJSON(&holiday); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseHoliday.Create(&holiday)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusCreated, holiday)
}

// @Summary Update Holiday
// @Description Update Holiday
// @Tags Holiday
// @Accept  json
// @Produce  json
// @Param holiday body entity.EntityHoliday true "Holiday"
// @Success 200 {object} entity.EntityHoliday "success"
// @Router /api/holiday [put]
func (h *HolidayHandler) UpdateHoliday(c *gin.Context) {
	var holiday entity.EntityHoliday

	if err := c.ShouldBindJSON(&holiday); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseHoliday.Update(&holiday)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, holiday)
}

// @Summary Delete Holiday
// @Description Delete Holiday
// @Tags Holiday
// @Accept  json
// @Produce  json
// @Param id path int true "Holiday ID"
// @Success 200 {object} entity.EntityHoliday "success"
// @Router /api/holiday/{id} [delete]
func (h *HolidayHandler) DeleteHoliday(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}

	err = h.UsecaseHoliday.Delete(id)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Holiday deleted successfully"})
}

func MountHolidayRoutes(gin *gin.Engine, conn *gorm.DB, usecase usecase_holiday.IUsecaseHoliday) {
	handler := NewHolidayHandler(usecase)

	group := gin.Group("/api/holiday")

	group.GET("/:id", handler.GetHoliday)
	group.GET("/", handler.GetAllHoliday)

	SetAdminMiddleware(conn, group)

	group.POST("/", handler.CreateHoliday)
	group.PUT("/:id", handler.UpdateHoliday)
	group.DELETE("/:id", handler.DeleteHoliday)
}
