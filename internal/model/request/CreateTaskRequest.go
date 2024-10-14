package request

import "time"

type CreateTaskRequest struct {
	ExpiredAt   time.Time `json:"expiredAt" binding:"required"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
}
