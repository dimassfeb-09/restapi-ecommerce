package response

import "time"

type UserResponse struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
