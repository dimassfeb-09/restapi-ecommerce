package response

import (
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
	"time"
)

type UserResponse struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Username  string    `json:"username"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}

func ToUserResponse(user *domain.Users) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}
}

func ToUserResponses(users []*domain.Users) []*UserResponse {
	var usersResponse []*UserResponse
	for i := 0; i < len(users); i++ {
		usersResponse = append(usersResponse, ToUserResponse(users[i]))
	}
	return usersResponse
}
