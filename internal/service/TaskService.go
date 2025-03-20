package service

import (
	"TODO-List/internal/auth"
	"TODO-List/internal/converter"
	"TODO-List/internal/model/request"
	"TODO-List/internal/repository/task"
	"TODO-List/internal/repository/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type TaskService struct {
	taskRepository *task.Repository
	userRepository *user.Repository
}

func NewTaskService(taskRepo *task.Repository, userRepo *user.Repository) *TaskService {
	return &TaskService{taskRepository: taskRepo, userRepository: userRepo}
}
func (service *TaskService) CreateTask(context *gin.Context) {

	token := context.GetHeader("Authorization")

	userId, err := auth.GetUserIdFromJwt(token)
	userEntity, err := service.userRepository.Find(userId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}

	var taskRequest request.CreateTaskRequest
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
	}

	if err := context.ShouldBindJSON(&taskRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity, err := converter.ConvertCreateTaskRequestToTaskEntity(&taskRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity.UserId = userId
	create, err := service.taskRepository.Create(*entity)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	taskResponse, err := converter.ConvertTaskEntityToTaskResponse(create, userEntity)
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
	taskEntity, err := service.taskRepository.Find(taskId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userEntity, err := service.userRepository.Find(taskEntity.UserId)
	taskResponse, err := converter.ConvertTaskEntityToTaskResponse(taskEntity, userEntity)
	context.JSON(http.StatusOK, taskResponse)
}
func (service *TaskService) DeleteTask(context *gin.Context) {
	taskId, err := strconv.Atoi(context.Param("taskId"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	err = service.taskRepository.Delete(taskId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
