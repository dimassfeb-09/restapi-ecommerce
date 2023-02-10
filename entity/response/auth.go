package response

import "github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"

type AuthLogin struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type AuthLoginResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func ToAuthLoginResponse(response *domain.AuthUser) *AuthLoginResponse {
	return &AuthLoginResponse{ID: response.ID, Username: response.Username}
}
