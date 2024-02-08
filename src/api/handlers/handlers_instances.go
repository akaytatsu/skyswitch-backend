package handlers

import (
	"app/api/middleware"
	"app/infrastructure/repository"
	usecase_instance "app/usecase/instance"
	usecase_user "app/usecase/user"
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
	instances, err := h.usecaseInstances.GetAll()
	if exception := handleError(c, err); exception {
		return
	}
	jsonResponse(c, http.StatusOK, instances)
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

func MountInstancesRoutes(r *gin.Engine, conn *gorm.DB) {
	instanceHandlers := NewIntancesHandlers(usecase_instance.NewService(
		repository.NewInstancePostgres(conn),
	))

	usecaseUser := usecase_user.NewService(
		repository.NewUserPostgres(conn),
	)

	group := r.Group("/api/instances")
	group.Use(middleware.AuthenticatedMiddleware(usecaseUser))

	group.GET("/", instanceHandlers.GetAllInstancesHandler)
	group.GET("/:id", instanceHandlers.GetByIDInstanceHandler)
}
