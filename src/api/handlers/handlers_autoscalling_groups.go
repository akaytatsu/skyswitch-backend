package handlers

import (
	"app/entity"
	usecase_autoscalling_groups "app/usecase/autoscalling_groups"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AutoScalingGroupHandler struct {
	UsecaseAutoScalingGroup usecase_autoscalling_groups.IUsecaseAutoScalingGroup
}

func NewAutoScalingGroupHandler(usecase usecase_autoscalling_groups.IUsecaseAutoScalingGroup) *AutoScalingGroupHandler {
	return &AutoScalingGroupHandler{UsecaseAutoScalingGroup: usecase}
}

// @Summary Get AutoScalingGroup
// @Description Get AutoScalingGroup
// @Tags AutoScalingGroup
// @Accept  json
// @Produce  json
// @Param id path int true "ID"
// @Success 200 {object} entity.EntityAutoScalingGroup "success"
// @Router /api/autoscallinggroup/{id} [get]
func (h *AutoScalingGroupHandler) GetAutoScalingGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}

	autoscallinggroup, err := h.UsecaseAutoScalingGroup.Get(id)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, autoscallinggroup)
}

// @Summary Get All AutoScalingGroup
// @Description Get All AutoScalingGroup
// @Tags AutoScalingGroup
// @Accept  json
// @Produce  json
// @Success 200 {object} []entity.EntityAutoScalingGroup "success"
// @Router /api/autoscallinggroup [get]
func (h *AutoScalingGroupHandler) GetAllAutoScalingGroup(c *gin.Context) {
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

	params := entity.SearchEntityAutoScalingGroupParams{
		OrderBy:   orderBy,
		SortOrder: sortOrder,
		Page:      pagina,
		PageSize:  tamanhoPagina,
		Q:         c.Query("q"),
		CreatedAt: c.Query("created_at"),
	}

	instances, totalRegs, err := h.UsecaseAutoScalingGroup.GetAll(params)
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

// @Summary Create AutoScalingGroup
// @Description Create AutoScalingGroup
// @Tags AutoScalingGroup
// @Accept  json
// @Produce  json
// @Param autoscallinggroup body entity.EntityAutoScalingGroup true "AutoScalingGroup"
// @Success 201 {object} entity.EntityAutoScalingGroup "success"
// @Router /api/autoscallinggroup [post]
func (h *AutoScalingGroupHandler) CreateAutoScalingGroup(c *gin.Context) {
	var autoscallinggroup entity.EntityAutoScalingGroup

	if err := c.ShouldBindJSON(&autoscallinggroup); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseAutoScalingGroup.Create(&autoscallinggroup)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusCreated, autoscallinggroup)
}

// @Summary Update AutoScalingGroup
// @Description Update AutoScalingGroup
// @Tags AutoScalingGroup
// @Accept  json
// @Produce  json
// @Param autoscallinggroup body entity.EntityAutoScalingGroup true "AutoScalingGroup"
// @Success 200 {object} entity.EntityAutoScalingGroup "success"
// @Router /api/autoscallinggroup [put]
func (h *AutoScalingGroupHandler) UpdateAutoScalingGroup(c *gin.Context) {
	var autoscallinggroup entity.EntityAutoScalingGroup

	if err := c.ShouldBindJSON(&autoscallinggroup); err != nil {
		handleError(c, err)
		return
	}

	err := h.UsecaseAutoScalingGroup.Update(&autoscallinggroup)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, autoscallinggroup)
}

// @Summary Delete AutoScalingGroup
// @Description Delete AutoScalingGroup
// @Tags AutoScalingGroup
// @Accept  json
// @Produce  json
// @Param id path int true "AutoScalingGroup ID"
// @Success 200 {object} entity.EntityAutoScalingGroup "success"
// @Router /api/autoscallinggroup/{id} [delete]
func (h *AutoScalingGroupHandler) DeleteAutoScalingGroup(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handleError(c, err)
		return
	}

	err = h.UsecaseAutoScalingGroup.Delete(id)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "AutoScalingGroup deleted successfully"})
}

func MountAutoScalingGroupRoutes(gin *gin.Engine, conn *gorm.DB, usecase usecase_autoscalling_groups.IUsecaseAutoScalingGroup) {
	handler := NewAutoScalingGroupHandler(usecase)

	group := gin.Group("/api/autoscallinggroup")

	group.GET("/:id", handler.GetAutoScalingGroup)
	group.GET("/", handler.GetAllAutoScalingGroup)

	SetAdminMiddleware(conn, group)

	group.POST("/", handler.CreateAutoScalingGroup)
	group.PUT("/:id", handler.UpdateAutoScalingGroup)
	group.DELETE("/:id", handler.DeleteAutoScalingGroup)
}
