package response

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

func ToAuthLoginResponse(response *AuthLoginResponse) *AuthLoginResponse {
	return &AuthLoginResponse{ID: response.ID, Username: response.Username, Token: response.Token}
}
