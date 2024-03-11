package handlers

import (
	"app/entity"
	usecase_log "app/usecase/log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type LogHandler struct {
	UsecaseLog usecase_log.IUsecaseLog
}

func NewLogHandler(usecase usecase_log.IUsecaseLog) *LogHandler {
	return &LogHandler{UsecaseLog: usecase}
}

// @Summary Get Log
// @Description Get Log
// @Tags Log
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} entity.EntityLog "success"
// @Router /api/log/{id} [get]
func (h *LogHandler) GetLog(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}

	log, err := h.UsecaseLog.Get(id)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, log)
}

// @Summary Get All Log
// @Description Get All Log
// @Tags Log
// @Accept  json
// @Produce  json
// @Success 200 {object} []entity.EntityLog "success"
// @Router /api/log [get]
func (h *LogHandler) GetAllLog(c *gin.Context) {
	orderBy, sortOrder := getOrderByParams(c, "created_at")
	pagina, tamanhoPagina := getPaginationParams(c)

	params := entity.SearchEntityLogParams{
		OrderBy:   orderBy,
		SortOrder: sortOrder,
		Page:      pagina,
		PageSize:  tamanhoPagina,
		Q:         c.Query("q"),
		CreatedAt: c.Query("created_at"),
	}

	log, totalRegs, err := h.UsecaseLog.GetAll(params)

	if exception := handleError(c, err); exception {
		return
	}

	paginationResponse := PaginationResponse{
		TotalPages:     getTotalPaginas(totalRegs, tamanhoPagina),
		Page:           pagina,
		TotalRegisters: int(totalRegs),
		Registers:      log,
	}

	c.JSON(http.StatusOK, paginationResponse)
}

func MountLogRoutes(gin *gin.Engine, conn *gorm.DB, usecase usecase_log.IUsecaseLog) {
	handler := NewLogHandler(usecase)

	group := gin.Group("/api/log")
	SetAdminMiddleware(conn, group)

	group.GET("/:id", handler.GetLog)
	group.GET("/", handler.GetAllLog)

}
