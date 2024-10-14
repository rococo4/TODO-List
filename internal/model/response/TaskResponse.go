package response

import "time"

type TaskResponse struct {
	Id          int       `json:"id"`
	CreatedAt   time.Time `json:"created"`
	ExpiresAt   time.Time `json:"expires"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
