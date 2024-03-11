package handlers

import (
	usecase_job "app/usecase/job"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type JobHandler struct {
	UsecaseJob usecase_job.IUsecaseJob
}

func NewJobHandler(usecase usecase_job.IUsecaseJob) *JobHandler {
	return &JobHandler{UsecaseJob: usecase}
}

// @Summary Get All Job
// @Description Get All Job
// @Tags Job
// @Accept  json
// @Produce  json
// @Success 200 {object} []entity.EntityJob "success"
// @Router /api/job [get]
func (h *JobHandler) GetAllJob(c *gin.Context) {
	jobs, err := h.UsecaseJob.GetAll()

	if exception := handleError(c, err); exception {
		return
	}

	c.JSON(http.StatusOK, jobs)
}

func MountJobRoutes(gin *gin.Engine, conn *gorm.DB, usecase usecase_job.IUsecaseJob) {
	handler := NewJobHandler(usecase)

	group := gin.Group("/api/job")

	SetAdminMiddleware(conn, group)
	group.GET("/", handler.GetAllJob)
}
