package handlers

import (
	"app/entity"
	"app/infrastructure/repository"
	usecase_instance "app/usecase/instance"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type IntancesHandlers struct {
	usecaseInstances usecase_instance.IUseCaseInstance
}

func NewIntancesHandlers(usecaseInstances usecase_instance.IUseCaseInstance) *IntancesHandlers {
	return &IntancesHandlers{usecaseInstances: usecaseInstances}
}

// @Summary Get all instances
// @Description Get all instances
// @Tags Instances
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer Token"
// @Security ApiKeyAuth
// @Success 200 {object} entity.EntityInstance "success"
// @Router /api/instances/ [get]
func (h *IntancesHandlers) GetAllInstancesHandler(c *gin.Context) {

	orderBy, sortOrder := getOrderByParams(c, "updated_at")
	pagina, tamanhoPagina := getPaginationParams(c)

	params := entity.SearchEntityInstanceParams{
		OrderBy:   orderBy,
		SortOrder: sortOrder,
		Page:      pagina,
		PageSize:  tamanhoPagina,
		Q:         c.Query("q"),
		CreatedAt: c.Query("created_at"),
	}

	instances, totalRegs, err := h.usecaseInstances.GetAll(params)
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

// @Summary Get instance by id
// @Description Get instance by id
// @Tags Instances
// @Accept  json
// @Produce  json
// @Param id path int true "Instance ID"
// @Param Authorization header string true "Bearer Token"
// @Security ApiKeyAuth
// @Success 200 {object} entity.EntityInstance "success"
// @Router /api/instances/{id} [get]
func (h *IntancesHandlers) GetByIDInstanceHandler(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	instance, err := h.usecaseInstances.GetByID(int64(id))
	if exception := handleError(c, err); exception {
		return
	}
	jsonResponse(c, http.StatusOK, instance)
}

// @Summary Update instance
// @Description Update instance
// @Tags Instances
// @Accept  json
// @Produce  json
// @Param id path int true "Instance ID"
// @Param entity.EntityInstance body entity.EntityInstance true "Instance"
// @Param Authorization header string true "
// @Security ApiKeyAuth
// @Success 200 {object} entity.EntityInstance "success"
// @Router /api/instances/{id} [put]
func (h *IntancesHandlers) UpdateInstanceHandler(c *gin.Context) {
	var instance entity.EntityInstance

	if err := c.ShouldBindJSON(&instance); err != nil {
		handleError(c, err)
		return
	}

	err := h.usecaseInstances.UpdateInstance(&instance, true)

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, instance)
}

func MountInstancesRoutes(r *gin.Engine, conn *gorm.DB) {
	instanceHandlers := NewIntancesHandlers(usecase_instance.NewService(
		repository.NewInstancePostgres(conn),
	))

	group := r.Group("/api/instances")
	SetAuthMiddleware(conn, group)

	group.GET("/", instanceHandlers.GetAllInstancesHandler)
	group.GET("/:id", instanceHandlers.GetByIDInstanceHandler)
	group.PUT("/:id", instanceHandlers.UpdateInstanceHandler)
}
