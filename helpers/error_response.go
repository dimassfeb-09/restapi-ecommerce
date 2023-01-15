package helpers

import "github.com/dimassfeb-09/restapi-ecommerce.git/entity/web"

func ToErrorResponse(code int, status string, msg any) *web.ErrorResponse {
	return &web.ErrorResponse{
		Code:    code,
		Status:  status,
		Message: msg,
	}
}
