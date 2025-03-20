package service

import (
	"TODO-List/internal/auth"
	"TODO-List/internal/converter"
	"TODO-List/internal/model/request"
	meth "TODO-List/internal/prometheus"
	"TODO-List/internal/repository/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type UserService struct {
	repo *user.Repository
}

func NewUserService(repo *user.Repository) *UserService {
	return &UserService{repo: repo}
}
func (service *UserService) CreateUser(context *gin.Context) {
	meth.RequestCounter.Inc()
	start := time.Now()
	var userRequest request.CreateUserRequest

	if err := context.ShouldBindJSON(&userRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		duration := time.Since(start).Seconds()
		meth.HttpDuration.WithLabelValues("CreateUser", "400").Observe(duration)
		return
	}

	entity, err := converter.ValidateUserRequestToEntity(&userRequest)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		duration := time.Since(start).Seconds()
		meth.HttpDuration.WithLabelValues("CreateUser", "400").Observe(duration)
		return
	}
	create, err := service.repo.Create(*entity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		duration := time.Since(start).Seconds()
		meth.HttpDuration.WithLabelValues("CreateUser", "500").Observe(duration)
		return
	}
	token, err := auth.GenerateJWT(create.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"token": token})
	duration := time.Since(start).Seconds()
	meth.HttpDuration.WithLabelValues("CreateUser", "200").Observe(duration)
}
func (service *UserService) GetUser(context *gin.Context) {
	meth.RequestCounter.Inc()
	start := time.Now()

	userId, err := strconv.Atoi(context.Param("userId"))

	userEntity, err := service.repo.Find(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		duration := time.Since(start).Seconds()
		meth.HttpDuration.WithLabelValues("GetUser", "500").Observe(duration)
		return
	}

	userResponse, err := converter.ConvertUserEntityToResponse(userEntity)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		duration := time.Since(start).Seconds()
		meth.HttpDuration.WithLabelValues("GetUser", "500").Observe(duration)
	}
	duration := time.Since(start).Seconds()
	meth.HttpDuration.WithLabelValues("GetUser", "200").Observe(duration)
	context.JSON(http.StatusOK, userResponse)
}
func (service *UserService) LoginUser(context *gin.Context) {
	meth.RequestCounter.Inc()
	start := time.Now()
	var loginRequest request.LoginRequest
	if err := context.ShouldBindJSON(&loginRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		duration := time.Since(start).Seconds()
		meth.HttpDuration.WithLabelValues("LoginUser", "400").Observe(duration)
		return
	}
	userLogin, err := service.repo.FindByUsername(loginRequest.Username)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"User with username not found": err.Error()})
		duration := time.Since(start).Seconds()
		meth.HttpDuration.WithLabelValues("LoginUser", "400").Observe(duration)
	}
	err = auth.CheckPasswordHash(userLogin.Password, loginRequest.Password)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"WrongPassword": err.Error()})
		duration := time.Since(start).Seconds()
		meth.HttpDuration.WithLabelValues("LoginUser", "400").Observe(duration)
	}
	token, err := auth.GenerateJWT(userLogin.Id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		duration := time.Since(start).Seconds()
		meth.HttpDuration.WithLabelValues("LoginUser", "500").Observe(duration)
	}
	context.JSON(http.StatusOK, gin.H{"token": token})
	duration := time.Since(start).Seconds()
	meth.HttpDuration.WithLabelValues("LoginUser", "200").Observe(duration)
}
