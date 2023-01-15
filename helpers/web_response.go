package response

import "github.com/dimassfeb-09/restapi-ecommerce.git/entity/web"

func ToWebResponse(code int, status string, data any) *web.WebResponse {
	return &web.WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
