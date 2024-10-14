package handler

import (
	"TODO-List/internal/service"
	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(service *service.TaskService) *TaskHandler {
	return &TaskHandler{service: service}
}

func (s *TaskHandler) RegisterEndpointsForTasks(router *gin.Engine) {
	router.GET("/task/:taskId", s.service.GetTask)
	router.POST("/task", s.service.CreateTask)
}
