package converter

import (
	"TODO-List/internal/model/request"
	"TODO-List/internal/model/response"
	"TODO-List/internal/repository/user/model"
)

func ValidateUserRequestToEntity(request *request.CreateUserRequest) (*model.User, error) {
	return &model.User{
		Username:  request.Username,
		Password:  request.Password,
		FirstName: request.FirstName,
		LastName:  request.LastName,
	}, nil
}
func ConvertUserEntityToResponse(user *model.User) (*response.UserResponse, error) {
	return &response.UserResponse{
		Id:        user.Id,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: user.CreatedAt,
	}, nil
}
