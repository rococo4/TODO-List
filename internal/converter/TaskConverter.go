package converter

import (
	"TODO-List/internal/model/request"
	"TODO-List/internal/model/response"
	"TODO-List/internal/repository/task/model"
)

func ValidateConvertRequestToEntity(request *request.CreateTaskRequest) (*model.Task, error) {
	return &model.Task{
		ExpiredAt:   request.ExpiredAt,
		Name:        request.Name,
		Description: request.Description,
	}, nil
}
func ConvertEntityToResponse(task *model.Task) (*response.TaskResponse, error) {
	return &response.TaskResponse{
		Id:          task.Id,
		CreatedAt:   task.CreatedAt,
		ExpiresAt:   task.ExpiredAt,
		Name:        task.Name,
		Description: task.Description,
	}, nil
}
