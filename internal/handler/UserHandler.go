package handler

import (
	"TODO-List/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (s *UserHandler) RegisterEndpointsForUser(router *gin.Engine) {
	router.GET("/user/:userId", s.service.GetUser)
	router.POST("/register", s.service.CreateUser)
}
