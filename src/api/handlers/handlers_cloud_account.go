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
	))

	usecaseUser := usecase_user.NewService(
		repository.NewUserPostgres(conn),
	)

	group := r.Group("/cloud_account")
	group.Use(middleware.AuthenticatedMiddleware(usecaseUser))

	group.GET("/cloud_account", handlers.GetAllCloudAccountHandle)
	group.GET("/cloud_account/:id", handlers.GetByIDCloudAccountHandle)
	group.POST("/cloud_account", handlers.CreateCloudAccountHandle)
	group.PUT("/cloud_account", handlers.UpdateCloudAccountHandle)
	group.DELETE("/cloud_account", handlers.DeleteCloudAccountHandle)
	group.GET("/cloud_account/:id/:status", handlers.ActiveDeactiveCloudAccountHandle)
}
