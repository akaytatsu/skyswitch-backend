package handlers

import (
	"app/api/middleware"
	"app/entity"
	"app/infrastructure/repository"
	usecase_cloud_account "app/usecase/cloud_account"
	usecase_user "app/usecase/user"
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
// @Router /api/cloud_account/ [get]
func (h *CloudAccountHandlers) GetAllCloudAccountHandle(c *gin.Context) {
	cloudAccounts, err := h.usecaseCloudAccount.GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"data": cloudAccounts,
	})
}

// @Summary Get cloud account by id
// @Description Get cloud account by id
// @Tags CloudAccount
// @Accept  json
// @Produce  json
// @Security ApiKeyAuth
// @Param id path int true "Cloud Account ID"
// @Success 200 {object} entity.EntityCloudAccount "success"
// @Router /api/cloud_account/{id} [get]
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
// @Router /api/cloud_account/ [put]
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
// @Router /api/cloud_account/{id} [delete]
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
// @Router /api/cloud_account/{id}/{status} [get]
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

func MountCloudAccountHandlers(r *gin.Engine, conn *gorm.DB) {
	handlers := NewCloudAccountHandlers(usecase_cloud_account.NewAWSService(
		repository.NewCloudAccountPostgres(conn),
		repository.NewInstancePostgres(conn),
	))

	usecaseUser := usecase_user.NewService(
		repository.NewUserPostgres(conn),
	)

	group := r.Group("api/cloud_account")
	group.Use(middleware.AuthenticatedMiddleware(usecaseUser))

	group.GET("/", handlers.GetAllCloudAccountHandle)
	group.GET("/:id", handlers.GetByIDCloudAccountHandle)
	group.POST("/", handlers.CreateCloudAccountHandle)
	group.PUT("/", handlers.UpdateCloudAccountHandle)
	group.DELETE("/:id", handlers.DeleteCloudAccountHandle)
	group.GET("/:id/:status", handlers.ActiveDeactiveCloudAccountHandle)
}
