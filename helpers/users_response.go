package response

import (
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/response"
)

func ToUserResponse(user *domain.Users) *response.UserResponse {
	return &response.UserResponse{
		Id:        user.ID,
		Name:      user.Name,
		Username:  user.Username,
		Balance:   user.Balance,
		CreatedAt: user.CreatedAt,
	}
}

func ToUserResponses(users []*domain.Users) []*response.UserResponse {
	var usersResponse []*response.UserResponse
	for i := 0; i < len(users); i++ {
		usersResponse = append(usersResponse, ToUserResponse(users[i]))
	}
	return usersResponse
}
