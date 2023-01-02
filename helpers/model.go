package helpers

import (
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/response"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/web"
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

func ToWebResponse(code int, status string, data any) *web.WebResponse {
	return &web.WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}

func ToErrorResponse(code int, status string, desc any) *web.ErrorResponse {
	return &web.ErrorResponse{
		Code:        code,
		Status:      status,
		Description: desc,
	}
}
