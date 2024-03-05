package handlers

import (
	"app/entity"
	"app/infrastructure/repository"
	usecase_cloud_account "app/usecase/cloud_account"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CloudAccountHandlers struct {
	usecaseCloudAccount usecase_cloud_account.IUsecaseCloudAccount
}

func NewCloudAccountHandlers(usecaseCloudAccount usecase_cloud_account.IUsecaseCloudAccount) *CloudAccountHandlers {
	return &CloudAccountHandlers{usecaseCloudAccount: usecaseCloudAccount}
}

// @Summary Get all cloud accounts
// @Description Get all cloud accounts
// @Tags CloudAccount
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} entity.EntityCloudAccount "success"
// @Router /api/cloudaccount/ [get]
func (h *CloudAccountHandlers) GetAllCloudAccountHandle(c *gin.Context) {

	orderBy, sortOrder := getOrderByParams(c, "updated_at")
	pagina, tamanhoPagina := getPaginationParams(c)

	params := entity.SearchEntityCloudAccountParams{
		OrderBy:   orderBy,
		SortOrder: sortOrder,
		Page:      pagina,
		PageSize:  tamanhoPagina,
		Q:         c.Query("q"),
		CreatedAt: c.Query("created_at"),
	}

	holiday, totalRegs, err := h.usecaseCloudAccount.GetAll(params)

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

// @Summary Get cloud account by id
// @Description Get cloud account by id
// @Tags CloudAccount
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Cloud Account ID"
// @Success 200 {object} entity.EntityCloudAccount "success"
// @Router /api/cloudaccount/{id} [get]
func (h *CloudAccountHandlers) GetByIDCloudAccountHandle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	cloudAccount, err := h.usecaseCloudAccount.GetByID(int64(id))
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": cloudAccount,
	})
}

// @Summary Create cloud account
// @Description Create cloud account
// @Tags CloudAccount
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param entity.EntityCloudAccount body entity.EntityCloudAccount true "Cloud Account"
// @Param Authorization header string true "Bearer Token"
// @Success 201 {string} string "success"
// @Router /api/cloudaccount/ [post]
func (h *CloudAccountHandlers) CreateCloudAccountHandle(c *gin.Context) {
	var cloudAccount *entity.EntityCloudAccount
	err := c.ShouldBindJSON(&cloudAccount)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.usecaseCloudAccount.CreateCloudAccount(cloudAccount)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(201, gin.H{
		"message": "success",
	})
}

// @Summary Update cloud account
// @Description Update cloud account
// @Tags CloudAccount
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Success 200 {object} entity.EntityCloudAccount "success"
// @Router /api/cloudaccount/ [put]
func (h *CloudAccountHandlers) UpdateCloudAccountHandle(c *gin.Context) {
	var cloudAccount *entity.EntityCloudAccount
	err := c.ShouldBindJSON(&cloudAccount)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.usecaseCloudAccount.UpdateCloudAccount(cloudAccount)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// @Summary Delete cloud account
// @Description Delete cloud account
// @Tags CloudAccount
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Cloud Account ID"
// @Success 200 {string} string "success"
// @Router /api/cloudaccount/{id} [delete]
func (h *CloudAccountHandlers) DeleteCloudAccountHandle(c *gin.Context) {
	var cloudAccount *entity.EntityCloudAccount
	err := c.ShouldBindJSON(&cloudAccount)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	err = h.usecaseCloudAccount.DeleteCloudAccount(cloudAccount)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success",
	})
}

// @Summary Active/Deactive cloud account
// @Description Active/Deactive cloud account
// @Tags CloudAccount
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Cloud Account ID"
// @Param status path bool true "Status"
// @Success 200 {object} entity.EntityCloudAccount "success"
// @Router /api/cloudaccount/{id}/{status} [get]
func (h *CloudAccountHandlers) ActiveDeactiveCloudAccountHandle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	status, _ := strconv.ParseBool(c.Param("status"))

	cloudAccount, err := h.usecaseCloudAccount.ActiveDeactiveCloudAccount(int64(id), status)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": cloudAccount,
	})
}

// @Summary Update all instances of cloud account provider
// @Description Update all instances of cloud account provider
// @Tags CloudAccount
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Cloud Account ID"
// @Success 200 {string} string "success"
// @Router /api/cloudaccount/update-all-instances/{id} [get]
func (h *CloudAccountHandlers) UpdateAllInstanceOFCloudAccountProviderHandle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := h.usecaseCloudAccount.UpdateAllInstancesOnCloudAccountProviderFromID(id)

	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": "success",
	})
}

func MountCloudAccountHandlers(r *gin.Engine, conn *gorm.DB) {
	handlers := NewCloudAccountHandlers(usecase_cloud_account.NewAWSService(
		repository.NewCloudAccountPostgres(conn),
		repository.NewInstancePostgres(conn),
	))

	group := r.Group("api/cloudaccount")
	SetAuthMiddleware(conn, group)

	group.PUT("/update-all-instances/:id", handlers.UpdateAllInstanceOFCloudAccountProviderHandle)
	group.GET("/", handlers.GetAllCloudAccountHandle)
	group.GET("/:id", handlers.GetByIDCloudAccountHandle)
	group.POST("/", handlers.CreateCloudAccountHandle)
	group.PUT("/", handlers.UpdateCloudAccountHandle)
	group.DELETE("/:id", handlers.DeleteCloudAccountHandle)
	group.GET("/:id/:status", handlers.ActiveDeactiveCloudAccountHandle)
}
