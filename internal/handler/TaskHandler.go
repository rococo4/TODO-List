package handler

import (
	"TODO-List/internal/auth"
	"TODO-List/internal/service"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (s *TaskHandler) RegisterEndpointsForTasks(r *gin.Engine) {
	r.GET("/task/:taskId", auth.VerifyToken(), s.service.GetTask)
	r.POST("/task", auth.VerifyToken(), s.service.CreateTask)
	r.DELETE("/task/:taskId", auth.VerifyToken(), s.service.DeleteTask)
}
