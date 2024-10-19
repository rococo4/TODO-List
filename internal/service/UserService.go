package service

import (
	"TODO-List/internal/auth"
	"TODO-List/internal/converter"
	"TODO-List/internal/model/request"
	"TODO-List/internal/repository/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserService struct {
	repo *user.Repository
}

func NewUserService(repo *user.Repository) *UserService {
	return &UserService{repo: repo}
}
func (service *UserService) CreateUser(context *gin.Context) {
	var userRequest request.CreateUserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	entity, err := converter.ValidateUserRequestToEntity(&userRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	create, err := service.repo.Create(*entity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	token, err := auth.GenerateJWT(create.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"token": token})
}
func (service *UserService) GetUser(context *gin.Context) {
	userId, err := strconv.Atoi(context.Param("userId"))

	userEntity, err := service.repo.Find(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userResponse, err := converter.ConvertUserEntityToResponse(userEntity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, userResponse)
}
func (service *UserService) LoginUser(context *gin.Context) {
	var loginRequest request.LoginRequest
	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userLogin, err := service.repo.FindByUsername(loginRequest.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"User with username not found": err.Error()})
	}
	err = auth.CheckPasswordHash(userLogin.Password, loginRequest.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"WrongPassword": err.Error()})
	}
	token, err := auth.GenerateJWT(userLogin.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusOK, gin.H{"token": token})
}
