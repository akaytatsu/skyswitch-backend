package handlers

import (
	"app/entity"
	usecase_dbinstance "app/usecase/dbinstance"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DbinstanceHandler struct {
	UsecaseDbinstance usecase_dbinstance.IUsecaseDbinstance
}

func NewDbinstanceHandler(usecase usecase_dbinstance.IUsecaseDbinstance) *DbinstanceHandler {
	return &DbinstanceHandler{UsecaseDbinstance: usecase}
}

// @Summary Get Dbinstance
// @Description Get Dbinstance
// @Tags Dbinstance
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} entity.EntityDbinstance "success"
// @Router /api/dbinstance/{id} [get]
func (h *DbinstanceHandler) GetDbinstance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}

	dbinstance, err := h.UsecaseDbinstance.Get(id)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, dbinstance)
}

// @Summary Get All Dbinstance
// @Description Get All Dbinstance
// @Tags Dbinstance
// @Accept  json
// @Produce  json
// @Success 200 {object} []entity.EntityDbinstance "success"
// @Router /api/dbinstance [get]
func (h *DbinstanceHandler) GetAllDbinstance(c *gin.Context) {
	orderBy, sortOrder := getOrderByParams(c, "updated_at")
	pagina, tamanhoPagina := getPaginationParams(c)

	OnlyStatusMonitor := c.Query("only_status_monitor")

	if OnlyStatusMonitor == "" {
		OnlyStatusMonitor = "true"
	}

	OnlyActive := c.Query("only_active")
	if OnlyActive == "" {
		OnlyActive = "true"
	}

	ExcludeBlankName := c.Query("exclude_blank_name")
	if ExcludeBlankName == "" {
		ExcludeBlankName = "true"
	}

	params := entity.SearchEntityDbinstanceParams{
		OrderBy:   orderBy,
		SortOrder: sortOrder,
		Page:      pagina,
		PageSize:  tamanhoPagina,
		Q:         c.Query("q"),
		CreatedAt: c.Query("created_at"),
	}

	instances, totalRegs, err := h.UsecaseDbinstance.GetAll(params)
	if exception := handleError(c, err); exception {
		return
	}

	paginationResponse := PaginationResponse{
		TotalPages:     getTotalPaginas(totalRegs, tamanhoPagina),
		Page:           pagina,
		TotalRegisters: int(totalRegs),
		Registers:      instances,
	}

	c.JSON(http.StatusOK, paginationResponse)
}

// @Summary Create Dbinstance
// @Description Create Dbinstance
// @Tags Dbinstance
// @Accept  json
// @Produce  json
// @Param dbinstance body entity.EntityDbinstance true "Dbinstance"
// @Success 201 {object} entity.EntityDbinstance "success"
// @Router /api/dbinstance [post]
func (h *DbinstanceHandler) CreateDbinstance(c *gin.Context) {
	var dbinstance entity.EntityDbinstance

	if err := c.ShouldBindJSON(&dbinstance); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseDbinstance.Create(&dbinstance)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusCreated, dbinstance)
}

// @Summary Update Dbinstance
// @Description Update Dbinstance
// @Tags Dbinstance
// @Accept  json
// @Produce  json
// @Param dbinstance body entity.EntityDbinstance true "Dbinstance"
// @Success 200 {object} entity.EntityDbinstance "success"
// @Router /api/dbinstance [put]
func (h *DbinstanceHandler) UpdateDbinstance(c *gin.Context) {
	var dbinstance entity.EntityDbinstance

	if err := c.ShouldBindJSON(&dbinstance); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseDbinstance.Update(&dbinstance, true)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, dbinstance)
}

// @Summary Delete Dbinstance
// @Description Delete Dbinstance
// @Tags Dbinstance
// @Accept  json
// @Produce  json
// @Param id path int true "Dbinstance ID"
// @Success 200 {object} entity.EntityDbinstance "success"
// @Router /api/dbinstance/{id} [delete]
func (h *DbinstanceHandler) DeleteDbinstance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}

	err = h.UsecaseDbinstance.Delete(id)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dbinstance deleted successfully"})
}

func MountDbinstanceRoutes(gin *gin.Engine, conn *gorm.DB, usecase usecase_dbinstance.IUsecaseDbinstance) {
	handler := NewDbinstanceHandler(usecase)

	group := gin.Group("/api/dbinstances")

	group.GET("/:id", handler.GetDbinstance)
	group.GET("/", handler.GetAllDbinstance)

	SetAdminMiddleware(conn, group)

	group.POST("/", handler.CreateDbinstance)
	group.PUT("/:id", handler.UpdateDbinstance)
	group.DELETE("/:id", handler.DeleteDbinstance)
}
