package response

import "time"

type UserResponse struct {
	Id        int       `json:"id"`
	Username  string    `json:"name"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
}
