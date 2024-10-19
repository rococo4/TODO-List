package service

import (
	"TODO-List/internal/auth"
	"TODO-List/internal/converter"
	"TODO-List/internal/model/request"
	"TODO-List/internal/repository/task"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskService struct {
	repo *task.Repository
}

func NewTaskService(repo *task.Repository) *TaskService {
	return &TaskService{repo: repo}
}
func (service *TaskService) CreateTask(context *gin.Context) {

	token := context.GetHeader("Authorization")
	userId, err := auth.GetUserIdFromJwt(token)
	if _, err := service.repo.Find(userId); err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	var taskRequest request.CreateTaskRequest
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	if err := context.ShouldBindJSON(taskRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity, err := converter.ValidateConvertRequestToEntity(&taskRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity.UserId = userId
	create, err := service.repo.Create(*entity)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	taskResponse, err := converter.ConvertEntityToResponse(create)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, taskResponse)
}
func (service *TaskService) GetTask(context *gin.Context) {
	taskId, err := strconv.Atoi(context.Param("taskId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	taskEntity, err := service.repo.Find(taskId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, taskEntity)
}
