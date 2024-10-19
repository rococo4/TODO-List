package handler

import (
	"TODO-List/internal/auth"
	"TODO-List/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (s *UserHandler) RegisterEndpointsForUser(r *gin.Engine) {

	r.GET("/user/:userId", auth.VerifyToken(), s.service.GetUser)

	r.POST("/register", s.service.CreateUser)

	r.POST("/login", s.service.LoginUser)
}
