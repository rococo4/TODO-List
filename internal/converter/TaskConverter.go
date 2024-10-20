package converter

import (
	"TODO-List/internal/model/request"
	"TODO-List/internal/model/response"
	"TODO-List/internal/repository/task/model"
	model2 "TODO-List/internal/repository/user/model"
)

func ConvertCreateTaskRequestToTaskEntity(request *request.CreateTaskRequest) (*model.Task, error) {
	return &model.Task{
		ExpiredAt:   request.ExpiredAt,
		Name:        request.Name,
		Description: request.Description,
	}, nil
}
func ConvertTaskEntityToTaskResponse(task *model.Task, user *model2.User) (*response.TaskResponse, error) {
	userResponse, err := ConvertUserEntityToResponse(user)
	if err != nil {
		return nil, err
	}
	return &response.TaskResponse{
		Id:          task.Id,
		CreatedAt:   task.CreatedAt,
		ExpiresAt:   task.ExpiredAt,
		Name:        task.Name,
		Description: task.Description,
		User:        *userResponse,
	}, nil
}
